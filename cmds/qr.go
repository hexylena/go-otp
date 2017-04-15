package cmds

import (
	"bytes"
	"database/sql"
	"fmt"
	"image"
	"image/png"
	"os"

	"code.google.com/p/rsc/qr"
	_ "github.com/xeodou/go-sqlcipher"
)

func QrCodes(dbPath, password string) error {
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
