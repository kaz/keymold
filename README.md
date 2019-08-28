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
   0.2.0

COMMANDS:
   new, n    add new OTP secret
   get, g    generate OTP
   proxy, p  create SSH proxy tunnel
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### add new OTP secret

```
$ keymold new --help
NAME:
   keymold new - add new OTP secret

USAGE:
   keymold new [command options] key_name

OPTIONS:
   --disable-touch-id  Allow generating OTP without TouchID authentication. (insecure!)

$ keymold new
Input your secret key: AAAAAAAAAAAAAAAAAAAAAAAAAA
```

### generate OTP

```
$ keymold get --help
NAME:
   keymold get - generate OTP

USAGE:
   keymold get key_name

$ keymold proxy
```

### create SSH proxy tunnel

```
$ keymold proxy --help
NAME:
   keymold proxy - create SSH proxy tunnel

USAGE:
   keymold proxy [command options] key_name

OPTIONS:
   -b value, --bastion value  Destination of bastion server. example: [USER_NAME@]HOST_NAME[:PORT]
   -t value, --target value   Destination of target server. example: [USER_NAME@]HOST_NAME[:PORT]

$ ssh-add # only publickey auth is supported. keymold uses ssh-agent to sign.

$ ssh -o "ProxyCommand keymold proxy -b user@bastion.local:10022 -t %h:%p" target-host -p 50022
```
