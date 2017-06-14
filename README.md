# OceanConnect

[![Build Status](https://travis-ci.org/dualinventive/go-oceanconnect.svg?branch=master)](https://travis-ci.org/dualinventive/go-oceanconnect)
[![Go Report Card](https://goreportcard.com/badge/github.com/dualinventive/go-oceanconnect)](https://goreportcard.com/report/github.com/dualinventive/go-oceanconnect)
[![GoDoc](https://godoc.org/github.com/dualinventive/go-oceanconnect?status.svg)](https://godoc.org/github.com/dualinventive/go-oceanconnect)
[![codecov](https://codecov.io/gh/dualinventive/go-oceanconnect/branch/master/graph/badge.svg)](https://codecov.io/gh/dualinventive/go-oceanconnect)

OceanConnect is the platform developed by [Huawei](http://developer.huawei.com/ict/en/site-oceanconnect) for the use with [NB-IOT](https://en.wikipedia.org/wiki/NarrowBand_IOT) devices.

NB-IOT is a LPWAN technology for bi-directional data traffic between devices and centralized cloud platforms. OceanConnect is the gateway for these devices.

This library uses the API of OceanConnect to retrieve data and register devices.

NOTE: The API is currently unstable

## Installation

To get started please get the Golang toolchains from the [Golang website](https://golang.org/). When you have a working go toolchain you can do:

```
go get github.com/dualinventive/go-oceanconnect
```

And you are ready to go!

## Included tools

Some simple tools for use with ocean-connect are included and located in the `cmd` folder of the root of the project.

### Register devices (regdevices)

Commandline tool to register devices at OceanConnect. See the readme in the designated folder.

## Contributing

Please read the [Contribution Guidelines](CONTRIBUTING.md). Furthermore: Fork -> Patch -> Push -> Pull Request

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/dualinventive/go-oceanconnect/blob/master/LICENSE) file for the full license text.
