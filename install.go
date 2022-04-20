package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

var PROJECT_TEMPLATE = "https://github.com/Dino16m/golearn.git"

var DEFAULT_APP_PATH = "./golearn-app"

func install() {
	root := getRoot()
	fetchTemplate(root)
	setupModules(root)
	tidyGoModules(root)
	fmt.Println("Ready to go")
}

func fetchTemplate(root string) {
	fmt.Println(root)
	err := exec.Command("git", "clone", PROJECT_TEMPLATE, root).Run()
	gitSettingsPath := path.Join(root, ".git")
	os.Remove(gitSettingsPath)
	if err != nil {
		panic(err)
	}
}

func tidyGoModules(root string) {
	fmt.Println("Tidying go module")
	cwd, _ := os.Getwd()
	os.Chdir(root)
	exec.Command("go", "mod", "tidy").Run()
	os.Chdir(cwd)
	fmt.Println("Module tidied")
}

func getRoot() string {
	root := getPathArg()
	fmt.Println("ROOT ", root)
	root = strings.TrimSpace(root)
	if len(root) < 1 || root == "." {
		root = DEFAULT_APP_PATH
	}
	return root
}

func getPathArg() string {
	args := os.Args[2:]
	if len(args) < 1 {
		return DEFAULT_APP_PATH
	}
	return args[0]
}
