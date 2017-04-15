package main

import (
	"bufio"
	"fmt"
	"github.com/erasche/go-otp/cmds"
	"github.com/urfave/cli"
	"os"
	"os/user"
	"path"
)

var (
	version   string
	builddate string
)

func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	text, _ := reader.ReadString('\n')
	return text
}

func main() {
	app := cli.NewApp()
	app.Name = "go-otp"
	app.Usage = "Simple TOTP tool"
	app.Version = fmt.Sprintf("%s (%s)", version, builddate)
	user, err := user.Current()

	var defaultDbPath string
	if err != nil {
		defaultDbPath = "auth.db"
	} else {
		defaultDbPath = path.Join(user.HomeDir, ".go-otp.db")
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "dbPath", Value: defaultDbPath, EnvVar: "GOTP_DB_PATH"},
		cli.StringFlag{Name: "password", EnvVar: "GOTP_DB_PASS"},
	}

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a TOTP secret",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "account", Usage: "Account Name (e.g. user@host)"},
				cli.StringFlag{Name: "issuer", Usage: "Issuer (e.g. example.com)"},
				cli.StringFlag{Name: "secretKey", Usage: "TOTP secret key"},
				cli.BoolFlag{Name: "overwrite"},
			},
			Action: func(c *cli.Context) error {
				cmds.AddCode(
					c.GlobalString("dbPath"),
					c.GlobalString("password"),
					c.GlobalString("account"),
					c.GlobalString("issuer"),
					c.GlobalString("secretKey"),
					c.GlobalBool("overwrite"),
				)
				return nil
			},
		},
		{
			Name:    "export",
			Aliases: []string{"e"},
			Usage:   "export all TOTP secrets",
			Action: func(c *cli.Context) error {
				cmds.ExportSecrets(
					c.GlobalString("dbPath"),
					c.GlobalString("password"),
				)
				return nil
			},
		},
		{
			Name:    "gen",
			Aliases: []string{"g"},
			Usage:   "Show current TOTP tokens",
			Action: func(c *cli.Context) error {
				cmds.GenerateCodes(
					c.GlobalString("dbPath"),
					c.GlobalString("password"),
				)
				return nil
			},
		},
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize database",
			Action: func(c *cli.Context) error {
				cmds.InitDb(
					c.GlobalString("dbPath"),
					c.GlobalString("password"),
				)
				return nil
			},
		},
		{
			Name:    "qr",
			Aliases: []string{"q"},
			Usage:   "Dump QR codes as PNG files",
			Action: func(c *cli.Context) error {
				cmds.QrCodes(
					c.GlobalString("dbPath"),
					c.GlobalString("password"),
				)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
