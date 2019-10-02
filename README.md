# Demon

[![godoc](https://godoc.org/github.com/streamwithme/vortex?status.svg)](https://godoc.org/github.com/streamwithme/vortex)
[![goreportcard](https://goreportcard.com/badge/github.com/streamwithme/vortex)](https://goreportcard.com/badge/github.com/streamwithme/vortex)

Run any command or script as a daemon.

### Usage
```shell
 _____   ______  __    __  ______  __   __    
/\  __-./\  ___\/\ "-./  \/\  __ \/\ "-.\ \   
\ \ \/\ \ \  __\\ \ \-./\ \ \ \/\ \ \ \-.  \  
 \ \____-\ \_____\ \_\ \ \_\ \_____\ \_\\"\_\ 
  \/____/ \/_____/\/_/  \/_/\/_____/\/_/ \/_/ 

Run any command or script as a daemon.

demon 
  help - show this usage
  spawn <command(s)> - create a new daemon
  destroy <pid|regex> - destroy a previously created daemon
  list <?pid|regex> - list running daemons
```

### Example

```console
$ demon spawn ./my-executable
$ demon list
PID     Uptime      Exec    Command
48392   5 seconds   demon   ./my-executable
```

### Contributing

Pull requests are encouraged.
