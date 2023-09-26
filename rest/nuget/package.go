package nuget

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

var (
	p           = make(map[ToDo]bool)
	netStandart = os.Getenv("NET_STANDART")
)

func (n *NugetPackage) GetPackage(packageName string, packageVersion string) (*Package, error) {
	name := strings.ToLower(packageName)
	version := packageVersion
	endpoint := fmt.Sprintf("/%s/%s.json", name, version)
	api, err := url.Parse(n.Client.NugetUrl + endpoint)
	if err != nil {
		log.Printf("Wrong url for nuget.org")
		return nil, err
	}
	res := Package{}
	if err := n.Client.SendRequest("GET", api, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (n *NugetPackage) findMetaData(name string, version string) (*MetaData, error) {
	pack, err := n.GetPackage(name, version)
	api, err := url.Parse(pack.CatalogEntry)
	if err != nil {
		log.Printf("Wrong catalog entry in package")
		return nil, err
	}
	res := MetaData{}
	if err := n.Client.SendRequest("GET", api, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (n *NugetPackage) checkDependency(pac ToDo) (map[ToDo]bool, error) {
	log.Printf("Processing package %s, %s ", pac.Name, pac.Version)
	data, err := n.findMetaData(pac.Name, pac.Version)
	if err != nil {
		log.Print("Can't find MetaData for package ")
		return nil, err
	}
	for _, group := range data.DependencyGroups {
		if group.TargetFramework == netStandart {
			for _, slice := range group.Dependencies {
				v, err := n.Client.NormalizeVersion(slice.Range, "[][, )]", "${1}")
				if err != nil {
					log.Print("Can't normalize version")
					break
				}
				depName := strings.ToLower(slice.DependencyID)
				newDep := ToDo{depName, v}
				p[newDep] = false
			}
		}
	}
	for dep, load := range p {
		if load == true {
			continue
		}
		p[dep] = true
		return n.checkDependency(dep)
	}
	return p, nil
}

func (n *NugetPackage) DownloadDependency(packageName string, packageVersion string) {
	name := strings.ToLower(packageName)
	pac := ToDo{name, packageVersion}
	packageToDownload, err := n.checkDependency(pac)
	if err != nil {
		return
	}
	packageToDownload[pac] = true
	for do := range packageToDownload {
		fileName := fmt.Sprintf("%s.%s.nupkg", do.Name, do.Version)
		nuget, err := n.GetPackage(do.Name, do.Version)
		if err != nil {
			return
		}
		date := n.Client.CheckDate(nuget.Published)
		if date == true {
			log.Printf(
				"Package: %s, version %s, published: %s",
				do.Name,
				do.Version,
				nuget.Published.Format("2006-01-02 15:04:05"),
			)
			err = n.Client.DownloadFile(nuget.PackageContent, fileName)
			if err != nil {
				log.Printf("Can's save the file")
				return
			}
			log.Printf("Downloaded")
		} else {
			log.Printf(
				"Package: %s, version %s, published: %s",
				do.Name,
				do.Version,
				nuget.Published.Format("2006-01-02 15:04:05"),
			)
			log.Printf("Skip")
		}
	}
	log.Print("Done")
}
