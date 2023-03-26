package main

import (
	"flag"
	"fmt"
	"log"
	"nuget-downloader/npm"
	"nuget-downloader/nuget"
	"os"
)

var (
	packageName    string
	packageVersion string
	repository     string
	packageType    string
)

func main() {
	log.SetOutput(os.Stdout)
	flag.StringVar(&packageName, "n", "", "Package name")
	flag.StringVar(&packageVersion, "v", "", "Package version")
	flag.StringVar(&packageType, "t", "", "Package type")
	flag.StringVar(&repository, "r", "nuget-freeze", "Package repository name in Nexus.Action")
	flag.Parse()
	switch packageType {
	case "nuget":
		nuget.DownloadNuget(packageName, packageVersion, repository)
	case "composer":
		fmt.Print("here will be a composer package downloader section")
	case "npm":
		pac := npm.ToDo{Name: packageName, Version: packageVersion}
		npmDep, _ := npm.CheckDependency(pac)
		//log.Printf("%v", npmPac.Versions[packageVersion].Dependencies)
		log.Printf("%v", npmDep)
		log.Printf("%s", len(npmDep))
	default:
		fmt.Print("here will be a default downloader section")
	}
}
