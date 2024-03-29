package nuget

import (
	"fmt"
	"log"
	"package-downloader/service"
	"strings"
)

const NUGET_URL = "https://api.nuget.org/v3/registration5-gz-semver2"

func DownloadNuget(packageName string, packageVersion string, repository string) {
	name := strings.ToLower(packageName)
	pac := ToDo{name, packageVersion}
	packageToDownload, err := CheckDependency(pac)
	if err != nil {
		return
	}
	packageToDownload[pac] = true
	for pac, _ := range packageToDownload {
		//nexusItems, err := nexus.CheckNexus(packageName, packageVersion, repository)
		//if len(nexusItems) > 0 {
		//	log.Print("Package already in Nexus")
		//	return
		//}
		fileName := fmt.Sprintf("%s.%s.nupkg", pac.Name, pac.Version)
		nuget, err := GetNugetPackage(pac.Name, pac.Version)
		if err != nil {
			return
		}
		date := service.CheckDate(nuget.Published)
		if date == true {
			log.Printf(
				"Package: %s, version %s, published: %s",
				pac.Name,
				pac.Version,
				nuget.Published.Format("2006-01-02 15:04:05"),
			)
			err = service.DownloadFile(nuget.PackageContent, fileName)
			if err != nil {
				log.Printf("Can's save the file")
				return
			}
			log.Printf("Downloaded")
		} else {
			log.Printf(
				"Package: %s, version %s, published: %s",
				pac.Name,
				pac.Version,
				nuget.Published.Format("2006-01-02 15:04:05"),
			)
			log.Printf("Skip")
		}
	}
	log.Print("Done")
}
