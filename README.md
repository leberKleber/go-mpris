# go-mpris

[![Go](https://github.com/leberKleber/go-mpris/workflows/go/badge.svg)](https://github.com/leberKleber/go-mpris/actions?query=workflow%3Ago)
[![Go Report Card](https://goreportcard.com/badge/github.com/leberKleber/go-mpris)](https://goreportcard.com/report/github.com/leberKleber/go-mpris)
[![codecov](https://codecov.io/gh/leberKleber/go-mpris/branch/main/graph/badge.svg)](https://codecov.io/gh/leberKleber/go-mpris)

go mpris version 2.2 implementation

Example cli has been implemented.

```shell
git clone git@github.com:leberKleber/go-mpris.git

go build examples/cli.go

./cli-client
```

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
- [X] PlaybackStatus
- [X] LoopStatus
- [X] Rate
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

#### Signals
- [ ] Seeked
