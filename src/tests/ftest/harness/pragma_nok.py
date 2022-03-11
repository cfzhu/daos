#!/usr/bin/python3
"""
  (C) Copyright 2021-2022 Intel Corporation.

  SPDX-License-Identifier: BSD-2-Clause-Patent
"""
from apricot import Test


class PragmaNok(Test):
    """Dummy test to illustrate pragma issue filtering with python3 features

    :avocado: recursive
    """

    def test_bar(self):
        """Simple test of apricot test code.

        :avocado: tags=all
        :avocado: tags=vm
        :avocado: tags=harness,pragma
        :avocado: tags=pragma_bar
        """

        bar="hello bar world"
        self.log.info(f"bar={bar}")
