# random-scripts

Collection of random scripts 

## pixel-i3-lock

### Requirements

```sh
sudo apt install scrot
```

## .sqliterc

Just a SQLite configuration file. Add this to your home directory.

## reading-time.sh

Reads the clipboard text, calculates and prints the appropriate reading time via system notification.

### Usage

1. Highlight any text with your cursor
2. Run this script
3. Profit

## removing-disable-snaps.sh

Removes old version of installed snap packages.

Credit: https://superuser.com/questions/1310825/how-to-remove-old-version-of-installed-snaps (25.06.2019)

## timer.go

A countdown timer in go lang.

### Requirements

Install this go package: [beeep](https://github.com/gen2brain/beeep)

```sh
go get github.com/gen2brain/beeep
```

### Usage

Set the countdown time to 10 minutes:

```go
go run timer.go -minutes=10
```

Print help:

```go
go run timer.go --help
```
