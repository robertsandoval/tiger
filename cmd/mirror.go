package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	versions         = make([]OCPversion, 0)
	filteredVersions = make([]OCPversion, 0)
)

type OCPversion struct {
	channel       string
	version       string
	majorVersion  int
	minorVersion  int
	patchVersion  int
	ocpInstallUrl string
	ocpCliUrl     string
}

func (ocp OCPversion) String() string {
	return fmt.Sprintf("%+v %+v", ocp.channel, ocp.version)
}

func getVersions(channel string) []OCPversion {
	queryAllVersions()
	if channel != "all" {
		for _, ocp := range versions {
			if ocp.channel == channel {
				filteredVersions = append(filteredVersions, ocp)
			}
		}
	} else {
		filteredVersions = versions
	}
	return sortVersions(filteredVersions)
}

//TODO need to move the filtering to a separate function
func queryAllVersions() {
	c := colly.NewCollector(
		colly.AllowedDomains("mirror.openshift.com"),
	)
	c.OnHTML(".file", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		links[0] = strings.Trim(links[0], "\\/")
		if links[0] != "candidate" && links[0] != "fast" && links[0] != "latest" && links[0] != "stable" && links[0] != "unreleased" {
			versions = append(versions, parseVersion(links[0]))
		}
	})
	c.Visit(OCP_MIRROR_URL)
}

//TODO Parse isn't right...need to account for stable- as well as -rc
//TODO change ocp.version to be built from ocp.major + ocp.minor + ocp.patch
func parseVersion(version string) OCPversion {
	//openshift-install-linux-4.10.13.tar.gz

	var ocp = OCPversion{}
	s := strings.Split(version, "-")
	if len(s) == 2 && (s[0] == "fast" || s[0] == "stable" || s[0] == "candidate") {
		ocp.channel = s[0]
		ocp.version = s[1]
		s2 := queryVersion(ocp.channel, ocp.version)
		sa := strings.Fields(s2)
		fullVersion := strings.Split(sa[1], ".")
		ocp.majorVersion, _ = strconv.Atoi(fullVersion[0])
		ocp.minorVersion, _ = strconv.Atoi(fullVersion[1])
		ocp.patchVersion, _ = strconv.Atoi(fullVersion[2])
		ocp.version = fullVersion[0] + "." + fullVersion[1] + "." + fullVersion[2]
		ocp.ocpInstallUrl = "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/" + version + "/" + "openshift-install-linux.tar.gz"
		ocp.ocpCliUrl = "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/" + ocp.version + "/" + "openshift-client-linux.tar.gz"
		//TODO this can go away
	} else {
		ocp.channel = "all"
		ocp.version = s[0]
	}

	return ocp
}

func queryVersion(channel string, version string) string {
	urlstring := "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/" + channel + "-" + version + "/release.txt"
	resp, _ := http.Get(urlstring)
	body, _ := ioutil.ReadAll(resp.Body)
	sb := string(body)
	result, _ := regexp.Compile(`Version:  \d+\.\d+\.\d+`)
	sb = result.FindString(sb)
	return sb
}

//TODO transition to this function after we get it working
func queryMirrorURLs(mirrorURL string) []string {
	var links []string
	c := colly.NewCollector(
		colly.AllowedDomains("mirror.openshift.com"),
	)
	c.OnHTML(".file", func(e *colly.HTMLElement) {
		links = e.ChildAttrs("a", "href")
	})
	c.Visit(mirrorURL)
	return links
}

func filterLinks() {

}

func getVersion(version string) {
	//	queryVersion()

}

//TODO Create sorted list return
// i.e. return the map and a sorted array of keys based on sort criteria (up,down..etc)
func sortVersions(versions []OCPversion) []OCPversion {
	sort.Slice(versions, func(i, j int) bool {
		return versions[i].minorVersion > versions[j].minorVersion
	})
	return versions
}

func parseSpecificVersionLinks() {

}
