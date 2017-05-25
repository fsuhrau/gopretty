# gopretty
Go Pretty a small beautifier for XCode builds similar to xcpretty

## Requirements
- Go 1.6+

## Installation
go get -u github.com/fsuhrau/gopretty

## Usage
xcodebuild -project 'testproject.xcodeproj' -configuration Release -target "test" | gopretty
