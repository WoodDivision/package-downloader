package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/go-version"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"package-downloader/rest/npm"
	"package-downloader/rest/nuget"
	"regexp"
	"strings"
	"time"
)

type Client struct {
	user, pass, token          string
	NexusUrl, NugetUrl, NpmUrl string
	HTTPClient                 *http.Client
	NugetClient                *nuget.NugetPackage
	NpmClient                  *npm.NpmPackage
}

type Error struct {
	Code    string `xml:"code,omitempty"`
	Message string `xml:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

const (
	urlNuget = "https://api.nuget.org/v3/registration5-gz-semver2"
	urlNpm   = "https://registry.npmjs.org"
)

func NewClient(usr string, pas string, url string) *Client {
	c := &Client{
		user:       usr,
		pass:       pas,
		HTTPClient: &http.Client{},
		NexusUrl:   url,
		NugetUrl:   urlNuget,
		NpmUrl:     urlNpm,
	}
	c.NugetClient = &nuget.NugetPackage{Client: c}
	c.NpmClient = &npm.NpmPackage{Client: c}
	return c
}

func (c *Client) PosString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

func (c *Client) ContainsString(slice []string, element string) bool {
	return !(c.PosString(slice, element) == -1)
}

func (c *Client) CheckDate(pac time.Time) bool {
	tYear, tMonth, tDay := time.Now().Date()
	if pac.Year() == tYear {
		if pac.Month() >= tMonth-2 {
			if pac.Day() <= tDay {
				return false
			}
		}
	}
	return true
}

func (c *Client) DownloadFile(URL, fileName string) error {

	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create an empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) NormalizeVersion(v string, regex string, replace string) (string, error) {
	reg := regexp.MustCompile(regex)
	ver := reg.ReplaceAllString(v, replace)
	nVersion, err := version.NewVersion(ver)
	if err != nil {
		return "", err
	}
	return nVersion.String(), nil
}

func (c *Client) NormalizeName(name string) string {
	var forbidenSymbols = []string{"@", "/", "."}
	for _, symbol := range forbidenSymbols {
		name, _ = strings.CutPrefix(name, symbol)
		name = strings.ReplaceAll(name, symbol, "-")
	}
	name = strings.ToLower(name)
	return name
}

func (c *Client) SendRequest(method string, u *url.URL, v interface{}) error {
	req := &http.Request{
		Method:     method,
		URL:        u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       u.Host,
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.user, c.pass)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes *Error
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			if errRes.Error != "" {
				return errors.New(errRes.Error)
			}
			if errRes.Message != "" {
				return errors.New(errRes.Message)
			}
		}
		return fmt.Errorf("status code: %d", res.StatusCode)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Unknow Error in Body.Close")
			return
		}
	}(res.Body)
	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return err
	}
	log.Printf("%#v", v)
	return nil
}
