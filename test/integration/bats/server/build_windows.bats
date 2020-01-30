#!/usr/bin/env bats
# Copyright 2017-2020 The Usackoud Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


load ${BASE_TEST_DIR}/helpers.bash

password="$TMP_PASSWORD"
hostname=${TEST_TARGET_NAME}01

@test "Usacloud: should be able to build server with using Windows 2016 RDS" {

  run usacloud_run_with_stderr server build -y -q \
          --os-type windows2016-rds \
          --password "$password" \
          --disk-size 100 \
          --name "$hostname"

  [ -n "${output}" ]
  [ ${status} -eq 1 ]
}

