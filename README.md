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
 The build needs Go 13.* and Yarn installed.
 
 ### Building the web pages
 The web pages use [create-react-app](https://github.com/facebook/create-react-app)
 which needs an up-to-date version of Node.
 ```
cd webapp
yarn install
./build.sh
```

### Building the application
The application is packaged as a [snap](https://snapcraft.io/docs) and can be
built using the `snapcraft` command. For testing purposes, it can be run via
```
go run device-config.go
```


## Contributing

This is an [open source](COPYING) project and we warmly welcome community
contributions, suggestions, and constructive feedback. If you're interested in
contributing, please take a look at our [Code of Conduct](https://ubuntu.com/community/code-of-conduct)
first.

- to report an issue, please file [a bug
  report](https://github.com/CanonicalLtd/device-config/issues/new) on our [GitHub issue
tracker](https://github.com/CanonicalLtd/device-config/issues)
