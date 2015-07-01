package cmds

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/xeodou/go-sqlcipher"
)

var AddDoc = `
Register a new TOTP/HOTP Private key
`

var (
	addAccountFlag   string
	addIssuerFlag    string
	addSecretKeyFlag string
	addPasswordFlag  string
	addUpdate        bool
)

func AddAction() error {
	if addSecretKeyFlag == "" || addIssuerFlag == "" || addAccountFlag == "" || addPasswordFlag == "" {
		return errors.New("Must provide -issuer, -account, -secretKey, and -password")
	}

	db, err := sql.Open("sqlite3", "auth.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	p := fmt.Sprintf("PRAGMA key = '%s';", addPasswordFlag)
	_, err = db.Exec(p)
	if err != nil {
		return err
	}

	if addUpdate {
		d := fmt.Sprintf(
			"UPDATE `users` SET password = '%s' WHERE issuer = '%s' AND account = '%s';",
			addSecretKeyFlag,
			addIssuerFlag,
			addAccountFlag,
		)
		_, err = db.Exec(d)
		if err != nil {
			return err
		}
	} else {

		d := fmt.Sprintf(
			"INSERT INTO `users` (issuer, account, password) values('%s', '%s', '%s');",
			addIssuerFlag,
			addAccountFlag,
			addSecretKeyFlag,
		)
		_, err = db.Exec(d)
		if err != nil {
			return err
		}
	}

	return nil
}

func AddFlagHandler(fs *flag.FlagSet) {
	fs.StringVar(&addAccountFlag, "account", "", "Account Name")
	fs.StringVar(&addIssuerFlag, "issuer", "", "Issuer Name")
	fs.StringVar(&addSecretKeyFlag, "secretKey", "", "Secret Key")

	fs.StringVar(&addPasswordFlag, "password", "", "Database Password")
	fs.BoolVar(&addUpdate, "update", false, "Overwrite any existing key")
}
