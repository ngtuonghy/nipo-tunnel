package main

import (
	"fmt"
	"os"

	"nipo-tunnel/pkg/command"
)

func main() {
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
