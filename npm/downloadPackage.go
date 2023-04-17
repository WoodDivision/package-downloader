package npm

import (
	"fmt"
	"log"
	"package-downloader/service"
	"strings"
)

const NPM_URL = "https://registry.npmjs.org"

func DownloadNpm(packageName string, packageVersion string, repository string) {
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
		fileName := fmt.Sprintf("%s-%s.tar", service.NormalizeName(pac.Name), pac.Version)
		npm, err := GetNpmPackage(pac.Name)
		if err != nil {
			return
		}
		date := service.CheckDate(npm.Time[pac.Version])
		if date == true {
			log.Printf(
				"Package: %s, version %s, published: %s",
				pac.Name,
				pac.Version,
				npm.Time[pac.Version].Format("2006-01-02 15:04:05"),
			)
			err = service.DownloadFile(npm.Versions[pac.Version].Dist.Tarball, fileName)
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
				npm.Time[pac.Version].Format("2006-01-02 15:04:05"),
			)
			log.Printf("Skip")
		}
	}
	log.Print("Done")
}
