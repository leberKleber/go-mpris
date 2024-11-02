# go-mpris

[![Go](https://github.com/leberKleber/go-mpris/workflows/go/badge.svg)](https://github.com/leberKleber/go-mpris/actions?query=workflow%3Ago)
[![GoDoc](https://godoc.org/github.com/leberKleber/go-mpris?status.png)](https://godoc.org/github.com/leberKleber/go-mpris)
[![Go Report Card](https://goreportcard.com/badge/github.com/leberKleber/go-mpris)](https://goreportcard.com/report/github.com/leberKleber/go-mpris)
[![codecov](https://codecov.io/gh/leberKleber/go-mpris/branch/main/graph/badge.svg)](https://codecov.io/gh/leberKleber/go-mpris)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

go-mpris is an implementation of the mpris dbus interface written in go (golang).
Implemented and tested against version 2.2. See: https://specifications.freedesktop.org/mpris-spec/2.2.

* [Example](#example)
* [Features](#features)
    * [Player](#player)
        * [Methods](#methods)
        * [Properties](#properties)
        * [Signals](#signals)
    * [TrackList](#tracklist)
        * [Methods](#methods-1)
        * [Properties](#properties-1)
        * [Signals](#signals-1)
    * MediaPlayer
* [Development](#development)
    * [Versioning](#versioning)
    * [Commits](#commits)
    * [Mocks](#mocks)
    * [Go Docs](#go-docs)

## Example

Example cli has been implemented.

```shell
git clone git@github.com:leberKleber/go-mpris.git

go build examples/cli.go

./cli-client
```

## Features

### MediaPlayer (root interface)

https://specifications.freedesktop.org/mpris-spec/2.2/Media_Player.html

#### Methods

| method | library path                | implemented        |
|--------|-----------------------------|--------------------|
| Raise  | `mpris.MediaPlayer.Raise()` | :heavy_check_mark: |
| Quit   | `mpris.MediaPlayer.Quit()`  | :heavy_check_mark: |

#### Properties

| property            | library path                                                | implemented              |
|---------------------|-------------------------------------------------------------|--------------------------|
| CanQuit             | `mpris.MediaPlayer.CanQuit() (bool, error)`                 | :heavy_multiplication_x: |
| Fullscreen          | `mpris.MediaPlayer.Fullscreen() (bool, error)`              | :heavy_multiplication_x: |
| CanSetFullscreen    | `mpris.MediaPlayer.CanSetFullscreen() (bool, error)`        | :heavy_multiplication_x: |
| CanRaise            | `mpris.MediaPlayer.CanRaise() (bool, error)`                | :heavy_multiplication_x: |
| HasTrackList        | `mpris.MediaPlayer.HasTrackList() (bool, error)`            | :heavy_multiplication_x: |
| Identity            | `mpris.MediaPlayer.Identity() (string, error)`              | :heavy_multiplication_x: |
| DesktopEntry        | `mpris.MediaPlayer.DesktopEntry() (string, error)`          | :heavy_multiplication_x: |
| SupportedUriSchemes | `mpris.MediaPlayer.SupportedUriSchemes() ([]string, error)` | :heavy_multiplication_x: |
| SupportedMimeTypes  | `mpris.MediaPlayer.SupportedMimeTypes() ([]string, error)`  | :heavy_multiplication_x: |

### Player

https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html

#### Methods

| method      | library path                                                            | implemented        |
|-------------|-------------------------------------------------------------------------|--------------------|
| Next        | `mpris.Player.Next()`                                                   | :heavy_check_mark: |
| Previous    | `mpris.Player.Previous()`                                               | :heavy_check_mark: |
| Pause       | `mpris.Player.Pause()`                                                  | :heavy_check_mark: |
| PlayPause   | `mpris.Player.PlayPause()`                                              | :heavy_check_mark: |
| Stop        | `mpris.Player.Stop()`                                                   | :heavy_check_mark: |
| Seek        | `mpris.Player.SeekTo(<offset> int64)`¹                                  | :heavy_check_mark: |
| SetPosition | `mpris.Player.SetPosition(<trackID> dbus.ObjectPath, <position> int64)` | :heavy_check_mark: |
| OpenUri     | `mpris.Player.OpenUri(<uri> string)`                                    | :heavy_check_mark: |

¹ Could not be named Seek, it's a reserved function name.

#### Properties

| property       | library path                                                            | implemented        |
|----------------|-------------------------------------------------------------------------|--------------------|
| PlaybackStatus | `mpris.Player.PlaybackStatus() (mpris.PlaybackStatus, error)`           | :heavy_check_mark: |
| LoopStatus     | `mpris.Player.LoopStatus() (mpris.LoopStatus, error)`                   | :heavy_check_mark: |
| LoopStatus     | `mpris.Player.SetLoopStatus(<loopStatus> mpris.LoopStatus) error`       | :heavy_check_mark: |
| Rate           | `mpris.Player.Rate() (float64, error)`                                  | :heavy_check_mark: |
| Rate           | `mpris.Player.SetRate(<rate> float64) error`                            | :heavy_check_mark: |
| Shuffle        | `mpris.Player.Shuffle() (bool, error)`                                  | :heavy_check_mark: |
| Shuffle        | `mpris.Player.SetShuffle(<shuffle> bool) error`                         | :heavy_check_mark: |
| Metadata       | `mpris.Player.Metadata() (mpris.Metadata, error)`                       | :heavy_check_mark: |
| Volume         | `mpris.Player.Volume() (float64, error)`                                | :heavy_check_mark: |
| Volume         | `mpris.Player.SetVolume(<volume> float64) (error)`                      | :heavy_check_mark: |
| Position       | `mpris.Player.Position() (int64, error)`                                | :heavy_check_mark: |
| Position       | `mpris.Player.SetPosition(<trackID> dbus.ObjectPath, <position> int64)` | :heavy_check_mark: |
| MinimumRate    | `mpris.Player.MinimumRate() (float64, error)`                           | :heavy_check_mark: |
| MaximumRate    | `mpris.Player.MaximumRate() (float64, error)`                           | :heavy_check_mark: |
| CanGoNext      | `mpris.Player.CanGoNext() (bool, error)`                                | :heavy_check_mark: |
| CanGoPrevious  | `mpris.Player.CanGoPrevious() (bool, error)`                            | :heavy_check_mark: |
| CanPlay        | `mpris.Player.CanPlay() (bool, error)`                                  | :heavy_check_mark: |
| CanPause       | `mpris.Player.CanPause() (bool, error)`                                 | :heavy_check_mark: |
| CanSeek        | `mpris.Player.CanSeek() (bool, error)`                                  | :heavy_check_mark: |
| CanControl     | `mpris.Player.CanControl() (bool, error)`                               | :heavy_check_mark: |

#### Signals

| signal | library path                                                      | implemented        |
|--------|-------------------------------------------------------------------|--------------------|
| Seeked | `mpris.Player.Seeked(<ctx> context.Context) (<-chan int, error) ` | :heavy_check_mark: |

### TrackList

https://specifications.freedesktop.org/mpris-spec/2.2/Track_List_Interface.html

#### Methods

| method            | library path        | implemented              |
|-------------------|---------------------|--------------------------|
| GetTracksMetadata | Not implemented yet | :heavy_multiplication_x: |
| AddTrack          | Not implemented yet | :heavy_multiplication_x: |
| RemoveTrack       | Not implemented yet | :heavy_multiplication_x: |
| GoTo              | Not implemented yet | :heavy_multiplication_x: |

#### Properties

| property      | library path        | implemented              |
|---------------|---------------------|--------------------------|
| Tracks        | Not implemented yet | :heavy_multiplication_x: |
| CanEditTracks | Not implemented yet | :heavy_multiplication_x: |

#### Signals

| signal               | library path        | implemented              |
|----------------------|---------------------|--------------------------|
| TrackListReplaced    | Not implemented yet | :heavy_multiplication_x: |
| TrackAdded           | Not implemented yet | :heavy_multiplication_x: |
| TrackRemoved         | Not implemented yet | :heavy_multiplication_x: |
| TrackMetadataChanged | Not implemented yet | :heavy_multiplication_x: |

## Development

### Versioning

This library follows the semantic versioning concept.

### Commits

Commits should follow the conventional commit rules.  
See: https://conventionalcommits.org.

### Mocks

Mocks will be generated with `github.com/matryer/moq`. It can be installed with
`go install github.com/matryer/moq@latest`. Generation can be triggered with `go generate ./...`.

### Go Docs

Read the docs at https://pkg.go.dev/github.com/leberKleber/go-mpris