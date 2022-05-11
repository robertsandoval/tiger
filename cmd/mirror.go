package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"sort"
	"strconv"
	"strings"
)

var (
	allVersions      = make(map[string]string)
	versions         = make([]OCPversion, 0)
	filteredVersions = make([]OCPversion, 0)
)

type OCPversion struct {
	category      string
	version       string
	majorVersion  int
	minorVersion  int
	patchVersion  int
	ocpInstallUrl string
	ocpCliUrl     string
}

func (ocp OCPversion) String() string {
	return fmt.Sprintf("%+v %+v", ocp.category, ocp.version)
}

func getAllVersions() {
	c := colly.NewCollector(
		colly.AllowedDomains("mirror.openshift.com"),
	)
	c.OnHTML(".file", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		links[0] = strings.Trim(links[0], "\\/")
		//		fmt.Println(links[0])
		if links[0] != "candidate" && links[0] != "fast" && links[0] != "latest" && links[0] != "stable" && links[0] != "unreleased" {
			versions = append(versions, parseVersion(links[0]))
		}
		//		allVersions[links[0]] = "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/" + links[0]
	})
	c.Visit("https://mirror.openshift.com/pub/openshift-v4/clients/ocp/")
}

func getVersions(category string) []OCPversion {
	getAllVersions()
	if category != "all" {
		for _, ocp := range versions {
			if ocp.category == category {
				filteredVersions = append(filteredVersions, ocp)
			}
		}
	} else {
		filteredVersions = versions
	}
	return sortVersions(filteredVersions)
}
func parseVersion(version string) OCPversion {
	var ocp = OCPversion{}
	var x []string
	s := strings.Split(version, "-")

	if len(s) == 2 {
		ocp.category = s[0]
		ocp.version = s[1]
		x = strings.Split(s[1], ".")
	} else {
		ocp.category = "all"
		ocp.version = s[0]
		x = strings.Split(s[0], ".")
	}
	ocp.majorVersion, _ = strconv.Atoi(x[0])
	ocp.minorVersion, _ = strconv.Atoi(x[1])
	ocp.ocpInstallUrl = "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/" + ocp.version + "/" + "openshift-install-linux.tar.gz"
	ocp.ocpCliUrl = "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/" + ocp.version + "/" + "openshift-client-linux.tar.gz"

	return ocp
}

//TODO Create sorted list return
// i.e. return the map and a sorted array of keys based on sort criteria (up,down..etc)
func sortVersions(versions []OCPversion) []OCPversion {
	sort.Slice(versions, func(i, j int) bool {
		return versions[i].minorVersion > versions[j].minorVersion
	})
	return versions
}
