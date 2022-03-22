APP="./install/lib/daos/TESTING/tests/dual_provider_client" 

HOST="wolf-55"

# These envariables will be used only by crt_launch
export OFI_INTERFACE=eth0
export OFI_DOMAIN=eth0
export CRT_PHY_ADDR_STR="ofi+sockets"

export OFI_INTERFACE=ib0
export OFI_DOMAIN=mlx5_0
export CRT_PHY_ADDR_STR="ofi+tcp;ofi_rxm"

ORTE_EXORTS="-x OFI_INTERFACE -x OFI_DOMAIN -x CRT_PHY_ADDR_STR"
set -x
${APP}
