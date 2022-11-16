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
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Main            string            `json:"main"`
	Scripts         map[string]string `json:"scripts"`
	Repository      Repository        `json:"repository"`
	Engines         Engines           `json:"engines"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	GitHead         string            `json:"gitHead"`
	Homepage        string            `json:"homepage"`
	Id              string            `json:"_id"`
	NodeVersion     string            `json:"_nodeVersion"`
	NpmVersion      string            `json:"_npmVersion"`
	Dist            Dist              `json:"dist"`
}

type Repository struct {
	Type      string `json:"type"`
	Url       string `json:"url"`
	Directory string `json:"directory"`
}

type Engines struct {
	Node string `json:"node"`
}
type Dist struct {
	Integrity    string       `json:"integrity"`
	Shasum       string       `json:"shasum"`
	Tarball      string       `json:"tarball"`
	FileCount    int          `json:"fileCount"`
	UnpackedSize int          `json:"unpackedSize"`
	Signatures   []Signatures `json:"signatures"`
	NpmSignature string       `json:"npm-signature"`
}

type Signatures struct {
	Keyid string `json:"keyid"`
	Sig   string `json:"sig"`
}

var npm *Package

const NPM_URL = "https://registry.npmjs.org"

func GetPackage(packageName string, packageVersion string) (*Package, error) {
	name := strings.ToLower(packageName)
	version := packageVersion
	endpoint := fmt.Sprintf("/%s/%s", name, version)
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
