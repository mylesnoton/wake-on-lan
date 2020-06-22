# Wake on Lan

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
