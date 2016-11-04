package cmds

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/xeodou/go-sqlcipher"
)

var ListDoc = `
List available accounts
`

var (
	listPasswordFlag string
)

func ListAction() error {
	if listPasswordFlag == "" {
		return errors.New("Must provide -password")
	}

	db, err := sql.Open("sqlite3", "auth.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	p := fmt.Sprintf("PRAGMA key = '%s';", listPasswordFlag)
	_, err = db.Exec(p)
	if err != nil {
		return err
	}

	e := "select account,issuer from users order by account asc, issuer asc;"
	rows, err := db.Query(e)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Printf("Services: Issuer : Account\n\n")

	i := 0
	for rows.Next() {
		var (
			account string
			issuer  string
		)
		rows.Scan(&account, &issuer)
		fmt.Printf("\t[%d] %s : %s\n", i, issuer, account)
		i += 1
	}
	rows.Close()

	return nil
}

func ListFlagHandler(fs *flag.FlagSet) {
	fs.StringVar(&listPasswordFlag, "password", "", "Database Password")
}
