package cmds

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/xeodou/go-sqlcipher"
)

var ListDoc = `
List available services
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

	e := "select name from users;"
	rows, err := db.Query(e)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Printf("Services:\n\n")

	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Printf("\t%s\n", name)
	}
	rows.Close()

	return nil
}

func ListFlagHandler(fs *flag.FlagSet) {
	fs.StringVar(&listPasswordFlag, "password", "", "Database Password")
}
