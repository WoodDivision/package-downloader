package nuget

import (
	"encoding/json"
	"fmt"
	"log"
	"package-downloader/service"
	"strings"
	"time"
)

type Package struct {
	Id             string    `json:"@id"`
	CatalogEntry   string    `json:"catalogEntry"`
	PackageContent string    `json:"packageContent"`
	Published      time.Time `json:"published"`
}

var nuget *Package

func GetNugetPackage(packageName string, packageVersion string) (*Package, error) {
	name := strings.ToLower(packageName)
	version := packageVersion
	endpoint := fmt.Sprintf("/%s/%s.json", name, version)
	apiURL := NUGET_URL + endpoint
	body := service.GetRequest(apiURL)
	err := json.Unmarshal(body, &nuget)
	if err != nil {
		log.Print("Can't unmarsh json for Nuget.Org")
		return nil, err
	}

	return nuget, nil
}
