package main

import (
	_ "embed"

	"github.com/0xDaoChain/dao-chain/command/root"
	"github.com/0xDaoChain/dao-chain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
