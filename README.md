# demon 😈

[![godoc](https://godoc.org/github.com/streamwithme/demon?status.svg)](https://godoc.org/github.com/streamwithme/demon)
[![goreportcard](https://goreportcard.com/badge/github.com/streamwithme/demon)](https://goreportcard.com/report/github.com/streamwithme/demon)

## Install

From console:

```console
bash <(curl -s https://raw.githubusercontent.com/streamwithme/demon/master/install.sh)
```

## Example

Use the spawn command to daemonize a script or executable.

```console
$ demon spawn ./my-executable
A new demon has been spawned with PID 12345.
```

View running daemons with the list command.

```console
$ demon list
PID  	Uptime       	Exec 	Command
64390	2 seconds ago	demon	./my-executable
```

## Contributing

Pull requests are encouraged.
