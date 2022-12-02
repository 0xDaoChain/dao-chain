package main

import (
	_ "embed"

	"github.com/0xDaoChain/dao-chain/command/root"
	"github.com/0xDaoChain/dao-chainlicenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
