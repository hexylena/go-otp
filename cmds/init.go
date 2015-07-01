package cmds

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/xeodou/go-sqlcipher"
)

var InitDoc = `
Setup a database for storage
`
var (
	passwordFlag string
)

func InitAction() error {
	if passwordFlag == "" {
		return errors.New("Password must not be empty")
	}

	// TODO check that file doesn't exist
	db, err := sql.Open("sqlite3", "auth.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	p := fmt.Sprintf("PRAGMA key = '%s';", passwordFlag)
	_, err = db.Exec(p)
	if err != nil {
		return err
	}

	c := "CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER PRIMARY KEY, `account` char, `issuer` char, `password` chart);"
    //account
    //issuer
    //password
	_, err = db.Exec(c)
	if err != nil {
		return err
	}

	return nil
}

func InitFlagHandler(fs *flag.FlagSet) {
	fs.StringVar(&passwordFlag, "password", "", "Master database password")
}
