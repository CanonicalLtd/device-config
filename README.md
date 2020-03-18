[![Build Status][travis-image]][travis-url]
[![Go Report Card][goreportcard-image]][goreportcard-url]
[![codecov][codecov-image]][codecov-url]
# Ubuntu Configuration

This configuration application for Ubuntu Core provides a password-protected web interface
to configure a device, listening on one specific interface. The features
that are provided are:

- Wired network configuration
- System proxy configuration
- System NTP server configuration
- System time zone configuration
- Basic health check of the system

 
 ## Development Environment
 The build needs Go 13.* and npm installed.
 
 ### Building the web pages
 The web pages use [create-react-app](https://github.com/facebook/create-react-app)
 which needs an up-to-date version of Node.
 ```
cd webapp
npm install
./build.sh
```

### Building the application
The application is packaged as a [snap](https://snapcraft.io/docs) and can be
built using the `snapcraft` command. The [snapcraft.yaml](snap/snapcraft.yaml)
is the source for building the application and the name of the snap needs to be
modified in that file. The application needs a number of privileged interfaces
that need to be enabled if it is to run in strict mode. However, it can be built
and installed in devmode without them.

For testing purposes, it can also be run via:
```
go run device-config.go
```

## Documentation
The [Getting Started](docs/GettingStarted.md) guide provides instructions on
installing and using the application.


## Contributing

This is an [open source](COPYING) project and we warmly welcome community
contributions, suggestions, and constructive feedback. If you're interested in
contributing, please take a look at our [Code of Conduct](https://ubuntu.com/community/code-of-conduct)
first.

- to report an issue, please file [a bug
  report](https://github.com/CanonicalLtd/device-config/issues/new) on our [GitHub issue
tracker](https://github.com/CanonicalLtd/device-config/issues)


[travis-image]: https://travis-ci.org/CanonicalLtd/device-config.svg?branch=master
[travis-url]: https://travis-ci.org/CanonicalLtd/device-config
[goreportcard-image]: https://goreportcard.com/badge/github.com/CanonicalLtd/device-config
[goreportcard-url]: https://goreportcard.com/report/github.com/CanonicalLtd/device-config
[codecov-url]: https://codecov.io/gh/CanonicalLtd/device-config
[codecov-image]: https://codecov.io/gh/CanonicalLtd/device-config/branch/master/graph/badge.svg
