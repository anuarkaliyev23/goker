# Goker

Goker is a CLI tool for utilities related to poker.  
This tool took heavy inspiration from [poker-odds](https://github.com/CookPete/poker-odds)

## Usage

## Installation

At this moment it can only be built from sources. See [build notes](#Build)

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

```
make build
```

This command produces binary to root project folder and can be locally executed as:

```shell
./goker hand-odds --hands KsTh,8d7d --board KdTsTd2d -i 1000 --texas
```


