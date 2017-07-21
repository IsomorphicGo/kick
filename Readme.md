<p align="center"><a href="http://isomorphicgo.org" target="_blank"><img src="https://github.com/isomorphicgo/isogoapp/blob/master/static/images/isomorphic_go_logo.png"></a></p>

# Kick

[![Go Report Card](https://goreportcard.com/badge/github.com/isomorphicgo/kick)](https://goreportcard.com/report/github.com/isomorphicgo/kick)

A lightweight mechanism to provide an *instant kickstart* to a Go web server instance, upon the modification of a Go source file within a particular project directory (including any subdirectories).

An *instant kickstart* consists of a recompilation of the Go code and a restart of the web server instance.

Kick comes with the ability to take both the `go` and `gopherjs` commands into consideration when performing the *instant kickstart*.

## Supported Operating Systems
Kick works on Windows and Unix-like operating systems (e.g., BSD, Linux, Mac OS).

## Installation

Before installing Kick, it is recommended, to install the barebones [isogoapp](https://github.com/isomorphicgo/isogoapp) first – since it will provide you with an example of how to use kick.

### Get Kick
`go get -u github.com/isomorphicgo/kick`

## Usage

### Getting Help

Issue the `help` flag to get help on using the kick command:

`kick --help`

### Running Kick

Example (using GopherJS):

`kick --appPath=$ISOGO_APP_ROOT --mainSourceFile=isogoapp.go --gopherjsAppPath=$ISOGO_APP_ROOT/client`

The `appPath` flag specifies the project directory where the Go application resides.

The `mainSourceFile` flag specifies the Go source file that implements the main function.

The `gopherjsAppPath` flag specifies the directory to the GopherJS client-side application. This flag is optional.

If your Go project is not using GopherJS, you can feel free to omit the `gopherjsAppPath` flag.

### Verify That Kick Is Functioning

Assuming that you've installed the [isogoapp](https://github.com/isomorphicgo/isogo), and you have issued the kick command to run the web server instance:

Access the [test page](http://localhost:8080) for the `isogoapp` using your web browser.

Open up the `client.go` source file in the `$ISOGO_APP_ROOT/client` directory. Change the message that is passed to the `SetInnerHTML()` function call.

Refresh your web browser. You should see your change reflected.

In the command line prompt where you issued the kick command, take note of the "Recompiling and Restarting" message. This is kick's way of telling you that an *instant kickstart* was performed.


## The Isomorphic Go Project
More information on the benefits of Isomorphic Go applications can be found at the [Isomorphic Go Website](http://isomorphicgo.org).

## License
Kick is licensed under the BSD License. Read the [LICENSE](https://github.com/isomorphicgo/kick/blob/master/LICENSE) file for more information.