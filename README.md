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
# validate a custom config
$ gopretty validate --config example.yaml
# beatufiy with default rules
$ xcodebuild -project 'testproject.xcodeproj' -configuration Release -target "test" | gopretty
# beautify with custom rules
$ xcodebuild -project 'testproject.xcodeproj' -configuration Release -target "test" | gopretty beautify --config example.yaml
```
