package npm

import (
	"log"
	"package-downloader/service"
)

type ToDo struct {
	Name    string
	Version string
}

var p = make(map[ToDo]bool)

func CheckDependency(pac ToDo) (map[ToDo]bool, error) {
	log.Printf("Processing package %s, %s ", pac.Name, pac.Version)
	npmPac, _ := GetNpmPackage(pac.Name)
	for n, v := range npmPac.Versions[pac.Version].Dependencies {
		v, err := service.NormalizeVersion(v, "[/^|~]", "")
		if err != nil {
			log.Print("Can't normalize version")
			break
		}
		newDep := ToDo{Name: n, Version: v}
		p[newDep] = false
	}
	for pac, load := range p {
		if load == true {
			continue
		}
		p[pac] = true
		return CheckDependency(pac)
	}
	return p, nil
}
