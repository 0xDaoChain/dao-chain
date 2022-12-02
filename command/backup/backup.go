package backup

import (
	"github.com/0xDaoChain/dao-chain/command"
	"github.com/spf13/cobra"

	"github.com/0xDaoChain/dao-chain/command/helper"
)

func GetCommand() *cobra.Command {
	backupCmd := &cobra.Command{
		Use:     "backup",
		Short:   "Backups the Chain(NEEDS A RUNNING NODE!)",
		PreRunE: runPreRun,
		Run:     runCommand,
	}

	helper.RegisterGRPCAddressFlag(backupCmd)

	setFlags(backupCmd)
	helper.SetRequiredFlags(backupCmd, params.getRequiredFlags())

	return backupCmd
}

func setFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&params.out,
		outFlag,
		"",
		"Backup path",
	)

	cmd.Flags().StringVar(
		&params.fromRaw,
		fromFlag,
		"0",
		"from which height should the backup be made?",
	)

	cmd.Flags().StringVar(
		&params.toRaw,
		toFlag,
		"",
		"till which height should the backup be made?",
	)
}

func runPreRun(_ *cobra.Command, _ []string) error {
	return params.validateFlags()
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter(cmd)
	defer outputter.WriteOutput()

	if err := params.createBackup(helper.GetGRPCAddress(cmd)); err != nil {
		outputter.SetError(err)

		return
	}

	outputter.SetCommandResult(params.getResult())
}
