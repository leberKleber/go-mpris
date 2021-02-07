# go-mpris

[![Go](https://github.com/leberKleber/go-mpris/workflows/Go/badge.svg?branch=master)](https://github.com/leberKleber/go-mpris/actions?query=workflow%3Ago)
[![Go Report Card](https://goreportcard.com/badge/github.com/leberKleber/go-mpris)](https://goreportcard.com/report/github.com/leberKleber/go-mpris)
[![codecov](https://codecov.io/gh/leberKleber/go-mpris/branch/master/graph/badge.svg)](https://codecov.io/gh/leberKleber/go-mpris)

go mpris version 2.2 implementation

## Functions:

### Player

https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html

#### Methods:
- [X] Next
- [X] Previous
- [X] Pause
- [X] PlayPause
- [X] Stop
- [X] Play
- [X] Seek > SeekTo (Seek is a reserved function name and can not be used in this case)
- [X] SetPosition
- [X] OpenUri

#### Properties
- [ ] PlaybackStatus
- [ ] LoopStatus
- [ ] Rate
- [ ] Shuffle
- [ ] Metadata
- [ ] Volume
- [ ] Position
- [ ] MinimumRate
- [ ] MaximumRate
- [ ] CanGoNext
- [ ] CanGoPrevious
- [ ] CanPlay
- [ ] CanPause
- [ ] CanSeek
- [ ] CanControl

## Example

Example cli has been implemented.

```shell
git clone git@github.com:leberKleber/go-mpris.git

go build examples/cli-client.go

./cli-client
```

---
Signals Seeked    (x: Position)

Types Track_Id Simple Type o 	
Playback_Rate Simple Type d 	
Volume Simple Type d 	
Time_In_Us Simple Type x 	
Playback_Status Enum s 	
Loop_Status Enum s 	