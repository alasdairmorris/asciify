# asciify

Transliterate file(s) containing Unicode characters into ASCII equivalents.

## Installation

`asciify` will run on most Linux, MacOS and Windows systems.

To install it, just `cd` into the directory in which you wish to install it and then copy-paste the appropriate one-liner from below.

### Linux (32-bit)

```
curl -s -L -o asciify https://github.com/alasdairmorris/asciify/releases/latest/download/asciify-linux-386 && chmod +x asciify
```

### Linux (64-bit)

```
curl -s -L -o asciify https://github.com/alasdairmorris/asciify/releases/latest/download/asciify-linux-amd64 && chmod +x asciify
```

### Mac OS X (Intel)

```
curl -s -L -o asciify https://github.com/alasdairmorris/asciify/releases/latest/download/asciify-darwin-amd64 && chmod +x asciify
```

### Mac OS X (Apple Silicon)

```
curl -s -L -o asciify https://github.com/alasdairmorris/asciify/releases/latest/download/asciify-darwin-arm64 && chmod +x asciify
```

### Windows (32-bit)

```
curl -s -L -o asciify.exe https://github.com/alasdairmorris/asciify/releases/latest/download/asciify-windows-386.exe
```

### Windows (64-bit)

```
curl -s -L -o asciify.exe https://github.com/alasdairmorris/asciify/releases/latest/download/asciify-windows-amd64.exe
```


### Build From Source

If you have Go installed and would prefer to build the app yourself, you can do:

```
go install github.com/alasdairmorris/asciify@latest
```


## Usage

```
Transliterate file(s) containing Unicode characters into ASCII equivalents.

Reads from given files, or stdin by default.

Outputs to stdout.

Usage:
  asciify [FILE...]
  asciify -h | --help
  asciify --version

Global Options:
  -h, --help             Show this screen.
  --version              Show version.

Homepage: https://github.com/alasdairmorris/asciify
```

## Examples

```
$ date | asciify
Mon 12 Aug 2024 10:13:19 BST
```

## License

[MIT](LICENSE)
