package nuget

import (
	"encoding/json"
	"log"
	"os"
	"package-downloader/service"
	"strings"
	"time"
)

type MetaData struct {
	Id                       string             `json:"@id"`
	Created                  time.Time          `json:"created"`
	FrameworkReferences      string             `json:"frameworkReferences"`
	IconFile                 string             `json:"iconFile"`
	IdMain                   string             `json:"id"`
	IsPrerelease             bool               `json:"isPrerelease"`
	LastEdited               time.Time          `json:"lastEdited"`
	LicenseFile              string             `json:"licenseFile"`
	LicenseUrl               string             `json:"licenseUrl"`
	Listed                   bool               `json:"listed"`
	PackageHash              string             `json:"packageHash"`
	PackageHashAlgorithm     string             `json:"packageHashAlgorithm"`
	PackageSize              int                `json:"packageSize"`
	ProjectUrl               string             `json:"projectUrl"`
	Published                time.Time          `json:"published"`
	ReleaseNotes             string             `json:"releaseNotes"`
	Repository               string             `json:"repository"`
	RequireLicenseAcceptance bool               `json:"requireLicenseAcceptance"`
	VerbatimVersion          string             `json:"verbatimVersion"`
	Version                  string             `json:"version"`
	DependencyGroups         []DependencyGroups `json:"dependencyGroups"`
}

type DependencyGroups struct {
	Dependencies    []Dependencies `json:"dependencies"`
	TargetFramework string         `json:"targetFramework"`
}

type Dependencies struct {
	DependencyID string `json:"id"`
	Range        string `json:"range"`
}

type ToDo struct {
	Name    string
	Version string
}

var (
	p           = make(map[ToDo]bool)
	netStandart = os.Getenv("NET_STANDART")
)

func CheckDependency(pac ToDo) (map[ToDo]bool, error) {
	log.Printf("Processing package %s, %s ", pac.Name, pac.Version)
	data, err := findMetaData(pac.Name, pac.Version)
	if err != nil {
		log.Print("Can't find MetaData for package ")
		return nil, err
	}
	for _, group := range data.DependencyGroups {
		if group.TargetFramework == netStandart {
			for _, slice := range group.Dependencies {
				v, err := service.NormalizeVersion(slice.Range, "[][, )]", "${1}")
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
		return CheckDependency(dep)
	}
	return p, nil
}

func findMetaData(name string, version string) (*MetaData, error) {
	var metaData *MetaData

	pack, err := GetNugetPackage(name, version)

	body := service.GetRequest(pack.CatalogEntry)
	err = json.Unmarshal(body, &metaData)
	if err != nil {
		log.Print("Can't unmarsh json for Dependency")
		return nil, err
	}
	return metaData, err
}
