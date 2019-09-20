# externalip
CLI tool that reports your external IP address. This uses ipify.org to report 
the IP address information.

## Compiling
Ensure [Go](https://golang.org) 1.12+ is installed. Clone this repository then
execute the following statements in a terminal.

```bash
$ cd externalip
$ go get
$ go build
```

## Running
Once compiled execute it.

```bash
$ ./externalip
```

This will display output similar to the following (note the IP addresses
are changed to protect the innocent!).

```bash
IPv4: 48.111.988.13
IPv6: 3491.1111.1111.1111.1111.1111.2222.333a
```

