package npmClient

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (n *NpmPackage) GetPackage(packageName string) (*Package, error) {

	name := strings.ToLower(packageName)
	endpoint := fmt.Sprintf("/%s", name)
	apiURL := n.Client.NpmUrl + endpoint
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		log.Print("Can't unmarsh json for Nuget.Org")
	}
	res := Package{}
	if err := n.Client.SendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, err
}

func (n *NpmPackage) DownloadDependency(packageName string, packageVersion string, repository string) {
	name := strings.ToLower(packageName)
	pac := ToDo{name, packageVersion}
	packageToDownload, err := n.checkDependency(pac)
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
		fileName := fmt.Sprintf("%s-%s.tar", n.Client.NormalizeName(pac.Name), pac.Version)
		npm, err := n.GetPackage(pac.Name)
		if err != nil {
			return
		}
		date := n.Client.CheckDate(npm.Time[pac.Version])
		if date == true {
			log.Printf(
				"Package: %s, version %s, published: %s",
				pac.Name,
				pac.Version,
				npm.Time[pac.Version].Format("2006-01-02 15:04:05"),
			)
			err = n.Client.DownloadFile(npm.Versions[pac.Version].Dist.Tarball, fileName)
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

func (n *NpmPackage) checkDependency(pac ToDo) (map[ToDo]bool, error) {

	var p = make(map[ToDo]bool)

	log.Printf("Processing package %s, %s ", pac.Name, pac.Version)
	npmPac, _ := n.GetPackage(pac.Name)
	for name, vers := range npmPac.Versions[pac.Version].Dependencies {
		log.Printf("%s %s", name, vers)
		v, err := n.Client.NormalizeVersion(vers, "[/^|~]", "")
		if err != nil {
			log.Print("Can't normalize version")
			break
		}
		newDep := ToDo{Name: name, Version: v}
		p[newDep] = false
	}
	for pac, load := range p {
		if load == true {
			continue
		}
		p[pac] = true
		return n.checkDependency(pac)
	}
	return p, nil
}
