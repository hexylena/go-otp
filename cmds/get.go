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
Generate a TOTP/HOTP Code for the given --service
`

var (
	getAccountFlag  string
    getIssuerFlag   string
	getPasswordFlag string
)

func genCode(key string) {
	key = strings.ToUpper(key)
	totp := &otp.TOTP{Secret: key, IsBase32Secret: true}
	token := totp.Get()
	fmt.Printf("[Valid for %02ds] %s\n", (30 - time.Now().Unix()%30), token)
}

func yieldCode(key string) {
	//generate one code
	genCode(key)
	var next_30 int = 30 - int(time.Now().Unix()%30)
	time.Sleep(time.Second * time.Duration(next_30))

	for {
		genCode(key)
		var next_30 int = 30 - int(time.Now().Unix()%30)
		time.Sleep(time.Second*time.Duration(next_30) + time.Millisecond*10)
	}

}

func GetAction() error {
	if getPasswordFlag == "" {
		return errors.New("Must provide -password")
	}

	//if getServiceFlag == "" {
	//fmt.Printf("Provide -service to generate codes for only one service")
	//}

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
        "select password from users where account = '%s' AND issuer = '%s';",
        getAccountFlag,
        getIssuerFlag,
    )
	rows, err := db.Query(e)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var pass string
		rows.Scan(&pass)
		yieldCode(pass)
	}
	rows.Close()

	return nil
}

func GetFlagHandler(fs *flag.FlagSet) {
	fs.StringVar(&getPasswordFlag, "password", "", "Database Password")

	fs.StringVar(&getAccountFlag, "account", "", "Account Name")
	fs.StringVar(&getIssuerFlag, "issuer", "", "Issuer Name")
}
