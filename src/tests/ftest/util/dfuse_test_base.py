#!/usr/bin/python
"""
(C) Copyright 2020-2022 Intel Corporation.

SPDX-License-Identifier: BSD-2-Clause-Patent
"""
import agent_utils as agu
from ClusterShell.NodeSet import NodeSet

from apricot import TestWithServers
from exception_utils import CommandFailure
from dfuse_utils import Dfuse


class DfuseTestBase(TestWithServers):
    """Runs Dfuse test suites.

    :avocado: recursive
    """

    def __init__(self, *args, **kwargs):
        """Initialize a TestWithServers object."""
        super().__init__(*args, **kwargs)
        self.dfuse = None

    def setUp(self):
        """Setup Test Case"""
        super().setUp()
        # using localhost as client if client list is empty
        if self.hostlist_clients is None:
            self.hostlist_clients = agu.include_local_host(None)

    def stop_job_managers(self):
        """Stop the test job manager followed by dfuse.

        Returns:
            list: a list of exceptions raised stopping the agents

        """
        error_list = super().stop_job_managers()
        try:
            self.stop_dfuse()
        except CommandFailure as error:
            error_list.append("Error stopping dfuse: {}".format(error))
        return error_list

    def start_dfuse(self, hosts, pool=None, container=None, mount_dir=None):
        """Create a DfuseCommand object and use it to start Dfuse.

        Args:
            hosts (list): list of hosts on which to start Dfuse
            pool (TestPool, optional): pool to use with Dfuse
            container (TestContainer, optional): container to use with Dfuse
            mount_dir (str, optional): updated mount dir name. Defaults to None.
        """
        self.dfuse = Dfuse(hosts, self.tmp)
        self.dfuse.get_params(self)

        dfuse_cores = self.params.get('cores', self.dfuse.namespace, None)

        # Update dfuse params
        if mount_dir:
            self.dfuse.mount_dir.update(mount_dir)
        if pool:
            self.dfuse.set_dfuse_params(pool)
        if container:
            self.dfuse.set_dfuse_cont_param(container)
        self.dfuse.set_dfuse_exports(self.server_managers[0], self.client_log)

        try:
            # Start dfuse
            self.dfuse.run(bind_cores=dfuse_cores)
        except CommandFailure as error:
            self.log.error(
                "Dfuse command %s failed on hosts %s", str(self.dfuse),
                str(NodeSet.fromlist(self.dfuse.hosts)), exc_info=error)
            self.fail("Test was expected to pass but it failed.")

    def stop_dfuse(self):
        """Stop Dfuse and unset the DfuseCommand object."""
        if self.dfuse:
            self.dfuse.stop()
            self.dfuse = None
