package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var TARGET = sanitizePath("golearn-api-template")

type Config struct {
	module string
	root   string
	target string
}

func setupModules(root string) {
	cfg := initialize(root)
	setupFiles(cfg)
}

func initialize(root string) Config {
	fmt.Println("Initializing....")

	module := read("What should we name your module? ", true)
	fmt.Println("Configuration acquired")
	return Config{
		root:   root,
		module: sanitizePath(module),
		target: sanitizePath(TARGET),
	}
}

func read(prompt string, required bool) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		content := strings.TrimSpace(text)
		if len(content) > 0 || !required {
			return text
		}
		fmt.Println("No input detected,")
	}
}

func setupFiles(cfg Config) {
	fmt.Println("Scanning for files")
	paths, err := getFiles(cfg.root)
	check(err)
	fmt.Println(len(paths), " files found")
	fmt.Println("Updating go files")
	for _, filepath := range paths {
		updateModuleName(filepath, cfg.target, cfg.module)
	}
	fmt.Println("Updating go.mod")
	updateModule(cfg.root, cfg.target, cfg.module)

	fmt.Println("Module files updated")
}

func sanitizePath(path string) string {
	path = strings.TrimSpace(path)
	path = strings.TrimSuffix(path, "/")
	return path
}

func updateModule(root string, target, replacement string) {
	goMod := path.Join(root, "go.mod")

	content, err := os.ReadFile(goMod)
	contentStr := string(content)
	check(err)
	if !strings.Contains(contentStr, target) {
		return
	}
	updatedContent := strings.ReplaceAll(contentStr, target, replacement)
	os.WriteFile(goMod, []byte(updatedContent), 0644)
}

func updateModuleName(path, targetStr, replacementStr string) {
	targetStr = targetStr + "/"
	replacementStr = replacementStr + "/"
	content, err := os.ReadFile(path)
	contentStr := string(content)
	check(err)
	if !strings.Contains(contentStr, targetStr) {
		return
	}
	updatedContent := strings.ReplaceAll(contentStr, targetStr, replacementStr)
	os.WriteFile(path, []byte(updatedContent), 0644)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func getFiles(root string) (paths []string, err error) {
	paths = []string{}
	err = filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		info, err := entry.Info()

		if err != nil {
			fmt.Println(err)
			return err
		}
		if !entry.IsDir() && isGoModule(path) && info.Mode().IsRegular() {
			paths = append(paths, path)
		}
		return nil
	})

	return paths, err
}

func isGoModule(path string) bool {
	return strings.Contains(path, ".go")
}
