package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := os.Mkdir("subdir", 0755)
	errCheck(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		errCheck(os.WriteFile(name, d, 0644))
	}

	err = os.MkdirAll("subdir/parent/child", 0755)
	errCheck(err)
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	errCheck(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child")
	errCheck(err)

	c, err = os.ReadDir(".")
	errCheck(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	errCheck(err)

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
