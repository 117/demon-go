# Demon

[![godoc](https://godoc.org/github.com/streamwithme/demon?status.svg)](https://godoc.org/github.com/streamwithme/demon)
![version](https://img.shields.io/github/v/release/streamwithme/demon?color=EB2C2C&include_prereleases)

Run any command or script as a daemon.

### Install

coming soon

### Usage
```console
$ demon
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
