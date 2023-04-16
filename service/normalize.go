package service

import (
	"github.com/hashicorp/go-version"
	"regexp"
	"strings"
)

var forbidenSymbols = []string{"@", "/", "."}

func NormalizeVersion(v string, regex string, replace string) (string, error) {
	reg := regexp.MustCompile(regex)
	ver := reg.ReplaceAllString(v, replace)
	nVersion, err := version.NewVersion(ver)
	if err != nil {
		return "", err
	}
	return nVersion.String(), nil
}

func NormalizeName(name string) string {
	for _, symbol := range forbidenSymbols {
		name, _ = strings.CutPrefix(name, symbol)
		name = strings.ReplaceAll(name, symbol, "-")
	}
	name = strings.ToLower(name)
	return name
}
