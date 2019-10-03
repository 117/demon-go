# Demon

[![godoc](https://godoc.org/github.com/streamwithme/demon?status.svg)](https://godoc.org/github.com/streamwithme/demon)
[![goreportcard](https://goreportcard.com/badge/github.com/streamwithme/demon)](https://goreportcard.com/report/github.com/streamwithme/demon)

Run any command or script as a daemon.

## Install

On Linux:

```console
$ wget https://github.com/streamwithme/demon/releases/download/v1.0.0/demon-linux-v1.0.0
$ chmod +x demon-linux-v1.0.0 && sudo mv demon-linux-v1.0.0 /usr/bin/demon
```

On Darwin (macOS):

```console
$ wget https://github.com/streamwithme/demon/releases/download/v1.0.0/demon-darwin-v1.0.0
$ chmod +x demon-darwin-v1.0.0 && sudo mv demon-darwin-v1.0.0 /usr/bin/demon
```

## Usage

```txt
 _____   ______  __    __  ______  __   __    
/\  __-./\  ___\/\ "-./  \/\  __ \/\ "-.\ \   
\ \ \/\ \ \  __\\ \ \-./\ \ \ \/\ \ \ \-.  \  
 \ \____-\ \_____\ \_\ \ \_\ \_____\ \_\\"\_\ 
  \/____/ \/_____/\/_/  \/_/\/_____/\/_/ \/_/ 

Run any command or script as a daemon.

demon help                      show this usage
demon spawn <command(s)>        create a new daemon
demon destroy <pid|regex>       destroy a previously created daemon
```

### Example

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

### Contributing

Pull requests are encouraged.
