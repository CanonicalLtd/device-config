# Getting Started
The device configuration application provides a web interface to configure an
Ubuntu Core device. It can also be used and tested on Ubuntu Classic systems,
but is primarily designed for use on IoT devices running Ubuntu Core.

## Installing the snap
To generate a dedicated snap for your devices, first edit the [snapcraft.yaml](../snap/snapcraft.yaml)
file and change the name of the snap. Then build the snap using the `snapcraft`
command. A pre-built demo is available in the global store: `device-config-demo`.
To try it out:

```
sudo snap install device-config-demo
sudo snap connect device-config-demo:network-setup-control :network-setup-control
sudo snap connect device-config-demo:time-control :time-control
sudo snap connect device-config-demo:timeserver-control :timeserver-control
sudo snap connect device-config-demo:timezone-control :timezone-control
sudo snap connect device-config-demo:system-observe :system-observe
```
The snap uses some privileged interfaces, that need to be connected. The option
to set the proxy configuration needs the `snapd-control` interface, which will
need a Brand store to host the snap - a commercial service from Canonical.

The following features need the `snapd-control` interface:
- Configuring the proxy
- Displaying the installed services
- Updating the configuration of snaps

The snap can be tested by installing it in `--devmode` and enabling a setting
that shows the privileged screens:
```
sudo snap set device-config-demo snapcontrol=true
```

### Snap configuration options
The snap allows a number of settings to be configured:
```
sudo snap set device-config-demo key1=value1 key2=value2 ...
```
The valid options are:
- `interface`: which interface the web service listens on (default=0.0.0.0)
- `listenon`: Force the service to listen a specific network device e.g. eth0 (default="")
- `hide`: Comma-separated list of interfaces to hide (default="")
- `proxy`: whether configuration of the proxy is enabled** (default=false)
- `nm`: whether network-manager configuration is used (default=false i.e. netplan is used)

** Needs `devmode` or the `snapd-control` interface to be connected

## Accessing the web interface
The web interface is accessible at [http://ip-address:8888/](http://<ip-address>:8888/)

To web interface requires authentication using a valid MAC address from the
device. On most devices this will be printed on the device, so it needs local
access to the device to access the configuration tool.

## Configuration Options
Once a user is logged into the device, the following configuration services
are available:

`/services`: a summary of the status of the application services running on the device.
`/network`: the ability to configure the network interfaces on the device.
`/proxy`: configure the HTTP, HTTPS and FTP proxy settings.
`/time`: configure the time, time zone, and the use of a time server

### Network configuration
The application generates a dedicated netplan file (`/etc/netplan/00-device-config.yaml`)
and does not remove or any other netplan YAML file that is on the device. The
netplan file will only be applied when the Apply button is clicked.
