package root

import (
	"fmt"
	"os"

	"github.com/0xDaoChain/dao-chain/command/backup"
	"github.com/0xDaoChain/dao-chain/command/genesis"
	"github.com/0xDaoChain/dao-chain/command/helper"
	"github.com/0xDaoChain/dao-chain/command/ibft"
	"github.com/0xDaoChain/dao-chain/command/license"
	"github.com/0xDaoChain/dao-chain/command/loadbot"
	"github.com/0xDaoChain/dao-chain/command/monitor"
	"github.com/0xDaoChain/dao-chain/command/peers"
	"github.com/0xDaoChain/dao-chain/command/secrets"
	"github.com/0xDaoChain/dao-chain/command/server"
	"github.com/0xDaoChain/dao-chain/command/status"
	"github.com/0xDaoChain/dao-chain/command/txpool"
	"github.com/0xDaoChain/dao-chain/command/version"
	"github.com/0xDaoChain/dao-chain/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Daoc   is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		loadbot.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
