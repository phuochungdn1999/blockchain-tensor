package main

import (
	"blockchain-golang/cli"
	"os"
)


func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
	
}