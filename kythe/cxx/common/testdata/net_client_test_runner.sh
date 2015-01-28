#!/bin/bash -e
# Copyright 2015 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
set -o pipefail
BASE_DIR="${PWD}/kythe/cxx/common/testdata"
TEST_BIN="${PWD}/campfire-out/bin/kythe/cxx/common/net_client_test"
KYTHE_WRITE_ENTRIES="${PWD}/campfire-out/bin/kythe/go/storage/tools/write_entries"
KYTHE_ENTRYSTREAM="${PWD}/campfire-out/bin/kythe/go/platform/tools/entrystream"
KYTHE_HTTP_SERVER="${PWD}/campfire-out/bin/kythe/go/serving/tools/http_server"
OUT_DIR="${PWD}/campfire-out/test/kythe/cxx/common/testdata/net_client_test_runner"

rm -rf -- "${OUT_DIR}"
mkdir -p "${OUT_DIR}/gs"

cat "${BASE_DIR}/net_client_test_data.json" \
    | "${KYTHE_ENTRYSTREAM}" -read_json=true  \
    | "${KYTHE_WRITE_ENTRIES}" -graphstore "${OUT_DIR}/gs"
"${KYTHE_HTTP_SERVER}" -graphstore "${OUT_DIR}/gs" -listen="localhost:0" &
SERVER_PID=$!
trap 'kill $SERVER_PID' EXIT ERR INT
LISTEN_AT=$((lsof -p "${SERVER_PID}" | grep -ohw "localhost:[0-9]*") || true)
while ! curl -s "${LISTEN_AT}" > /dev/null; do
  echo 'Waiting for server...' >&2
  sleep 1s
  LISTEN_AT=$((lsof -p "${SERVER_PID}" | grep -ohw "localhost:[0-9]*") || true)
done

"${TEST_BIN}" --xrefs="http://${LISTEN_AT}"
