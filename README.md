# gopretty
Go Pretty a small beautifier for XCode builds similar to xcpretty

## Requirements
- Go 1.6+

## Installation
``` bash
$ go get -u github.com/fsuhrau/gopretty
```

## Usage
``` bash
$ xcodebuild -project 'testproject.xcodeproj' -configuration Release -target "test" | gopretty
```
