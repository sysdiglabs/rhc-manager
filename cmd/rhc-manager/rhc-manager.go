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
