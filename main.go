package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getLastGoPath(gopath string) string {
	sepRegex := regexp.MustCompile("[:;]")
	pathList := sepRegex.Split(gopath, -1)
	if len(pathList) < 1 {
		return ""
	}

	target := pathList[0]
	for i := 1; i < len(pathList); i++ {
		if pathList[i] != "" {
			target = pathList[i]
		}
	}

	return target
}

func main() {
	var err error
	projectRoot, err := os.Getwd()
	checkError(err)
	importPaths := ""
	gopath := getLastGoPath(os.Getenv("GOPATH"))

	force := false

	flag.StringVar(&projectRoot, "root", projectRoot, "golang project root path")
	flag.StringVar(&importPaths, "path", "", "project import path")
	flag.StringVar(&gopath, "gopath", gopath, "go path")
	flag.BoolVar(&force, "force", force, "force create mode")
	flag.Parse()

	if importPaths == "" {
		importPaths = filepath.Base(projectRoot)
	}

	info, err := os.Stat(projectRoot)
	checkError(err)

	projectPath := filepath.Join(gopath, "src", filepath.Dir(importPaths))
	os.MkdirAll(projectPath, info.Mode())

	projectLink := filepath.Join(projectPath, filepath.Base(importPaths))
	_, err = os.Stat(projectLink)
	if !os.IsNotExist(err) {
		if force {
			os.Remove(projectLink)
		} else {
			fmt.Printf("path[%v] existed\n", projectLink)
			os.Exit(1)
		}
	}

	code := 1
	err = os.Symlink(projectRoot, projectLink)
	if err == nil {
		fmt.Printf("create link %s -> %s\n", projectLink, projectRoot)
		code = 0
	} else if runtime.GOOS == "windows" {
		fmt.Printf("create link failed: %v, are you run as administrator?\n", err)
	} else {
		fmt.Printf("create link failed: %v\n", err)
	}

	os.Exit(code)
}
