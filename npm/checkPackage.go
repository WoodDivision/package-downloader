package npm

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Package struct {
	Id       string             `json:"_id"`
	Rev      string             `json:"_rev"`
	Name     string             `json:"name"`
	DistTags struct{}           `json:"dist-tags"`
	Versions map[string]Version `json:"versions"`
	Time     map[string]string  `json:"time"`
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

var npm *Package

const NPM_URL = "https://registry.npmjs.org"

func GetPackage(packageName string) (*Package, error) {
	name := strings.ToLower(packageName)
	//version := packageVersion
	endpoint := fmt.Sprintf("/%s", name)
	apiURL := NPM_URL + endpoint
	resp, err := http.Get(apiURL)
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
	err = json.Unmarshal(body, &npm)
	if err != nil {
		log.Print("Can't unmarsh json for registry.npmjs.org")
		log.Print(err.Error())
		return nil, err
	}
	return npm, nil
}
