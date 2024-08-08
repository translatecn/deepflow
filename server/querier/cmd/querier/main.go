/*
 * Copyright (c) 2024 Yunshan Networks
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"github.com/deepflowio/deepflow/server/libs/over_logger"
	"os"
	"strings"

	logging "github.com/op/go-logging"

	"github.com/deepflowio/deepflow/server/common"
	"github.com/deepflowio/deepflow/server/querier/over_querier"
)

var configPath = flag.String("f", "/etc/server.yaml", "Specify config file location")

func execName() string {
	splitted := strings.Split(os.Args[0], "/")
	return splitted[len(splitted)-1]
}

var log = logging.MustGetLogger(execName())

func main() {
	if os.Getppid() != 1 {
		over_logger.EnableStdoutLog()
	}
	shared := common.NewControllerIngesterShared()
	over_querier.Start(*configPath, "", shared)
}
