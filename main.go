//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"os"

	"github.com/edgexfoundry/app-functions-sdk-go/appsdk"
	"github.com/edgexfoundry/app-functions-sdk-go/pkg/transforms"

	"app-send-cmd/functions"
)

const (
	serviceKey = "httpCommandService"
)

func main() {
	// turn off secure mode for examples. Not recommended for production
	os.Setenv("EDGEX_SECURITY_SECRET_STORE", "false")
	// 1) First thing to do is to create an instance of the EdgeX SDK with TargetType Switch
	//    and initialize it.
	edgexSdk := &appsdk.AppFunctionsSDK{ServiceKey: serviceKey, TargetType: &functions.Switch{}}
	if err := edgexSdk.Initialize(); err != nil {
		edgexSdk.LoggingClient.Error(fmt.Sprintf("SDK initialization failed: %v\n", err))
		os.Exit(-1)
	}

	// 2) This is our functions pipeline configuration, the collection of functions to
	// execute every time an event is triggered.
	err := edgexSdk.SetFunctionsPipeline(
		functions.SendSwitchCommand,
		functions.PrintToConsole,
		transforms.NewOutputData().SetOutputData,
	)

	if err != nil {
		edgexSdk.LoggingClient.Error("Setting Functions Pipeline failed: " + err.Error())
		os.Exit(-1)
	}

	// 3) Lastly, we'll go ahead and tell the SDK to "start" and begin listening for Persons
	// to trigger the pipeline.
	err = edgexSdk.MakeItRun()
	if err != nil {
		edgexSdk.LoggingClient.Error("MakeItRun returned error: ", err.Error())
		os.Exit(-1)
	}

	// Do any required cleanup here
	os.Exit(0)
}
