package main

import (
	"bufio"
	"fmt"
	"github.com/erasche/go-otp/cmds"
	"github.com/robmerrell/comandante"
	"os"
)

func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	text, _ := reader.ReadString('\n')
	return text
}

func main() {
	bin := comandante.New("auth", "Authenticator app")
	bin.IncludeHelp()

	initCmd := comandante.NewCommand("init", "Initialize storage", cmds.InitAction)
	initCmd.FlagInit = cmds.InitFlagHandler
	initCmd.Documentation = cmds.InitDoc
	bin.RegisterCommand(initCmd)

	listCmd := comandante.NewCommand("list", "List stored services", cmds.ListAction)
	listCmd.FlagInit = cmds.ListFlagHandler // use the flag package to get a url
	listCmd.Documentation = cmds.ListDoc
	bin.RegisterCommand(listCmd)

	addCmd := comandante.NewCommand("add", "Add a new service", cmds.AddAction)
	addCmd.FlagInit = cmds.AddFlagHandler // use the flag package to add a url
	addCmd.Documentation = cmds.AddDoc
	bin.RegisterCommand(addCmd)

	getCmd := comandante.NewCommand("gen", "Generate a TOTP code for a given service", cmds.GetAction)
	getCmd.FlagInit = cmds.GetFlagHandler // use the flag package to get a url
	getCmd.Documentation = cmds.GetDoc
	bin.RegisterCommand(getCmd)

	qrCmd := comandante.NewCommand("qr", "qr PNG QR Code", cmds.QrAction)
	qrCmd.FlagInit = cmds.QrFlagHandler // use the flag package to qr a url
	qrCmd.Documentation = cmds.QrDoc
	bin.RegisterCommand(qrCmd)

	exportCmd := comandante.NewCommand("export", "export database to json", cmds.ExportAction)
	exportCmd.FlagInit = cmds.ExportFlagHandler
	exportCmd.Documentation = cmds.ExportDoc
	bin.RegisterCommand(exportCmd)

	// run the command
	if err := bin.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
