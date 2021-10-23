# gwall

![Github license](https://img.shields.io/github/license/tboerc/gwall)

`gwall` is a CLI tool written in [Go](https://golang.org/) that uses [WinDivert](https://reqrypt.org/windivert.html) to make a firewall for Grand Theft Auto Online.

# Contents

- [Why?](#why)
- [How it works?](#how-it-works)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Avaliable commands](#avaliable-commands)
- [Unistallation](#unistallation)
- [Contributions](#contributions)
- [Thanks to](#thanks-to)

## Why?

Have you ever played Grand Theft Auto Online on PC? Its a swarm of hackers or modders and I'm just tired of it. Sometimes is impossible to have fun with friends or do some open world missions without a griefer or something.

There are some programs already that do this task, like [Guardian](https://www.thedigitalarc.com/software/guardian). But I made a lightweight CLI tool with less functionalities, for now at least, that suit my needs.

## How it works?

This tool uses [WinDivert](https://reqrypt.org/windivert.html) under the hood, feel free to understand how it works, but it's basically a user-mode packet capture-and-divert package. With that, the CLI can capture packets that make other users connect on your online session, and block it by the source IP. So if there is a packet with unknown source IP, one that is not on whitelist, it will be blocked, and the user with that IP will not join on your session.

Most times the tool works as expected, but there is sometimes that Rockstar changes how your friend connects on your session, so your friend will get blocked. If that happens, turn off your gwall and ask your friend to run it and then join his session, that should do the trick.

## Requirements

- Administrator privileges
- Windows 7 or greater 64 Bits

## Installation

- Download the latest release [here](https://github.com/tboerc/gwall/releases/latest)
- Extract it anywhere you want, but make sure is a folder whitout special characters
- Add the folder with extracted content to your `PATH`, you can follow this [guide](https://github.com/tboerc/gwall/wiki/Adding-gwall-to-system-PATH) if you don't know how to do it

## Usage

If you successfully added the folder to your `PATH` you can run the `gwall` command anywhere with `CMD` or `PowerShell`. Make sure you are executing the commands with Administrator privileges.

### Solo session

To start a solo session just run the following command:

```bash
gwall solo
```

### Whitelist session

There are a few more steps to start a whitelist session, first you need to add all your friends public IP to the list. They can find their public IP with the following command:

```bash
gwall ip
```

But, if somenone don't want to use `gwall`, it's possible to get the public IP on this [site](https://ip.bramp.net). Just remember to use the IPv4 value.

To add the values to the whitelist, run the following command for each IP:

```bash
gwall add 220.191.42.195 # Replace with your friend IP
```

After setting up the list, just run:

```bash
gwall whitelist
```

## Avaliable commands

All `gwall` commands are available below for further usage.

### gwall solo

Block any user from your session.

```bash
gwall solo
```

### gwall whitelist

Allow users from your whitelist to be on your session.

```bash
gwall whitelist
```

### gwall list

List all IP addresses on your whitelist.

```bash
gwall list
# output ↓
Index     IP
0         220.191.42.191
1         220.191.42.192
2         220.191.42.193
```

### gwall add

Add an IP to the whitelist. Most of the time you want to add your friend's public IP, but if your friend plays on the same internet as you, you need to add their local IP.

```bash
gwall add 220.191.42.191
```

### gwall remove

Remove an IP from whitelist. You can get the IP list with `gwall list`.

```bash
gwall remove 220.191.42.191
```

### gwall ip

Display your public IP address. You can use this to send your public IP to a friend and add it to their whitelist.

```bash
gwall ip
# output ↓
220.191.42.191
```

### gwall stop

Stop, if needed, the WinDivert service. Run this if you wish to remove `gwall` and are having trouble deleting it's folder.

```bash
gwall stop
```

## Unistallation

If you want to uninstall `gwall`, just delete it's folder. If you have problems deleting the folder, just type `gwall stop` on CMD or PowerShell to fully stop the WinDivert service and you will be able to delete the folder.

## Build from source

You can build this from source by cloning this repo on your machine. You will need [Go](https://golang.org/) and [make](http://gnuwin32.sourceforge.net/packages/make.htm) for that.

```bash
git clone https://github.com/tboerc/gwall.git
go mod tidy
make build
```

## Contributions

All contributions are helpful, feel free to make a Pull Request.

## Thanks to

- [Guardian](https://gitlab.com/digitalarc/guardian)
- [WinDivert](https://reqrypt.org/windivert.html)
