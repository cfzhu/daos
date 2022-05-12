#!/usr/bin/python
"""
  (C) Copyright 2020-2022 Intel Corporation.

  SPDX-License-Identifier: BSD-2-Clause-Patent
"""

import time
import os
from ior_test_base import IorTestBase


class DfuseSpaceCheck(IorTestBase):
    # pylint: disable=too-many-ancestors
    """Base Parallel IO test class.

    :avocado: recursive
    """

    def __init__(self, *args, **kwargs):
        """Initialize a DfuseSpaceCheck object."""
        super().__init__(*args, **kwargs)
        self.initial_space = None
        self.block_size = None

    def get_nvme_free_space(self, display=True):
        """Display pool free space.

        Args:
            display (bool): boolean to display output of free space.

        Returns:
            int: Free space available in nvme.

        """
        free_space_nvme = self.pool.get_pool_free_space("nvme")
        if display:
            self.log.info("Free nvme space: %s", free_space_nvme)

        return free_space_nvme

    def wait_for_aggregation(self, retries=4, interval=60):
        """Wait for aggregation to finish.

        Args:
            retries (int, optional): number of times to retry.
                Default is 4.
            interval (int, optional): seconds to wait before retrying.
                Default is 60.

        """
        for _ in range(retries):
            current_space = self.get_nvme_free_space()
            if current_space == self.initial_space:
                return
            time.sleep(interval)

        self.log.info("Free space when test terminated: %s", current_space)
        self.fail("Aggregation did not complete within {} seconds".format(retries * interval))

    def write_multiple_files(self):
        """Write multiple files.

        Returns:
            int: Total number of files created before going out of space.

        """
        file_count = 0
        while self.get_nvme_free_space(False) >= self.block_size:
            file_loc = str(self.dfuse.mount_dir.value +
                           "/largefile_{}.txt".format(file_count))
            write_dd_cmd = "dd if=/dev/zero of={} bs={} count=1".format(
                file_loc, self.block_size)
            if 0 in self.execute_cmd(write_dd_cmd, False, False):
                file_count += 1

        return file_count

    def test_dfusespacecheck(self):
        """Jira ID: DAOS-3777.

        Test Description:
            Purpose of this test is to mount dfuse and verify aggregation
            to return space when pool is filled with once large file and
            once with small files.
        Use cases:
            Create Pool
            Create Posix container
            Mount dfuse
            Create largefile.txt and write to it until pool is out of space.
            Remove largefile.txt and wait for aggregation to return all the
            space back.
            Now create many small files until pool is out of space again and
            store the number of files created.
            Remove all the small files created and wait for aggregation to
            return all the space back.
            Now create small files again until pool is out of space and check,
            whether out of space happens at the same file count as before.
        :avocado: tags=all,full_regression
        :avocado: tags=hw,small
        :avocado: tags=aggregation,daosio,dfuse
        :avocado: tags=dfusespacecheck
        """
        # get test params for cont and pool count
        self.block_size = self.params.get('block_size', '/run/dfusespacecheck/*')

        # Create a pool, container, and start dfuse
        self.create_pool()
        self.create_cont()
        self.start_dfuse(self.hostlist_clients, self.pool, self.container)

        # get nvme space before write
        self.initial_space = self.get_nvme_free_space()

        # Create a file as large as we can
        large_file = os.path.join(self.dfuse.mount_dir.value, 'largefile.txt')
        self.execute_cmd('touch {}'.format(large_file))
        dd_count = (self.initial_space // self.block_size) + 1
        write_dd_cmd = "dd if=/dev/zero of={} bs={} count={}".format(
            large_file, self.block_size, dd_count)
        self.execute_cmd(write_dd_cmd, False)

        # Remove the file
        self.execute_cmd('rm -rf {}'.format(large_file))

        # Wait for aggregation to complete
        self.wait_for_aggregation()

        # Write small files until we run out of space
        file_count1 = self.write_multiple_files()

        # remove all the small files created above.
        self.execute_cmd("rm -rf {}".format(os.path.join(self.dfuse.mount_dir.value, '*')))

        # Wait for aggregation to complete after file removal
        self.wait_for_aggregation()

        # Write small files again until we run out of space
        # and verify that we wrote at least as many as file_count1
        file_count2 = self.write_multiple_files()
        self.log.info('file_count1 = %s', file_count1)
        self.log.info('file_count2 = %s', file_count2)
        self.assertGreaterEqual(
            file_count2, file_count1,
            'Space was not returned. Expected to write at least {} files'.format(file_count1))
