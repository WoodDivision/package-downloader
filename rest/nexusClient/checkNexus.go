package nexusClient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Items             Item        `json:"items"`
	ContinuationToken interface{} `json:"continuationToken"`
}

type Item []struct {
	Id         string      `json:"id"`
	Repository string      `json:"repository"`
	Format     string      `json:"format"`
	Group      interface{} `json:"group"`
	Name       string      `json:"name"`
	Version    string      `json:"version"`
	Assets     Assets      `json:"assets"`
}

type Assets []struct {
	DownloadUrl  string    `json:"downloadUrl"`
	Path         string    `json:"path"`
	Id           string    `json:"id"`
	Repository   string    `json:"repository"`
	Format       string    `json:"format"`
	Checksum     Checksum  `json:"checksum"`
	ContentType  string    `json:"contentType"`
	LastModified time.Time `json:"lastModified"`
	Nuget        Nuget     `json:"nuget"`
}

type Checksum struct {
	Sha1   string `json:"sha1"`
	Sha256 string `json:"sha256"`
	Sha512 string `json:"sha512"`
	Md5    string `json:"md5"`
}
type Nuget struct {
	IsLatestVersion bool   `json:"is_latest_version"`
	IsPrerelease    bool   `json:"is_prerelease"`
	Id              string `json:"id"`
	Version         string `json:"version"`
}

var nexusUrl = os.Getenv("NEXUS_URL")

func CheckNexus(packageName string, packageVersion string, repository string) (Item, error) {

	apiURL := fmt.Sprintf("%s/service/rest/v1/search?repository=%s&name=%s&version=%s", nexusUrl, repository, packageName, packageVersion)

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

	var nexus *Response
	err = json.Unmarshal(body, &nexus)
	if err != nil {
		log.Print("Can't unmarsh json for Nexus")
		return nil, err
	}

	if len(nexus.Items) > 0 {
		for _, value := range nexus.Items {
			msg := fmt.Sprintf("Find package:%s version:%s in Nexus.Action \n", value.Name, value.Version)
			log.Print(msg)
		}
		return nexus.Items, nil
	}
	return nil, nil
}
