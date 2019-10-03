# Demon

[![godoc](https://godoc.org/github.com/streamwithme/demon?status.svg)](https://godoc.org/github.com/streamwithme/demon)
[![goreportcard](https://goreportcard.com/badge/github.com/streamwithme/demon)](https://goreportcard.com/report/github.com/streamwithme/demon)

Run any command or script as a daemon.

### Install

On linux systems:

```console
$ wget https://github.com/streamwithme/demon/releases/download/v0.0.9/demon-linux-v0.0.9
$ chmod +x demon-linux-v0.0.9
$ mv demon-linux-v0.0.9 /usr/bin/demon
```

### Example

Want to run a command or script as a daemon?

```console
$ demon spawn ping -c 15 example.com
demon > A new demon has been spawned with PID 64390.
$ demon list
PID  	Uptime       	Exec 	Command
64390	2 seconds ago	demon	ping -c 15 example.com
```

### Contributing

Pull requests are encouraged.
