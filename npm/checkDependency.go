package npm

import (
	"log"
	"regexp"
)

type ToDo struct {
	Name    string
	Version string
}

var d = make(map[ToDo]bool)

func GetDependencies(p *Package) (map[ToDo]bool, error) {
	for name, version := range p.Dependencies {
		reg := regexp.MustCompile("[/^]")
		depVersion := reg.ReplaceAllString(version, "")
		newDep := ToDo{Name: name, Version: depVersion}
		d[newDep] = false
	}

	for pac, load := range d {
		log.Printf("Processing package %s, %s ", pac.Name, pac.Version)
		if load == true {
			break
		}
		d[pac] = true
		nextPac, err := GetPackage(pac.Name, pac.Version)
		if err != nil {
			log.Print(err.Error())
			return nil, err
		}
		return GetDependencies(nextPac)
	}
	return d, nil
}
