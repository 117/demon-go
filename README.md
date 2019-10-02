# Demon

[![godoc](https://godoc.org/github.com/streamwithme/demon?status.svg)](https://godoc.org/github.com/streamwithme/demon)
[![goreportcard](https://goreportcard.com/badge/github.com/streamwithme/demon)](https://goreportcard.com/report/github.com/streamwithme/demon)

Run any command or script as a daemon.

### Usage
```txt
$ demon help
 _____   ______  __    __  ______  __   __    
/\  __-./\  ___\/\ "-./  \/\  __ \/\ "-.\ \   
\ \ \/\ \ \  __\\ \ \-./\ \ \ \/\ \ \ \-.  \  
 \ \____-\ \_____\ \_\ \ \_\ \_____\ \_\\"\_\ 
  \/____/ \/_____/\/_/  \/_/\/_____/\/_/ \/_/ 

Run any command or script as a daemon.

demon
        help                    show this usage
        spawn <command(s)>      create a new daemon
        destroy <pid|regex>     destroy a previously created daemon
        list <?pid|regex>       ist running daemons
```

### Example

Spawn a new demon with a provided command.
```console
$ demon spawn ping -c 15 example.com
demon > A new demon has been spawned with PID 64390.
```
That's it, your demon is now running in the background.
```console
$ demon list
PID  	Uptime       	Exec 	Command
64390	2 seconds ago	demon	ping -c 15 example.com
```

### Contributing

Pull requests are encouraged.
