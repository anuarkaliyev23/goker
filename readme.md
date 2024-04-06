# Goker

Goker is a CLI tool for utilities related to poker.  
This tool took heavy inspiration from [poker-odds](https://github.com/CookPete/poker-odds)

## Usage

## Installation

### Download Release

Every merge to `main` branch creates release accessible from GH page

This repo follows semantic versioning. By default it bumps version by `minor` release.
To change this behaviour you can pass these strings in your message

- `#none`
- `#patch`
- `#minor`
- `#major`

### Build from sources

See [build notes](#Build)

### Hand Odds calculation

#### Texas Hold'em

```shell
goker hand-odds --hands KsTh,8d7d --board KdTsTd2d -i 1000 --texas
```

```
[KsTh]: 100.0%
[8d7d]: 0.0%
Ties: 0.0%
947 ms
```

#### Short-Deck

```shell
goker hand-odds --hands KsTh,8d7d --board KdTsTd2d -i 1000 --short-deck
```

```
[KsTh]: 3.0%
[8d7d]: 97.0%
Ties: 0.0%
953 ms
```

#### Omaha 

```shell
goker hand-odds --hands KsThAcAd,8d7d5c4c --board KdTd9d2d -i 1000 --omaha
```

```
[KsThAcAd]: 10.7%
[8d7d5c4c]: 89.3%
Ties: 0.0%
4737 ms
```

## Roadmap

Technical Stuff:

- [ ] Add Benchmarks for combination calculations

Features:

- [ ] Hand-Odds calculations
    - [x] Texas Hold'em
    - [ ] Omaha
    - [x] Short-Deck
- [ ] Event Possibilities
    - [ ] Draw a specific combination
        - [ ] Texas Hold'em
        - [ ] Omaha
        - [ ] Short-Deck

## Useful Links:

- [awesome-poker](https://github.com/apehex/awesome-poker/tree/master)
- [Basic poker odds and outs](https://www.cardplayer.com/poker-tools/odds-and-outs)
- [hand-odds](https://github.com/CookPete/poker-odds)

## Developer Notes

### Build

#### For development purposes

```
make build
```

This command produces binary to root project folder and can be locally executed as:

```shell
./goker hand-odds --hands KsTh,8d7d --board KdTsTd2d -i 1000 --texas
```

#### Multiplatform

```shell
make build-multiplatform
```

Supported Platforms:

- darwin/amd64 
- darwin/arm64 
- freebsd/amd64 
- freebsd/arm 
- linux/amd64 
- linux/arm 
- linux/arm64 
- windows/amd64

To add new platform please refer to [build script](./build.sh)

