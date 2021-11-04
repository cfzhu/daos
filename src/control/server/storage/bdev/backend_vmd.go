//
// (C) Copyright 2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

package bdev

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"

	"github.com/daos-stack/daos/src/control/common"
	"github.com/daos-stack/daos/src/control/logging"
	"github.com/daos-stack/daos/src/control/provider/system"
	"github.com/daos-stack/daos/src/control/server/storage"
)

// backingAddrToVMD converts a VMD backing devices address e.g. 5d0505:03.00.0
// to the relevant logical VMD address e.g. 0000:5d:05.5.
func backingAddrToVMD(in common.PCIAddress) (common.PCIAddress, bool) {
	if !in.IsVMDBackingAddress() {
		return common.PCIAddress{}, false
	}

	dom := in.Domain
	newAddrStr := fmt.Sprintf("0000:%c%c:%c%c.%c", dom[0], dom[1], dom[2], dom[3], dom[5])

	addr, err := common.NewPCIAddress(newAddrStr)
	if err != nil {
		panic(err)
	}

	return *addr, true
}

// mapVMDToBackingAddrs stores found vmd backing addresses under vmd address key.
func mapVMDToBackingAddrs(foundCtrlrs storage.NvmeControllers) (map[common.PCIAddress]*common.PCIAddressList, error) {
	vmds := make(map[common.PCIAddress]*common.PCIAddressList)

	for _, ctrlr := range foundCtrlrs {
		addr, err := common.NewPCIAddress(ctrlr.PciAddr)
		if err != nil {
			return nil, errors.Wrap(err, "controller pci address invalid")
		}

		// find backing device addresses from vmd address
		vmdAddr, isVMDBackingAddr := backingAddrToVMD(*addr)
		if !isVMDBackingAddr {
			continue
		}

		if _, exists := vmds[vmdAddr]; !exists {
			vmds[vmdAddr] = new(common.PCIAddressList)
		}

		vmds[vmdAddr].Add(*addr)
	}

	return vmds, nil
}

// substVMDAddrs replaces VMD PCI addresses in input device list with the PCI
// addresses of the backing devices behind the VMD.
//
// Return new device list with PCI addresses of devices behind the VMD.
func substVMDAddrs(inPCIAddrs *common.PCIAddressList, foundCtrlrs storage.NvmeControllers) (*common.PCIAddressList, error) {
	if len(foundCtrlrs) == 0 {
		return nil, nil
	}

	vmds, err := mapVMDToBackingAddrs(foundCtrlrs)
	if err != nil {
		return nil, err
	}

	// swap input vmd addresses with respective backing addresses
	outPCIAddrs := make(common.PCIAddressList, 0, len(*inPCIAddrs))
	for _, inAddr := range *inPCIAddrs {
		if backingAddrs, exists := vmds[*inAddr]; exists {
			outPCIAddrs = append(outPCIAddrs, *backingAddrs...)
			continue
		}
		outPCIAddrs = append(outPCIAddrs, inAddr)
	}

	return &outPCIAddrs, nil
}

// substituteVMDAddresses wraps around substVMDAddrs and takes a BdevScanResponse
// reference along with a logger.
func substituteVMDAddresses(log logging.Logger, inPCIAddrs *common.PCIAddressList, bdevCache *storage.BdevScanResponse) (*common.PCIAddressList, error) {
	if bdevCache == nil || len(bdevCache.Controllers) == 0 {
		log.Debugf("no bdev cache to find vmd backing devices (devs: %v)", inPCIAddrs)
		return nil, nil
	}

	msg := fmt.Sprintf("vmd detected, processing addresses (input %v, existing %v)",
		inPCIAddrs, bdevCache.Controllers)

	dl, err := substVMDAddrs(inPCIAddrs, bdevCache.Controllers)
	if err != nil {
		return nil, errors.Wrapf(err, msg)
	}
	log.Debugf("%s: new %s", msg, dl)

	return dl, nil
}

// detectVMD returns whether VMD devices have been found and a slice of VMD
// PCI addresses if found. Implements vmdDetectFn.
func detectVMD() (*common.PCIAddressList, error) {
	distro := system.GetDistribution()
	var lspciCmd *exec.Cmd

	// Check available VMD devices with command:
	// "$lspci | grep  -i -E "Volume Management Device"
	switch {
	case distro.ID == "opensuse-leap" || distro.ID == "opensuse" || distro.ID == "sles":
		lspciCmd = exec.Command("/sbin/lspci")
	default:
		lspciCmd = exec.Command("lspci")
	}

	vmdCmd := exec.Command("grep", "-i", "-E", "Volume Management Device")
	var cmdOut bytes.Buffer
	var prefixIncluded bool

	vmdCmd.Stdin, _ = lspciCmd.StdoutPipe()
	vmdCmd.Stdout = &cmdOut
	_ = lspciCmd.Start()
	_ = vmdCmd.Run()
	_ = lspciCmd.Wait()

	if cmdOut.Len() == 0 {
		return common.NewPCIAddressList()
	}

	vmdCount := bytes.Count(cmdOut.Bytes(), []byte("0000:"))
	if vmdCount == 0 {
		// sometimes the output may not include "0000:" prefix
		// usually when muliple devices are in PCI_ALLOWED
		vmdCount = bytes.Count(cmdOut.Bytes(), []byte("Volume"))
	} else {
		prefixIncluded = true
	}
	vmdAddrs := make([]string, 0, vmdCount)

	i := 0
	scanner := bufio.NewScanner(&cmdOut)
	for scanner.Scan() {
		if i == vmdCount {
			break
		}
		s := strings.Split(scanner.Text(), " ")
		if !prefixIncluded {
			s[0] = "0000:" + s[0]
		}
		vmdAddrs = append(vmdAddrs, strings.TrimSpace(s[0]))
		i++
	}

	if len(vmdAddrs) == 0 {
		return nil, errors.New("error parsing cmd output")
	}

	return common.NewPCIAddressList(vmdAddrs...)
}

// vmdFilterAddresses takes an input request and a list of discovered VMD addresses.
// The VMD addresses are validated against the input request allow and block lists.
// The output allow list will only contain VMD addresses if either both input allow
// and block lists are empty or if included in allow and not included in block lists.
func vmdFilterAddresses(inReq *storage.BdevPrepareRequest, vmdPCIAddrs *common.PCIAddressList) (*storage.BdevPrepareRequest, error) {
	outAllowList := new(common.PCIAddressList)
	outReq := *inReq

	inAllowList, err := common.NewPCIAddressListFromString(inReq.PCIAllowList)
	if err != nil {
		return nil, err
	}
	inBlockList, err := common.NewPCIAddressListFromString(inReq.PCIBlockList)
	if err != nil {
		return nil, err
	}

	// Set allow list to all VMD addresses if no allow or block lists in request.
	if inAllowList.IsEmpty() && inBlockList.IsEmpty() {
		outReq.PCIAllowList = vmdPCIAddrs.String()
		outReq.PCIBlockList = ""
		return &outReq, nil
	}

	// Add VMD addresses to output allow list if included in request allow list.
	if !inAllowList.IsEmpty() {
		inclAddrs := inAllowList.Intersect(vmdPCIAddrs)

		if inclAddrs.IsEmpty() {
			// no allowed vmd addresses
			outReq.PCIAllowList = ""
			outReq.PCIBlockList = ""
			return &outReq, nil
		}

		outAllowList = inclAddrs
	}

	if !inBlockList.IsEmpty() {
		// use outAllowList in case vmdPCIAddrs list has already been filtered
		inList := outAllowList

		if inList.IsEmpty() {
			inList = vmdPCIAddrs
		}

		exclAddrs := inList.Difference(inBlockList)

		if exclAddrs.IsEmpty() {
			// all vmd addresses are blocked
			outReq.PCIAllowList = ""
			outReq.PCIBlockList = ""
			return &outReq, nil
		}

		outAllowList = exclAddrs
	}

	outReq.PCIAllowList = outAllowList.String()
	outReq.PCIBlockList = ""
	return &outReq, nil
}

// getVMDPrepReq determines if VMD devices are going to be used and returns a
// bdev prepare request with the VMD addresses explicitly set in PCI_ALLOWED list.
//
// If VMD is not to be prepared, a nil request is returned.
func getVMDPrepReq(log logging.Logger, req *storage.BdevPrepareRequest, vmdDetect vmdDetectFn) (*storage.BdevPrepareRequest, error) {
	if !req.EnableVMD {
		return nil, nil
	}

	vmdPCIAddrs, err := vmdDetect()
	if err != nil {
		return nil, errors.Wrap(err, "VMD could not be enabled")
	}

	if vmdPCIAddrs.IsEmpty() {
		log.Debug("vmd prep: no vmd devices found")
		return nil, nil
	}
	log.Debugf("volume management devices detected: %v", vmdPCIAddrs)

	vmdReq, err := vmdFilterAddresses(req, vmdPCIAddrs)
	if err != nil {
		return nil, err
	}

	// No addrs left after filtering
	if vmdReq.PCIAllowList == "" {
		if req.PCIAllowList != "" {
			log.Debugf("vmd prep: %q devices not allowed", vmdPCIAddrs)
			return nil, nil
		}
		if req.PCIBlockList != "" {
			log.Debugf("vmd prep: %q devices blocked", vmdPCIAddrs)
			return nil, nil
		}
	}

	log.Debugf("volume management devices selected: %q", vmdReq.PCIAllowList)

	return vmdReq, nil
}
