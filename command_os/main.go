package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("ls", "-a")
	cmd.Dir = "/test/workspace"

	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
}
