package main

import (
	"fmt"

	"hcli/cli"
)

const (
	version   = "0.2.0"
	author    = "Sony AK <sony@sony-ak.com>"
	shortName = "hcli"
	longName  = "HandyCLI"
)

func main() {
	fmt.Println(longName + " v" + version + " by " + author + "\n")
	cli.DetectAvailableCli()
}
