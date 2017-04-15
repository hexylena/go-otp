package cmds

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/hgfischer/go-otp"
	"github.com/maxmclau/gput"
	_ "github.com/xeodou/go-sqlcipher"
)

type Account struct {
	account string
	issuer  string
	pass    string
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
	var count int64 = 30 - time.Now().Unix()%30
	fmt.Printf(
		"%s%s\n",
		strings.Repeat(".", int(count)),
		strings.Repeat(" ", 30-int(count)),
	)
	for _, account := range accounts {
		genCode(account)
	}
}

func GenerateCodes(dbPath, password string) error {
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
