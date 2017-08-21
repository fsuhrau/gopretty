# gopretty
Go Pretty a small beautifier for XCode builds similar to xcpretty

## Requirements
- Go 1.6+

## Installation
### form source
``` bash
$ go get -u github.com/fsuhrau/gopretty
```

##ä via brew
``` bash
$ brew install fsuhrau/homebrew-tap/gopretty
```

## Usage
``` bash
$ xcodebuild -project 'testproject.xcodeproj' -configuration Release -target "test" | gopretty
```
