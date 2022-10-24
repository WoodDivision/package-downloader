package nuget

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type MetaData struct {
	Id                       string             `json:"@id"`
	Type                     []string           `json:"@type"`
	Authors                  string             `json:"authors"`
	CatalogCommitId          string             `json:"catalog:commitId"`
	CatalogCommitTimeStamp   time.Time          `json:"catalog:commitTimeStamp"`
	Copyright                string             `json:"copyright"`
	Created                  time.Time          `json:"created"`
	Description              string             `json:"description"`
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
	PackageEntries           []PackageEntries   `json:"packageEntries"`
	Tags                     []string           `json:"tags"`
	PackageContext           PackageContext     `json:"@context"`
}
type DependencyGroups struct {
	Id              string         `json:"@id"`
	Type            string         `json:"@type"`
	Dependencies    []Dependencies `json:"dependencies"`
	TargetFramework string         `json:"targetFramework"`
}

type Dependencies struct {
	Id           string `json:"@id"`
	Type         string `json:"@type"`
	DependencyID string `json:"id"`
	Range        string `json:"range"`
}
type PackageEntries struct {
	Id               string `json:"@id"`
	Type             string `json:"@type"`
	CompressedLength int    `json:"compressedLength"`
	FullName         string `json:"fullName"`
	Length           int    `json:"length"`
	Name             string `json:"name"`
}
type PackageContext struct {
	Vocab                   string                  `json:"@vocab"`
	Catalog                 string                  `json:"catalog"`
	Xsd                     string                  `json:"xsd"`
	DependenciesContext     DependenciesContext     `json:"dependencies"`
	DependencyGroupsContext DependencyGroupsContext `json:"dependencyGroups"`
	PackageEntriesContext   PackageEntriesContext   `json:"packageEntries"`
	PackageTypes            PackageTypes            `json:"packageTypes"`
	SupportedFrameworks     SupportedFrameworks     `json:"supportedFrameworks"`
	Tags                    Tags                    `json:"tags"`
	Vulnerabilities         Vulnerabilities         `json:"vulnerabilities"`
	PublishedType           PublishedType           `json:"published"`
	Created                 Created                 `json:"created"`
	LastEdited              LastEdited              `json:"lastEdited"`
	CatalogCommitTimeStamp  CatalogCommitTimeStamp  `json:"catalog:commitTimeStamp"`
	Reasons                 Reasons                 `json:"reasons"`
}

type DependenciesContext struct {
	Id        string `json:"@id"`
	Container string `json:"@container"`
}
type DependencyGroupsContext struct {
	Id        string `json:"@id"`
	Container string `json:"@container"`
}
type PackageEntriesContext struct {
	Id        string `json:"@id"`
	Container string `json:"@container"`
}
type PackageTypes struct {
	Id        string `json:"@id"`
	Container string `json:"@container"`
}
type SupportedFrameworks struct {
	Id        string `json:"@id"`
	Container string `json:"@container"`
}

type Tags struct {
	Id        string `json:"@id"`
	Container string `json:"@container"`
}
type Vulnerabilities struct {
	Id        string `json:"@id"`
	Container string `json:"@container"`
}
type PublishedType struct {
	Type string `json:"@type"`
}
type Created struct {
	Type string `json:"@type"`
}
type LastEdited struct {
	Type string `json:"@type"`
}

type CatalogCommitTimeStamp struct {
	Type string `json:"@type"`
}
type Reasons struct {
	Container string `json:"@container"`
}

type ToDo struct {
	Name    string
	Version string
}

var (
	p        = make(map[ToDo]bool)
	metaData *MetaData
)

func FindDependencies(pac ToDo) (map[ToDo]bool, error) {
	data, err := findMetaData(pac.Name, pac.Version)
	if err != nil {
		log.Print("Can't find MetaData for package ")
		return nil, err
	}
	for _, target := range data.DependencyGroups {
		if len(target.Dependencies) != 0 {
			for _, slice := range target.Dependencies {
				reg := regexp.MustCompile("[][, )]")
				depVersion := reg.ReplaceAllString(slice.Range, "${1}")
				depName := strings.ToLower(slice.DependencyID)
				newDep := ToDo{depName, depVersion}
				p[newDep] = false
			}
		}
	}
	for pac, load := range p {
		if load == true {
			break
		}
		p[pac] = true
		return FindDependencies(pac)
	}
	return p, nil
}

func findMetaData(name string, version string) (*MetaData, error) {

	pack, err := FindPackage(name, version)
	resp, err := http.Get(pack.CatalogEntry)
	if err != nil {
		log.Print("Wrong request.Check that you enter correct package name or version ")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print("There is error in body pars")
		return nil, err
	}
	err = json.Unmarshal(body, &metaData)
	if err != nil {
		log.Print("Can't unmarsh json for Dependency")
		return nil, err
	}
	return metaData, err
}
