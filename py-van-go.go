// go run py-van-go.go
package main

import (
	"fmt"
	"os/exec"
	"path"
	"os"
)

func main() {
	// where python
	python := path.Clean("C:/Users/123/AppData/Local/Programs/Python/Python38/python.exe")
	
	cmd := exec.Command(python, "hoi.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	e := cmd.Run()
	CheckError(e)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}