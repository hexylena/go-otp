package cmds

import (
	"database/sql"
	"fmt"

	_ "github.com/xeodou/go-sqlcipher"
)

func ExportSecrets(dbLoc, exportPasswordFlag string) error {
	db, err := sql.Open("sqlite3", dbLoc)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	p := fmt.Sprintf("PRAGMA key = '%s';", exportPasswordFlag)
	_, err = db.Exec(p)
	if err != nil {
		return err
	}

	e := "select account,issuer,password from users;"
	rows, err := db.Query(e)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			account  string
			issuer   string
			password string
		)
		rows.Scan(&account, &issuer, &password)
		fmt.Printf("%s\t%s\t%s\n", account, issuer, password)
	}
	rows.Close()

	return nil
}
