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

## Developer Notes

### Build

```
make build
```

This command produces binary to root project folder and can be locally executed as:

```shell
./goker texas hand-odds --hands KsTs,Kc9c --board KdKh8h -i 1000; 
```


