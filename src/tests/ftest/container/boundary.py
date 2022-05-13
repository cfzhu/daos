#!/usr/bin/python3
"""
  (C) Copyright 2022 Intel Corporation.

  SPDX-License-Identifier: BSD-2-Clause-Patent
"""

import time
from apricot import TestWithServers
from general_utils import DaosTestError
from avocado.core.exceptions import TestFail
from thread_manager import ThreadManager

class BoundaryTest(TestWithServers):
    """
    Epic: Create system level tests that cover boundary tests and
          functionality.
    Testcase:
          DAOS-8464: Test lots of pools and connections
    Test Class Description:
          Start DAOS servers, create pools and containers to the support limit.
    :avocado: recursive
    """

    def __init__(self, *args, **kwargs):
        """Initialize a BoundaryTest object."""
        super().__init__(*args, **kwargs)
        self.with_io = False
        self.io_run_time = None
        self.io_rank = None
        self.io_obj_classs = None

    def setUp(self):
        """Set Up BoundaryTest"""
        super().setUp()
        self.pool = []
        self.with_io = self.params.get("with_io", '/run/boundary_test/*')
        self.io_run_time = self.params.get("run_time", '/run/container/execute_io/*')
        self.io_rank = self.params.get("rank", '/run/container/execute_io/*')
        self.io_obj_classs = self.params.get("obj_classs", '/run/container/execute_io/*')

    def get_pool(self, *args, **kwargs):
        """Get a test pool object and append to list.

        Returns:
            TestPool: the created test pool object.

        """
        pool = super().get_pool(*args, **kwargs)
        self.pool.append(pool)
        return pool

    def create_container_and_test(self, pool=None, cont_num=1):
        """To create single container on pool.

        Args:
            pool (str): pool handle to create container.
            container_num (int): container number to create.

        """
        try:
            container = self.get_container(pool)
        except (DaosTestError, TestFail) as err:
            msg = "#(3.{}.{}) container create failed. err={}".format(pool.label, cont_num, err)
            self.fail(msg)

        self.log.info("===(3.%s.%d)create_container_and_test, container %s created..",
            pool.label, cont_num, container)
        if self.with_io:
            try:
                data_bytes = container.execute_io(
                    self.io_run_time, self.io_rank, self.io_obj_classs)
                self.log.info(
                    "===(3.%s.%d)Wrote %d bytes to container %s", pool.label, cont_num,
                    data_bytes, container)
            except (DaosTestError, TestFail) as err:
                msg = "#(3.{}.{}) container IO failed, err: {}".format(pool.label, cont_num, err)
                self.fail(msg)
        time.sleep(2)  #to syncup containers before close

        try:
            self.log.info("===(4.%s.%d)create_container_and_test, container closing.",
                pool.label, cont_num)
            container.close()
            self.log.info("===(4.%s.%d)create_container_and_test, container closed.",
                pool.label, cont_num)
        except (DaosTestError, TestFail) as err:
            msg = "#(4.{}.{}) container close fail, err: {}".format(pool.label, cont_num, err)
            self.fail(msg)

    def create_containers(self, pool, num_containers):
        """To create number of containers parallelly on pool.

        Args:
            pool(str): pool handle.
            num_containers (int): number of containers to create.
        """
        self.log.info("==(2.%s)create_containers start.", pool.label)
        thread_manager = ThreadManager(self.create_container_and_test, self.timeout - 30)

        for cont_num in range(num_containers):
            thread_manager.add(pool=pool, cont_num=cont_num)

        # Launch the create_container_and_test threads
        self.log.info("==Launching %d create_container_and_test threads", thread_manager.qty)
        failed_thread_count = thread_manager.check_run()
        self.log.info(
            "==(2.%s) after thread_manager_run, %d containers created.", pool.label, num_containers)
        if failed_thread_count > 0:
            msg = "#(2.{}) FAILED create_container_and_test Threads".format(failed_thread_count)
            self.d_log.error(msg)
            self.fail(msg)

    def create_pools(self, num_pools, num_containers):
        """To create number of pools and containers parallelly.

        Args:
            num_pools (int): number of pools to create.
            num_containers (int): number of containers to create.

        """
        # Create pools in parallel
        pool_manager = ThreadManager(self.get_pool, self.timeout - 30)
        for _ in range(num_pools):
            pool_manager.add()
        self.log.info('Creating %d pools', num_pools)
        num_failed = pool_manager.check_run()
        if num_failed > 0:
            self.fail('{} pool create threads failed'.format(num_failed))
        self.log.info('Created %d pools', num_pools)

        # Create containers for each pool in parallel
        container_manager = ThreadManager(self.create_containers, self.timeout - 30)
        for pool in self.pool:
            container_manager.add(pool=pool, num_containers=num_containers)
        self.log.info('Creating %d containers for each pool', num_containers)
        num_failed = container_manager.check_run()
        if num_failed > 0:
            self.fail('{} container create threads failed'.format(num_failed))
        self.log.info('Created %d * %d containers', num_pools, num_containers)

    def test_container_boundary(self):
        """JIRA ID: DAOS-8464 Test lots of pools and containers in parallel.
        Test Description:
            Testcase 1: Test 1 pool with containers boundary condition in parallel.
            Testcase 2: Test large number of pools and containers in parallel.
            Testcase 3: Test pools and containers with io.
            log.info: (a.b.c) a: test-step,  b: pool.label,  c: container_number
        Use case:
            0. Bring up DAOS server.
            1. Create pools and create containers_test by ThreadManager.
            2. Create containers and test under each pool by sub ThreadManager.
            3. Launch io and syncup each container.
            4. Close container.
        :avocado: tags=all,full_regression
        :avocado: tags=hw,medium
        :avocado: tags=container,pool
        :avocado: tags=container_boundary,pool_boundary
        """
        num_pools = self.params.get("num_pools", '/run/boundary_test/*')
        num_containers = self.params.get("num_containers", '/run/boundary_test/*')
        self.create_pools(num_pools=num_pools, num_containers=num_containers)
        self.log.info("===>Boundary test passed.")
