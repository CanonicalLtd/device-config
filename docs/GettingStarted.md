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
sudo snap install --devmode device-config-demo
```
The snap uses some privileged interfaces, so `--devmode` is required. To use the
snap in `strict` mode will need a Brand store, which is a commercial service
from Canonical.

## Accessing the web interface
The web interface is accessible at [http://<ip-address>:8888/](http://<ip-address>:8888/)

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
