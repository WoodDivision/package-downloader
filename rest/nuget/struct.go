package nuget

import (
	"package-downloader/rest"
	"time"
)

type NugetPackage struct {
	Client *rest.Client
}

type Package struct {
	Id             string    `json:"@id"`
	CatalogEntry   string    `json:"catalogEntry"`
	PackageContent string    `json:"packageContent"`
	Published      time.Time `json:"published"`
}

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
