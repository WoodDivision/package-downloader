package npm

import (
	"log"
	"regexp"
)

type ToDo struct {
	Name    string
	Version string
}

var p = make(map[ToDo]bool)

func CheckDependency(pac ToDo) (map[ToDo]bool, error) {
	npmPac, _ := GetPackage(pac.Name)
	for name, version := range npmPac.Versions[pac.Version].Dependencies {
		reg := regexp.MustCompile("[/^|~]")
		depVersion := reg.ReplaceAllString(version, "")
		newDep := ToDo{Name: name, Version: depVersion}
		p[newDep] = false
	}
	for pac, load := range p {
		if load == true {
			continue
		}
		log.Printf("Processing package %s, %s ", pac.Name, pac.Version)
		p[pac] = true
		return CheckDependency(pac)
	}
	return p, nil
}
