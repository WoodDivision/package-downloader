package npmClient

import (
	"package-downloader/rest"
	"time"
)

type NpmPackage struct {
	Client *rest.Client
}

type Package struct {
	Id       string               `json:"_id"`
	Rev      string               `json:"_rev"`
	Name     string               `json:"name"`
	DistTags struct{}             `json:"dist-tags"`
	Versions map[string]Version   `json:"versions"`
	Time     map[string]time.Time `json:"time"`
}

type Version struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Id              string            `json:"_id"`
	NodeVersion     string            `json:"_nodeVersion"`
	NpmVersion      string            `json:"_npmVersion"`
	Dist            Dist              `json:"dist"`
}

type Dist struct {
	Shasum     string `json:"shasum"`
	Tarball    string `json:"tarball"`
	Integrity  string `json:"integrity"`
	Signatures []struct {
		Keyid string `json:"keyid"`
		Sig   string `json:"sig"`
	} `json:"signatures"`
}

type ToDo struct {
	Name    string
	Version string
}
