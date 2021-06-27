# Wake on Lan

[![CircleCI](https://circleci.com/gh/mylesnoton/wake-on-lan.svg?style=svg&circle-token=7c3b4212d5c7125641ed2855569cd26da48bdcd8)](https://circleci.com/gh/mylesnoton/wake-on-lan)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mylesnoton/password-generator/blob/master/LICENSE)

A command line tool to broadcast a Wake on LAN magic packet to devices on your network, written in Go

## Usage

```bash
Usage:
  wake-on-lan <Options>

Command Options:
  --mac             The MAC address of the machine you want to wake up, eg: 192.168.0.1
  --bcast           The IP address you want to broadcast the magic packet to on your network
```

## Example

```bash
wake-on-lan --mac 00-00-00-00-00-00
```

## Testing

```bash
go test -v
```
