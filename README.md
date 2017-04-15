# OATH-TOTP App

- Store TOTP account credentials in an sqlcipher'd database
- Generate TOTP Codes
- Generate QR Codes to add your accounts to multiple devices

This was born from an awful experience with *Google Authenticator* where I
realised I could not export or backup my codes. This is really just temporary
until KeePassXC gets TOTP support.

# N.B.

- Takes password from CLI. You should be using FDE anyway.
- No HOTP support

# Installation

```console
$ go install github.com/erasche/go-otp
```

# Usage

## Initialize the database

```console
$ go-otp init -password 'blah blah blah'
```

## Register New Services

```console
$ go-otp add -password $password -secretKey LONGSECRETKEY -account alice@local.host -issuer AWS
```

## Overwrite Existing Service Entries

```console
$ go-otp add -password $password -secretKey LONGSECRETKEY -account alice@local.host -issuer AWS -update
```

## Generate Codes

```console
$ go-otp -password $password gen
...............
[        alice@local.host ][             example.com ] 584325
```

## Generate QR Codes

```console
$ go-otp -password $password qr
QR Code stored to AWS__alice@local.host.png
QR Code stored to AWS__jane@university.edu.png
QR Code stored to DigitalOcean__alice@gmail.com.png
QR Code stored to GitHub__alice@gmail.com.png
```

# LICENSE

GPLv3
