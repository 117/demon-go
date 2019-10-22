# demon

[![godoc](https://godoc.org/github.com/117/demon?status.svg)](https://godoc.org/github.com/117/demon)
[![goreportcard](https://goreportcard.com/badge/github.com/117/demon)](https://goreportcard.com/report/github.com/117/demon)

Run any command or script as a daemon.

## Install

On Linux:

```console
wget https://github.com/117/demon/releases/download/v1.1.0/demon-linux-v1.1.0 && chmod +x demon-linux-v1.1.0 && sudo mv demon-linux-v1.1.0 /usr/bin/demon
```

On Darwin (macOS):

```console
wget https://github.com/117/demon/releases/download/v1.1.0/demon-darwin-v1.1.0 && chmod +x demon-darwin-v1.1.0 && sudo mv demon-darwin-v1.1.0 /usr/local/bin/demon
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
