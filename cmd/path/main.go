package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// demonstrateFilepath shows functions that only handle path string manipulation.
// They do not interact with the file system.
func demonstrateFilepath() {
	fmt.Println("--- Demonstration: 'path/filepath' only ---")

	// Building a path safely for any OS.
	path := filepath.Join("home", "user", "documents", "report.docx")
	fmt.Printf("Path with Join: %s\n", path)

	// Extracting parts of the path.
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	ext := filepath.Ext(path)
	fmt.Printf("Dir: %s, Base: %s, Ext: %s\n", dir, base, ext)

	// Cleaning up a path.
	dirtyPath := filepath.Join("home", "user", "..", "docs", ".", "file.txt")
	cleanPath := filepath.Clean(dirtyPath)
	fmt.Printf("Dirty path: %s\nClean path: %s\n", dirtyPath, cleanPath)

	// Splitting the path into directory and file.
	dir, file := filepath.Split(path)
	fmt.Printf("Split -> Directory: %s, File: %s\n", dir, file)

	// Checking if a name matches a pattern.
	matched, _ := filepath.Match("*.docx", base)
	fmt.Printf("Does the file '%s' match '*.docx'? %t\n", base, matched)

	// Calculating the relative path between two paths.
	basepath := "/home/user/go"
	targetpath := "/home/user/go/src/project"
	relativePath, _ := filepath.Rel(basepath, targetpath)
	fmt.Printf("Relative path from '%s' to '%s' is: '%s'\n", basepath, targetpath, relativePath)
}

// demonstrateFilepathWithOS shows functions that interact with the file system.
func demonstrateFilepathWithOS() {
	fmt.Println("\n--- Demonstration: 'path/filepath' with 'os' ---")

	// Checking if a path is absolute.
	fmt.Printf("Is '/home/user' absolute? %t\n", filepath.IsAbs("/home/user"))
	fmt.Printf("Is 'documents' absolute? %t\n", filepath.IsAbs("documents"))

	// Getting the absolute path of a file (requires OS query).
	absPath, err := filepath.Abs("main.go")
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
	} else {
		fmt.Printf("Absolute path of 'main.go': %s\n", absPath)
	}

	// Getting the current working directory.
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
	} else {
		fmt.Printf("Current working directory: %s\n", wd)
	}

	// Checking if a file or directory exists.
	if info, err := os.Stat("main.go"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("'main.go' does not exist.")
		}
	} else {
		fmt.Printf("'main.go' exists and is a directory? %t\n", info.IsDir())
	}
}

func walkDir() {
	fmt.Println("----------WALKDIR-----------")
	root := "."

	extToFind := ".go"

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Erro ao tentar ler o diretório %q: %v\n", path, err)
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) == extToFind {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Erro ao percorrer o diretório: %v\n", err)
	}
}

func main() {
	demonstrateFilepath()
	demonstrateFilepathWithOS()

	walkDir()
}
