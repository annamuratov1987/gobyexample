package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.CreateTemp("", "sample")
	checkErr(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	checkErr(err)

	dname, err := os.MkdirTemp("", "sampledir")
	checkErr(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	checkErr(err)
}
