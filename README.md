# whapp

This is a WhatsApp robot proof of concept. You can look at this code if you like, but it won't do that much for you. It's a long way from being pulled from GIT and being a Generally Useful Distribution (git-gud).

## How to run it

- Get the modules: `go mod tidy`
- Build the bot: `go build whapp.go`
- See what flags are supported: `whapp -h`
- Run it: `whapp -v`

This generates the logfile `/tmp/whapp.log`. Leave out `-v` if you don't want debug messages clogging the logfile. During the first invocation you'll be presented a QR code, that you need to scan to authenticate the bot. To do so, open WhatsApp on your phone, press the hamburger menu (three vertical dots) at the right top and select *Settings*. Then press the QR code symbol in the right top corner, then *Scan code*.

This is only needed upon the first invocation. The authentication credentials and keys are stored in a a sqlite3 database called `store.db` (in the current directory from which you invoke `whapp`). If you remove `store.db` then you'll have to re-authenticate.

## What does it do

Currently `whapp` doesn't do anything useful beyond listening to what happens and to print a message to `stdout` when one is received. It will listen forever, until you press `^C`. Or you can let it run for a given duration using the flag `-q` as in `whapp -q 30s`.

## Links

Under the hood `whapp` uses the stock WhatsApp client [whatsmeow](http://go.mau.fi/whatsmeow). Very neat and nice implementation, so respect where it's due.

I've externalized some tooling that I think may be handy into https://github.com/KarelKubat/whatsmeow (events and logging), hoping it's useful to you.