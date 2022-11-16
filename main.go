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
		npmPac, _ := npm.GetPackage(packageName, packageVersion)
		npmDep, _ := npm.GetDependencies(npmPac)
		log.Print(npmDep)
	default:
		fmt.Print("here will be a default downloader section")
	}
}
