package main

import (
	"fmt"

	"hcli/cli"
)

const (
	version   = "0.1.0"
	author    = "Sony AK <sony@sony-ak.com>"
	shortName = "hcli"
	longName  = "Handly CLI"
)

func main() {
	fmt.Println(longName + " v" + version + " by " + author + "\n")
	cli.DetectAvailableCli()
}
