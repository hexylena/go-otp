package cmds

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/xeodou/go-sqlcipher"
)

var ExportDoc = `
Export available services
`

var (
	exportPasswordFlag string
)

func ExportAction() error {
	if exportPasswordFlag == "" {
		return errors.New("Must provide -password")
	}

	db, err := sql.Open("sqlite3", "auth.db")
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
            account string
            issuer string
            password string
        )
		rows.Scan(&account, &issuer, &password)
		fmt.Printf("%s\t%s\t%s\n", account, issuer, password)
	}
	rows.Close()

	return nil
}

func ExportFlagHandler(fs *flag.FlagSet) {
	fs.StringVar(&exportPasswordFlag, "password", "", "Database Password")
}
