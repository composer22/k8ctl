# k8ctl

A simple client to communicate to a k8ctl-server

Written in [Golang.](http://golang.org)

## About

CLI tool useful for displaying and managing deployments, pods, and services
within kubernetes clusters, while restricting more volatile commands.

## Command Line Usage

```
A command line client for deploying and managing applications and releases in a cluster/namespace.

Usage:
  k8ctl [command]

Available Commands:
  cronjobs    Display cronjob infomation
  deployments Display and restart deployments
  guide       Usage guide for the application
  help        Help about any command
  ingresses   Display ingress infomation
  jobs        Display job infomation
  pods        Display pod infomation
  releases    Display and manage helm releases
  services    Display service infomation
  version     Version of the application

Flags:
  -l, --cluster string   Cluster to access (mandatory)
  -c, --config string    config file (default is $HOME/.k8ctl.yaml)
  -h, --help             help for k8ctl

Use "k8ctl [command] --help" for more information about a command.
```
## Configuration

A config file is mandatory. You can place it in the same directory as the
application or in your home directory. The name, with period, is:

.k8ctl.yaml

This file can be empty, but must exist either in your root directory or the
same directory as this application.

An example config file is included under /examples

## Building

This code currently requires version 1.14.1 or higher of Go.

. build.sh is the tool to create multiple executables. Edit what you need/don't need.

For package management, look to dep for instructions: <https://github.com/golang/dep>

commands:
```
dep init # only once.
dep ensure -add <another package>
dep ensure -update
dep status
```

Information on Golang installation, including pre-built binaries, is available at <http://golang.org/doc/install>.

Run `go version` to see the version of Go which you have installed.

Run `go build` inside the directory to build.

Run `go test ./...` to run the unit regression tests.

A successful build run produces no messages and creates an executable called `k8ctl` in this
directory.

Run `go help` for more guidance, and visit <http://golang.org/> for tutorials, presentations, references and more.

## License

(The MIT License)

Copyright (c) 2020 Pyxxel Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
