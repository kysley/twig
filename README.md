# twig

## Overview
A tool for me to quickly switch between playback devices on Windows.

## Get device ids

`> Get-AudioDevice -List`

## Installation

Requires [AudioDeviceCmdlets](https://github.com/frgnca/AudioDeviceCmdlets) to be installed.

- Clone repo
- run `go build`
- run `go install`

## Usage
```bash
> twig speakers
# Default playback & comm device is now my speakers :)

> twig headphones
# Default playback & comm device is now my headphones >:)
```

### With [PowerToys](https://docs.microsoft.com/en-us/windows/powertoys/)

shift + space

`'> twig <speakers | headphones>'`
