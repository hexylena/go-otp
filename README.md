# OATH-TOTP App

- Store TOTP account credentials in an sqlcipher'd database
- Generate TOTP Codes

This was born from an awful experience with *Google Authenticator* where I
realised I could not export or backup my codes. Not wanting to move from F/OSS
GA to the closed source Authy app, I wrote a simple piece of software to
securely store my OATH-TOTP private keys, and generate codes on demand at the
command line, which will let me add the accounts to whatever MFA devices I want.

# N.B.

- I'm still learning this language, the code is probably rubbish
- Yes, the database password is taken in plaintext from the CLI, that could
  probably be improved a LOT

# Installation

```console
$ go install github.com/erasche/go-otp
```

# Usage

Currently all commands must be run in the same directory each and every time.
This is bad, I am aware. Future versions will improve upon this.

## Initialize the database

```console
$ go-otp init -password your_space_free_password
```

(If someone wants to PR the password stuff, I would be VERY thankful)

## Register New Services

```console
$ go-otp add -password $password -secretKey 20CHARSTRING -service MyServiceName
```

## Overwrite Existing Service Entries

```console
$ go-otp add -password $password -secretKey NEWSERVICEKEY -service MyServiceName -update
```

## Generate Codes

```console
$ go-otp gen -password $password -service MyServiceName
[Valid for 16s] 123456
[Valid for 30s] 111111
[Valid for 30s] 222222
```

# LICENSE

OATH-TOTP storage and password generator
Copyright © 2015 Eric Rasche <rasche.eric@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

