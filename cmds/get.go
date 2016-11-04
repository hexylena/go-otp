package cmds

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/hgfischer/go-otp"
	_ "github.com/xeodou/go-sqlcipher"
	"strings"
	"time"
)

var GetDoc = `
Generate a TOTP/HOTP Code for the given
`

var (
	getIndex        int
	getPasswordFlag string
)

func genCode(key, account, issuer string) {
	key = strings.ToUpper(key)
	totp := &otp.TOTP{Secret: key, IsBase32Secret: true}
	token := totp.Get()
	fmt.Printf("[%24s%16s] [Valid for %02ds] %s\n", account, issuer, (30 - time.Now().Unix()%30), token)
}

func yieldCode(key, account, issuer string) {
	//generate one code
	genCode(key, account, issuer)
}

func GetAction() error {
	if getPasswordFlag == "" {
		return errors.New("Must provide -password")
	}

	db, err := sql.Open("sqlite3", "auth.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	p := fmt.Sprintf("PRAGMA key = '%s';", getPasswordFlag)
	_, err = db.Exec(p)
	if err != nil {
		return err
	}

	e := fmt.Sprintf(
		"select account,issuer,password from users order by account asc, issuer asc;",
	)
	rows, err := db.Query(e)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			account string
			issuer  string
			pass    string
		)
		rows.Scan(&account, &issuer, &pass)
        yieldCode(pass, account, issuer)
	}
	rows.Close()

	return nil
}

func GetFlagHandler(fs *flag.FlagSet) {
	fs.StringVar(&getPasswordFlag, "password", "", "Database Password")
}
