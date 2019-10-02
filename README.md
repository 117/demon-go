# Demon

[![godoc](https://godoc.org/github.com/streamwithme/demon?status.svg)](https://godoc.org/github.com/streamwithme/demon)
[![goreportcard](https://goreportcard.com/badge/github.com/streamwithme/demon)](https://goreportcard.com/report/github.com/streamwithme/demon)
![version](https://img.shields.io/github/v/release/streamwithme/demon?color=EF4C4C&include_prereleases)

Run any command or script as a daemon.

### Install

coming soon

### Example

Want to run a command or script as a daemon?
```console
$ demon spawn ping -c 15 example.com
demon > A new demon has been spawned with PID 64390.
```

Done. It's that simple!
```console
$ demon list
PID  	Uptime       	Exec 	Command
64390	2 seconds ago	demon	ping -c 15 example.com
```

### Contributing

Pull requests are encouraged.
