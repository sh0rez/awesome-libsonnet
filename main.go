package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	yaml "gopkg.in/yaml.v3"
)

type Project struct {
	// Friendly name of the project
	Name string `yaml:"name" mixins:"name"`
	// Online home of it
	URL string `yaml:"url" mixins:"url"`
	// Optional description
	Description string `yaml:"description" mixins:"description"`

	// Application modules for this project
	Modules []Package `yaml:"modules" mixins:"modules"`
	// Configuration mixins for this project
	Mixins []Package `mixins:"mixins"`
}

type Package struct {
	// Name of the package (optional)
	Name string `yaml:"name"`
	// Import (gitub.com/user/repo/...)
	Import string `yaml:"import"`
	// If the package is not on GitHub, set this to the online location
	URL string `yaml:"url"`
	// Maintainer status of the project
	Maintainer string `yaml:"maintainer"`
	// Optional description
	Description string
}

type AwesomeLibsonnet struct {
	Projects []Project
	Helpers  []Package
}

var Maintainers = map[string]string{
	"unmaintained": ":no_entry:",
	"project":      ":crown:",
	"community":    ":busts_in_silhouette:",
	"":             ":busts_in_silhouette:",
}

func main() {
	awesome := AwesomeLibsonnet{
		Projects: LoadProjects("./projects"),
	}

	templates := template.Must(template.ParseFiles("README.tmpl"))
	if err := templates.Execute(os.Stdout, awesome); err != nil {
		log.Fatalln(err)
	}
}

func LoadProjects(dir string) []Project {
	list := []Project{}

	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if ext := filepath.Ext(path); !(ext == ".yaml" || ext == ".yml") {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		p := Project{}
		if err := yaml.Unmarshal(data, &p); err != nil {
			return err
		}

		list = append(list, p)
		return nil
	}); err != nil {
		log.Fatalln(err)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})

	for _, m := range list {
		for i, p := range m.Modules {
			m.Modules[i] = preparePackage(p)
		}
		for i, p := range m.Mixins {
			m.Mixins[i] = preparePackage(p)
		}
	}

	return list
}

func preparePackage(p Package) Package {
	// fix URL
	re := regexp.MustCompile(`(?mU)(github\.com\/.*\/.*\/)(.*)$`)
	if p.URL == "" {
		p.URL = re.ReplaceAllString(p.Import, "https://${1}tree/master/${2}")
	}

	// maintainer
	if k, ok := Maintainers[p.Maintainer]; ok {
		p.Maintainer = k
	} else {
		p.Name = strings.TrimPrefix(p.Name+fmt.Sprintf(" by %s", p.Maintainer), " ")
		p.Maintainer = Maintainers["community"]
	}

	return p
}
