package main

import (
    "fmt"
    "github.com/robmerrell/comandante"
    "github.com/erasche/go-otp/cmds"
    "bufio"
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

    // run the command
    if err := bin.Run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
