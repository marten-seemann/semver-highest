package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Masterminds/semver/v3"
)

func main() {
	var targetInput string
	var versionsInput string
	var includePrereleases bool
	flag.StringVar(&targetInput, "target", "", "the target version")
	flag.StringVar(&versionsInput, "versions", "", "comma-separated list of versions (e.g. v0.1.0,v0.2.0,v0.3.0-alpha)")
	flag.BoolVar(&includePrereleases, "prerelease", false, "include pre-releases")
	flag.Parse()

	if targetInput == "" || versionsInput == "" {
		flag.Usage()
		os.Exit(1)
	}

	highest, err := GetHighestBefore(targetInput, versionsInput, includePrereleases)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(highest)
}

func GetHighestBefore(targetInput string, versionsInput string, includePrereleases bool) (string, error) {
	if targetInput[0] != 'v' {
		return "", errors.New("target version is not a version number")
	}
	target, err := semver.StrictNewVersion(targetInput[1:])
	if err != nil {
		return "", fmt.Errorf("target version is not a semver version: %s", targetInput)
	}

	versions := make([]*semver.Version, 0, len(versionsInput))
	for _, v := range strings.Split(versionsInput, ",") {
		if v[0] != 'v' { // not a version tag
			continue
		}
		ver, err := semver.StrictNewVersion(v[1:])
		if err != nil { // not a semver version tag
			continue
		}
		versions = append(versions, ver)
	}

	highest, err := getHighestBefore(target, versions, includePrereleases)
	if err != nil {
		return "", err
	}
	return "v" + highest.String(), nil
}

func getHighestBefore(target *semver.Version, versions []*semver.Version, includePrereleases bool) (*semver.Version, error) {
	highest := semver.MustParse("0.0.0")
	for _, v := range versions {
		if target.Equal(v) {
			return nil, fmt.Errorf("target version v%s already exists", v)
		}
		if !includePrereleases && v.Prerelease() != "" {
			continue
		}
		if target.GreaterThan(v) {
			if v.GreaterThan(highest) {
				highest = v
			}
		}
	}
	return highest, nil
}
