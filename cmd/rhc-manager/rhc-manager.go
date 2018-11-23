/*
   Copyright 2018 Sysdig

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"github.com/tembleking/rhc-manager"
	"fmt"
	"flag"
)

var (
	flagId    = flag.String("id", "", "Project ID")
	flagBuild = flag.String("b", "", "Build using the specified Tag")
)

func main() {
	flag.Parse()

	client := rhc_manager.ApiClient{}

	if *flagId == "" {
		fmt.Println("Project Id must be specified using -id. See -h for more information")
		return
	}

	switch {
	case *flagBuild != "":
		response, err := client.BuildProject(*flagId, *flagBuild)
		if err != nil {
			panic(fmt.Sprintf("%+v", err))
		}
		fmt.Println(response.Message)
	default:
		project, err := client.GetProject(*flagId)
		if err != nil {
			panic(fmt.Sprintf("%+v", err))
		}
		fmt.Println(project.Project)
	}

	return

}
