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
- [Build from source](#build-from-source)
- [Contributions](#contributions)
- [Thanks to](#thanks-to)

## Why?

Have you ever played Grand Theft Auto Online on PC? Its a swarm of hackers or modders and I'm just tired of it. Sometimes is impossible to have fun with friends or do some open world missions without a griefer or something.

There are some programs already that do this task, like [Guardian](https://www.thedigitalarc.com/software/guardian). But I made a lightweight CLI tool with less functionalities, for now at least, that suit my needs.

## How it works?

This tool uses [WinDivert](https://reqrypt.org/windivert.html) under the hood, feel free to understand how it works, but it's basically a user-mode packet capture-and-divert package. With that, the CLI can capture packets that make other users connect on your online session, and block it by the source IP. So if there is a packet with unknown source IP, one that is not on whitelist, it will be blocked, and the user with that IP will not join on your session.

That's the theory, but in practice sometimes some random guy joins the session or your friend gets blocked, even with his IP on the whitelist. However, at most times it works as expected.

## Requirements

- Administrator privileges
- Windows 7 or greater 64 Bit

## Installation

- Download the latest release [here](https://github.com/tboerc/gwall/releases/latest)
- Extract it anywhere you want, but make sure is a folder whitout special characters
- Add the folder with extracted content to your `PATH`, you can follow this [guide](https://gist.github.com/nex3/c395b2f8fd4b02068be37c961301caa7) if you don't know how to do it

## Usage

If you successfully added the folder to your `PATH` you can run the `gwall` command anywhere with `CMD` or `PowerShell`. Make sure you are executing the commands with Administrator privileges.

### Solo session

To start a solo session just run the following command:

```bash
gwall start solo
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
gwall start whitelist
```

## Avaliable commands

All `gwall` commands are available below for further usage.

### gwall start solo

Block any user from your session.

```bash
gwall start solo
```

### gwall start whitelist

Allow users from your whitelist to be on your session.

```bash
gwall start whitelist
```

### gwall list

List all IP addresses on your whitelist.

```bash
gwall list
# output ↓
Index     Public IP         Local IP
0         220.191.42.195    220.191.42.115
1         220.191.42.195    <nil>
2         220.191.42.195    220.191.42.115
```

### gwall add

Add a IP to whitelist, you can add a public or local IP. Most times you just need the public IP, but if your friend plays on the same internet as you, you need to add his local IP.

```bash
gwall add 220.191.42.196 # for only public IP

gwall add nil 220.191.42.196 # for only local IP

gwall add 220.191.42.198 220.191.42.196 # for both IPs
```

### gwall remove

Remove an whitelist row based on it index. You can get the index with `gwall list`.

```bash
gwall remove 1
```

### gwall ip

Display your public IP address. You can use this to send your public IP to a friend add it on his whitelist.

```bash
gwall ip
# output ↓
220.191.42.195
```

### gwall stop

Stop, if needed, the WinDivert service. Run this if you wish to remove `gwall` and are having trouble deleting its folder.

```bash
gwall stop
```

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
