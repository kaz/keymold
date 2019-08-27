# keymold

OTP generator

- works on command line, only has CLI
- works with macOS Keychain
- integrated with TouchID

## install

- Go https://github.com/kaz/keymold/releases
- Download
- Put it on `/usr/local/bin`

## usage

```
$ keymold
NAME:
   keymold - OTP generator, works on command line.

USAGE:
   keymold [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   new, n   add new OTP secret
   get, g   generate OTP
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

$ keymold new test
Input your secret key: AAAAAAAAAAAAAAAAAAAAAAAAAA

$ keymold get test
035093
```
