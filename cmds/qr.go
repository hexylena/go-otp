package cmds

import (
	"database/sql"
"bytes"
         "image"
         "image/png"
	"errors"
    "os"
	"flag"
	"fmt"
	_ "github.com/xeodou/go-sqlcipher"
    "code.google.com/p/rsc/qr"
)

var QrDoc = `

`

var (
	qrPasswordFlag string
)

func QrAction() error {
	if qrPasswordFlag == "" {
		return errors.New("Must provide -password")
	}

	db, err := sql.Open("sqlite3", "auth.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	p := fmt.Sprintf("PRAGMA key = '%s';", qrPasswordFlag)
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

	fmt.Printf("Services:\n\n")

	for rows.Next() {
        var (
            account string
            issuer string
            password string
        )
		rows.Scan(&account, &issuer, &password)
        oath_url := fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s", issuer, account, password, issuer)
        //otpauth://totp/Example:alice@google.com?secret=JBSWY3DPEHPK3PXP&issuer=Example
        code, err := qr.Encode(oath_url, qr.H)

        if err != nil {
            return err
        }

        imgByte := code.PNG()
        // convert byte to image for saving to file
        img, _, _ := image.Decode(bytes.NewReader(imgByte))

        //save the imgByte to file
        var filename = fmt.Sprintf("%s__%s.png", issuer, account)

        out, err := os.Create(filename)
        if err != nil {
            return err
        }

        err = png.Encode(out, img)
        if err != nil {
            return err
        }

        fmt.Printf("QR Code stored to %s\n", filename)
	}
	rows.Close()

	return nil
}

func QrFlagHandler(fs *flag.FlagSet) {
	fs.StringVar(&qrPasswordFlag, "password", "", "Database Password")
}
