# Goker

Goker is a CLI tool for utilities related to poker.  
This tool took heavy inspiration from [poker-odds](https://github.com/CookPete/poker-odds)

## Usage

## Installation

At this moment it can only be built from sources. See [build notes](#Build)

### Texas Hold'em

#### Hand Odds calculation

```shell
goker texas hand-odds --hands KsTs,Kc9c --board KdKh8h -i 1000; 
```

```
[KsTs]: 65.6%
[Kc9c]: 12.2%
Ties: 22.2%
1041 ms
```

## Roadmap

Technical Stuff:

- [ ] Add Benchmarks for combination calculations

Features:

- [ ] Hand-Odds calculations
    - [x] Texas Hold'em
    - [ ] Omaha
    - [ ] Short-Deck
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
./goker texas hand-odds --hands KsTs,Kc9c --board KdKh8h -i 1000; 
```


