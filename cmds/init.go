package cmds

import (
	"database/sql"
	"fmt"

	_ "github.com/xeodou/go-sqlcipher"
)

func InitDb(dbPath, password string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	p := fmt.Sprintf("PRAGMA key = '%s';", password)
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
