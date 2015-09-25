package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/blang/semver"
)

func main() {
	var (
		e           error
		l           string
		versionList semver.Versions
	)
	exec.Command("git", "fetch", "--tags").Run()
	cmd := exec.Command("git", "tag", "-l")
	b, _ := cmd.Output()
	bfr := bytes.NewBuffer(b)

	for e == nil {
		l, e = bfr.ReadString('\n')

		if v, ve := semver.Make(strings.Replace(strings.TrimSpace(l), "v", "", -1)); ve == nil {
			versionList = append(versionList, v)
		}
	}
	semver.Sort(versionList)
	currentVersion := versionList[len(versionList)-1]

	switch os.Args[1] {
	case "bump_patch":
		currentVersion.Patch++

	case "bump_minor":
		currentVersion.Patch = 0
		currentVersion.Minor++

	case "bump_major":
		currentVersion.Patch = 0
		currentVersion.Minor = 0
		currentVersion.Major++
	}
	fmt.Printf("v%s", currentVersion)
}
