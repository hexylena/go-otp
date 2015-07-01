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
	addServiceFlag   string
	addSecretKeyFlag string
	addPasswordFlag  string
    addUpdate        bool
)

func AddAction() error {
	if addSecretKeyFlag == "" || addServiceFlag == "" || addPasswordFlag == "" {
		return errors.New("Must provide -service, -secretKey, and -password")
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
            "UPDATE `users` set password = '%s' where name = '%s';",
            addSecretKeyFlag,
            addServiceFlag,
        )
        _, err = db.Exec(d)
        if err != nil {
            return err
        }
    } else {

        d := fmt.Sprintf(
            "INSERT INTO `users` (name, password) values('%s', '%s');",
            addServiceFlag,
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
	fs.StringVar(&addServiceFlag, "service", "", "Service Name")
	fs.StringVar(&addPasswordFlag, "password", "", "Database Password")
	fs.StringVar(&addSecretKeyFlag, "secretKey", "", "Secret Key")
	fs.BoolVar(&addUpdate, "update", false, "Overwrite any existing key")
}
