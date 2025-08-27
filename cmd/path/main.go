package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func demonstrateFilepath() {
	path := filepath.Join("home", "user", "documents", "reports.docx")
	fmt.Printf("Caminho construído com Join: %s\n", path)

	dir := filepath.Dir(path)
	fmt.Printf("Diretório (dir): %s\n", dir)

	base := filepath.Base(path)
	fmt.Printf("Nome base (Base) %s\n", base)

	ext := filepath.Ext(path)
	fmt.Printf("Extensão (Ext): %s\n", ext)

	fmt.Printf("É absoluto? %t\n", filepath.IsAbs(path))
	fmt.Printf("'/home/user' é absoluto? %t\n", filepath.IsAbs("/home/user"))

	joinPath := filepath.Join("home", "user", "..", "user", "docs", ".", "file.txt")
	cleanPath := filepath.Clean(joinPath)

	fmt.Printf("Caminho %s\n", joinPath)
	fmt.Printf("Caminho limpo %s\n", cleanPath)

	absPath, err := filepath.Abs("main.go")
	if err != nil {
		fmt.Printf("Erro ao obter o caminho absoluto:", err)
	} else {
		fmt.Printf("Caminho absoluto (abs) de 'main.go': %s\n", absPath)
	}

	// dividir um caminho em diretório e nome nome do arquivo
	dir, file := filepath.Split(path)
	fmt.Printf("Split -> Diretório: %s, Arquivo: %s\n", dir, file)
}

func demonstrateFilepathWithOS() {
	fmt.Printf("'/home/user é absoluto? %t\n", filepath.IsAbs("/home/user"))
	fmt.Printf("'documents' é absoluto? %t\n", filepath.IsAbs("documents"))

	absPath, err := filepath.Abs("main.go")
	if err != nil {
		fmt.Println("Erro ao obter caminho absoluto: ", err)
	} else {
		fmt.Printf("Caminho absoluto de 'main.go' %s\n", absPath)
	}

	workDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho atual ", err)
	} else {
		fmt.Printf("Diretório de trabalho atual: %s\n", workDirectory)
	}

	// verificando se um arquivo ou diretório existe
	if info, err := os.Stat("main.go"); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("'main.go' não existe.")
		}
	} else {
		fmt.Printf("'main.go' existe e é um diretório? %t\n", info)
	}

}

func main() {
	demonstrateFilepath()
	demonstrateFilepathWithOS()
}
