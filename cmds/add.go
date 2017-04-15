package cmds

import (
	"database/sql"
	"fmt"

	_ "github.com/xeodou/go-sqlcipher"
)

func AddCode(dbPath, password, account, issuer, secretKey string, overwrite bool) error {
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

	if overwrite {
		d := fmt.Sprintf(
			"UPDATE `users` SET password = '%s' WHERE issuer = '%s' AND account = '%s';",
			secretKey,
			issuer,
			account,
		)
		_, err = db.Exec(d)
		if err != nil {
			return err
		}
	} else {

		d := fmt.Sprintf(
			"INSERT INTO `users` (issuer, account, password) values('%s', '%s', '%s');",
			issuer,
			account,
			secretKey,
		)
		_, err = db.Exec(d)
		if err != nil {
			return err
		}
	}

	return nil
}
