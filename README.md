# gopretty
Go Pretty a small beautifier for XCode and Unity builds similar to xcpretty

## Requirements
- Go 1.8+

## Installation
### form source
``` bash
$ go get -u github.com/fsuhrau/gopretty
```

### via brew
``` bash
$ brew install fsuhrau/homebrew-tap/gopretty
```

## Usage
``` bash
$ xcodebuild -project 'testproject.xcodeproj' -configuration Release -target "test" | gopretty
```
