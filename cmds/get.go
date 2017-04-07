package cmds

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/hgfischer/go-otp"
	_ "github.com/xeodou/go-sqlcipher"
	"github.com/maxmclau/gput"
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

type Account struct {
	account string
	issuer string
	pass string
}

func genCode(an Account) int {
	key := strings.ToUpper(an.pass)
	totp := &otp.TOTP{Secret: key, IsBase32Secret: true}
	token := totp.Get()
	data := fmt.Sprintf("[%24s][%24s] %s\n", an.account, an.issuer, token)
	fmt.Print(data)
	return len(data)
}

func printIt(accounts []Account) {
	var count int64 = 30 - time.Now().Unix() % 30
	fmt.Printf(
		"%s%s\n",
		strings.Repeat(".", int(count)),
		strings.Repeat(" ", 30 - int(count)),
	)
	for _, account := range accounts {
		genCode(account)
	}
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

	accounts := make([]Account, 0)

	for rows.Next() {
		var (
			account string
			issuer  string
			pass    string
		)
		rows.Scan(&account, &issuer, &pass)
		accounts = append(accounts, Account{account, issuer, pass})
	}
	rows.Close()

	// Print immediately
	printIt(accounts)

	// Then only print if we roll over.
	for _ = range time.Tick(1 * time.Second) {
		// Clear
		gput.Cuu1()
		for _ = range accounts {
			gput.Cuu1()
			fmt.Print("\r")
		}

		printIt(accounts)
	}

	return nil
}

func GetFlagHandler(fs *flag.FlagSet) {
	fs.StringVar(&getPasswordFlag, "password", "", "Database Password")
}
