package mpris

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlayer_Methods(t *testing.T) {
	tests := []struct {
		name           string
		givenName      string
		action         func(p *Player)
		expectedDest   string
		expectedPath   dbus.ObjectPath
		expectedMethod string
		expectedFlags  dbus.Flags
		expectedArgs   []interface{}
	}{
		{
			name:      "Next",
			givenName: "next",
			action: func(p *Player) {
				p.Next()
			},
			expectedDest:   "next",
			expectedPath:   "/org/mpris/MediaPlayer2",
			expectedMethod: "org.mpris.MediaPlayer2.Player.Next",
			expectedFlags:  0,
			expectedArgs:   nil,
		},
		{
			name:      "Previous",
			givenName: "previous",
			action: func(p *Player) {
				p.Previous()
			},
			expectedDest:   "previous",
			expectedPath:   "/org/mpris/MediaPlayer2",
			expectedMethod: "org.mpris.MediaPlayer2.Player.Previous",
			expectedFlags:  0,
			expectedArgs:   nil,
		},
		{
			name:      "Pause",
			givenName: "pause",
			action: func(p *Player) {
				p.Pause()
			},
			expectedDest:   "pause",
			expectedPath:   "/org/mpris/MediaPlayer2",
			expectedMethod: "org.mpris.MediaPlayer2.Player.Pause",
			expectedFlags:  0,
			expectedArgs:   nil,
		},
		{
			name:      "Play",
			givenName: "play",
			action: func(p *Player) {
				p.Play()
			},
			expectedDest:   "play",
			expectedPath:   "/org/mpris/MediaPlayer2",
			expectedMethod: "org.mpris.MediaPlayer2.Player.Play",
			expectedFlags:  0,
			expectedArgs:   nil,
		},
		{
			name:      "PlayPause",
			givenName: "play-pause",
			action: func(p *Player) {
				p.PlayPause()
			},
			expectedDest:   "play-pause",
			expectedPath:   "/org/mpris/MediaPlayer2",
			expectedMethod: "org.mpris.MediaPlayer2.Player.PlayPause",
			expectedFlags:  0,
			expectedArgs:   nil,
		},
		{
			name:      "SeekTo",
			givenName: "seek-to",
			action: func(p *Player) {
				p.SeekTo(12356789)
			},
			expectedDest:   "seek-to",
			expectedPath:   "/org/mpris/MediaPlayer2",
			expectedMethod: "org.mpris.MediaPlayer2.Player.SeekTo",
			expectedFlags:  0,
			expectedArgs:   []interface{}{int64(12356789)},
		},
		{
			name:      "Stop",
			givenName: "stop",
			action: func(p *Player) {
				p.Stop()
			},
			expectedDest:   "stop",
			expectedPath:   "/org/mpris/MediaPlayer2",
			expectedMethod: "org.mpris.MediaPlayer2.Player.Stop",
			expectedFlags:  0,
			expectedArgs:   nil,
		},
		{
			name:      "SetPosition",
			givenName: "set-position",
			action: func(p *Player) {
				p.SetPosition("/my/path", 123456789)
			},
			expectedDest:   "set-position",
			expectedPath:   "/org/mpris/MediaPlayer2",
			expectedMethod: "org.mpris.MediaPlayer2.Player.SetPosition",
			expectedFlags:  0,
			expectedArgs:   []interface{}{dbus.ObjectPath("/my/path"), int64(123456789)},
		},
		{
			name:      "OpenUri",
			givenName: "open-uri",
			action: func(p *Player) {
				p.OpenURI("file://my/uri")
			},
			expectedDest:   "open-uri",
			expectedPath:   "/org/mpris/MediaPlayer2",
			expectedMethod: "org.mpris.MediaPlayer2.Player.OpenUri",
			expectedFlags:  0,
			expectedArgs:   []interface{}{"file://my/uri"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var givenDest string
			var givenPath dbus.ObjectPath
			var givenMethod string
			var givenFlags dbus.Flags
			var givenArgs []interface{}

			tt.action(&Player{
				name: tt.givenName,
				connection: &dbusConnMock{
					ObjectFunc: func(dest string, path dbus.ObjectPath) dbusBusObject {
						givenDest = dest
						givenPath = path
						return &dbusBusObjectMock{
							CallFunc: func(method string, flags dbus.Flags, args ...interface{}) dbusCall {
								givenMethod = method
								givenFlags = flags
								givenArgs = args
								return nil
							},
						}
					},
				},
			})

			assert.Equal(t, tt.expectedDest, givenDest, "given dest is not as expected")
			assert.Equal(t, tt.expectedPath, givenPath, "given path is not as expected")
			assert.Equal(t, tt.expectedMethod, givenMethod, "given method is not as expected")
			assert.Equal(t, tt.expectedFlags, givenFlags, "given flags is not as expected")
			assert.EqualValues(t, tt.expectedArgs, givenArgs, "given args is not as expected")
		})
	}
}

func TestPlayer_GetProperties(t *testing.T) {
	tests := []struct {
		name           string
		givenName      string
		callVariant    dbus.Variant
		callError      error
		runAndValidate func(t *testing.T, p *Player)
		expectedDest   string
		expectedPath   dbus.ObjectPath
		expectedKey    string
	}{
		{
			name:        "PlaybackStatus",
			callVariant: dbus.MakeVariant("Paused"),
			givenName:   "playback-status",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.PlaybackStatus()
				assert.NoError(t, err)
				assert.Equal(t, PlaybackStatusPaused, s, "status is not as expected")
			},
			expectedDest: "playback-status",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.PlaybackStatus",
		},
		{
			name:      "PlaybackStatus error",
			callError: errors.New("nope"),
			givenName: "playback-status",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.PlaybackStatus()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.PlaybackStatus\": nope", fmt.Sprint(err))
			},
			expectedDest: "playback-status",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.PlaybackStatus",
		},
		{
			name:        "LoopStatus",
			callVariant: dbus.MakeVariant("Track"),
			givenName:   "loop-status",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.LoopStatus()
				assert.NoError(t, err)
				assert.Equal(t, LoopStatusTrack, s, "status is not as expected")
			},
			expectedDest: "loop-status",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.LoopStatus",
		},
		{
			name:      "LoopStatus error",
			callError: errors.New("nope"),
			givenName: "loop-status",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.LoopStatus()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.LoopStatus\": nope", fmt.Sprint(err))
			},
			expectedDest: "loop-status",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.LoopStatus",
		},
		{
			name:        "Rate",
			callVariant: dbus.MakeVariant(float64(3)),
			givenName:   "rate",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.Rate()
				assert.NoError(t, err)
				assert.Equal(t, float64(3), s, "rate is not as expected")
			},
			expectedDest: "rate",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.Rate",
		},
		{
			name:      "Rate error",
			callError: errors.New("nope"),
			givenName: "rate",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.Rate()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.Rate\": nope", fmt.Sprint(err))
			},
			expectedDest: "rate",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.Rate",
		},
		{
			name:        "Shuffle",
			callVariant: dbus.MakeVariant(false),
			givenName:   "shuffle",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.Shuffle()
				assert.NoError(t, err)
				assert.Equal(t, false, s, "shuffle is not as expected")
			},
			expectedDest: "shuffle",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.Shuffle",
		},
		{
			name:      "Shuffle error",
			callError: errors.New("nope"),
			givenName: "shuffle",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.Shuffle()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.Shuffle\": nope", fmt.Sprint(err))
			},
			expectedDest: "shuffle",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.Shuffle",
		},
		{
			name: "Metadata",
			callVariant: dbus.MakeVariant(map[string]dbus.Variant{
				"myKey1": dbus.MakeVariant(true),
				"myKey2": dbus.MakeVariant("key2"),
			}),
			givenName: "metadata",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.Metadata()
				assert.NoError(t, err)
				assert.Equal(t, Metadata{
					"myKey1": dbus.MakeVariant(true),
					"myKey2": dbus.MakeVariant("key2"),
				}, s, "metadata is not as expected")
			},
			expectedDest: "metadata",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.Metadata",
		},
		{
			name:      "Metadata error",
			callError: errors.New("nope"),
			givenName: "metadata",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.Metadata()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.Metadata\": nope", fmt.Sprint(err))
			},
			expectedDest: "metadata",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.Metadata",
		},
		{
			name:        "Volume",
			callVariant: dbus.MakeVariant(0.5),
			givenName:   "volume",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.Volume()
				assert.NoError(t, err)
				assert.Equal(t, 0.5, s, "volume is not as expected")
			},
			expectedDest: "volume",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.Volume",
		},
		{
			name:      "Volume error",
			callError: errors.New("nope"),
			givenName: "volume",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.Volume()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.Volume\": nope", fmt.Sprint(err))
			},
			expectedDest: "volume",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.Volume",
		},
		{
			name:        "Position",
			callVariant: dbus.MakeVariant(int64(220342)),
			givenName:   "position",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.Position()
				assert.NoError(t, err)
				assert.Equal(t, int64(220342), s, "position is not as expected")
			},
			expectedDest: "position",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.Position",
		},
		{
			name:      "Position error",
			callError: errors.New("nope"),
			givenName: "position",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.Position()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.Position\": nope", fmt.Sprint(err))
			},
			expectedDest: "position",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.Position",
		},
		{
			name:        "MinimumRate",
			callVariant: dbus.MakeVariant(0.000001),
			givenName:   "minimum-rate",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.MinimumRate()
				assert.NoError(t, err)
				assert.Equal(t, 0.000001, s, "minimum-rate is not as expected")
			},
			expectedDest: "minimum-rate",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.MinimumRate",
		},
		{
			name:      "MinimumRate error",
			callError: errors.New("nope"),
			givenName: "minimum-rate",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.MinimumRate()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.MinimumRate\": nope", fmt.Sprint(err))
			},
			expectedDest: "minimum-rate",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.MinimumRate",
		},
		{
			name:        "MaximumRate",
			callVariant: dbus.MakeVariant(0.000001),
			givenName:   "maximum-rate",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.MaximumRate()
				assert.NoError(t, err)
				assert.Equal(t, 0.000001, s, "maximum-rate is not as expected")
			},
			expectedDest: "maximum-rate",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.MaximumRate",
		},
		{
			name:      "MaximumRate error",
			callError: errors.New("nope"),
			givenName: "maximum-rate",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.MaximumRate()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.MaximumRate\": nope", fmt.Sprint(err))
			},
			expectedDest: "maximum-rate",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.MaximumRate",
		},
		{
			name:        "CanGoNext",
			callVariant: dbus.MakeVariant(true),
			givenName:   "can-go-next",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.CanGoNext()
				assert.NoError(t, err)
				assert.Equal(t, true, s, "can-go-next is not as expected")
			},
			expectedDest: "can-go-next",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanGoNext",
		},
		{
			name:      "CanGoNext error",
			callError: errors.New("nope"),
			givenName: "can-go-next",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.CanGoNext()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.CanGoNext\": nope", fmt.Sprint(err))
			},
			expectedDest: "can-go-next",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanGoNext",
		},
		{
			name:        "CanGoPrevious",
			callVariant: dbus.MakeVariant(true),
			givenName:   "can-go-previous",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.CanGoPrevious()
				assert.NoError(t, err)
				assert.Equal(t, true, s, "can-go-previous is not as expected")
			},
			expectedDest: "can-go-previous",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanGoPrevious",
		},
		{
			name:      "CanGoPrevious error",
			callError: errors.New("nope"),
			givenName: "can-go-previous",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.CanGoPrevious()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.CanGoPrevious\": nope", fmt.Sprint(err))
			},
			expectedDest: "can-go-previous",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanGoPrevious",
		},
		{
			name:        "CanPlay",
			callVariant: dbus.MakeVariant(true),
			givenName:   "can-play",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.CanPlay()
				assert.NoError(t, err)
				assert.Equal(t, true, s, "can-play is not as expected")
			},
			expectedDest: "can-play",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanPlay",
		},
		{
			name:      "CanPlay error",
			callError: errors.New("nope"),
			givenName: "can-play",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.CanPlay()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.CanPlay\": nope", fmt.Sprint(err))
			},
			expectedDest: "can-play",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanPlay",
		},
		{
			name:        "CanPause",
			callVariant: dbus.MakeVariant(true),
			givenName:   "can-pause",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.CanPause()
				assert.NoError(t, err)
				assert.Equal(t, true, s, "can-pause is not as expected")
			},
			expectedDest: "can-pause",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanPause",
		},
		{
			name:      "CanPause error",
			callError: errors.New("nope"),
			givenName: "can-pause",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.CanPause()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.CanPause\": nope", fmt.Sprint(err))
			},
			expectedDest: "can-pause",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanPause",
		},
		{
			name:        "CanSeek",
			callVariant: dbus.MakeVariant(true),
			givenName:   "can-seek",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.CanSeek()
				assert.NoError(t, err)
				assert.Equal(t, true, s, "can-seek is not as expected")
			},
			expectedDest: "can-seek",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanSeek",
		},
		{
			name:      "CanSeek error",
			callError: errors.New("nope"),
			givenName: "can-seek",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.CanSeek()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.CanSeek\": nope", fmt.Sprint(err))
			},
			expectedDest: "can-seek",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanSeek",
		},
		{
			name:        "CanControl",
			callVariant: dbus.MakeVariant(true),
			givenName:   "can-control",
			runAndValidate: func(t *testing.T, p *Player) {
				s, err := p.CanControl()
				assert.NoError(t, err)
				assert.Equal(t, true, s, "can-control is not as expected")
			},
			expectedDest: "can-control",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanControl",
		},
		{
			name:      "CanControl error",
			callError: errors.New("nope"),
			givenName: "can-control",
			runAndValidate: func(t *testing.T, p *Player) {
				_, err := p.CanControl()
				assert.Equal(t, "failed to get property \"org.mpris.MediaPlayer2.Player.CanControl\": nope", fmt.Sprint(err))
			},
			expectedDest: "can-control",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.CanControl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var calledDest string
			var calledPath dbus.ObjectPath
			var calledKey string

			m := &dbusConnMock{
				ObjectFunc: func(dest string, path dbus.ObjectPath) dbusBusObject {
					calledDest = dest
					calledPath = path
					return &dbusBusObjectMock{
						GetPropertyFunc: func(key string) (dbus.Variant, error) {
							calledKey = key
							return tt.callVariant, tt.callError
						},
					}
				},
			}

			tt.runAndValidate(t, &Player{name: tt.givenName, connection: m})

			assert.Equal(t, tt.expectedDest, calledDest, "called dest is not as expected")
			assert.Equal(t, tt.expectedPath, calledPath, "called path is not as expected")
			assert.Equal(t, tt.expectedKey, calledKey, "called key is not as expected")
		})
	}
}

func TestPlayer_SetProperties(t *testing.T) {
	tests := []struct {
		name             string
		givenName        string
		callError        error
		runAndValidate   func(t *testing.T, p *Player)
		expectedDest     string
		expectedPath     dbus.ObjectPath
		expectedProperty string
		expectedValue    interface{}
	}{
		{
			name:      "LoopStatus",
			givenName: "loop-status",
			runAndValidate: func(t *testing.T, p *Player) {
				err := p.SetLoopStatus(LoopStatusTrack)
				assert.NoError(t, err)
			},
			expectedDest:     "loop-status",
			expectedPath:     "/org/mpris/MediaPlayer2",
			expectedProperty: "org.mpris.MediaPlayer2.Player.LoopStatus",
			expectedValue:    dbus.MakeVariant("Track"),
		},
		{
			name:      "LoopStatus error",
			callError: errors.New("nope"),
			givenName: "loop-status",
			runAndValidate: func(t *testing.T, p *Player) {
				err := p.SetLoopStatus(LoopStatusPlaylist)
				assert.Equal(t, "failed to set property \"org.mpris.MediaPlayer2.Player.LoopStatus\": nope", fmt.Sprint(err))
			},
			expectedDest:     "loop-status",
			expectedPath:     "/org/mpris/MediaPlayer2",
			expectedProperty: "org.mpris.MediaPlayer2.Player.LoopStatus",
			expectedValue:    dbus.MakeVariant("Playlist"),
		},
		{
			name:      "Rate",
			givenName: "rate",
			runAndValidate: func(t *testing.T, p *Player) {
				err := p.SetRate(0.5)
				assert.NoError(t, err)
			},
			expectedDest:     "rate",
			expectedPath:     "/org/mpris/MediaPlayer2",
			expectedProperty: "org.mpris.MediaPlayer2.Player.Rate",
			expectedValue:    dbus.MakeVariant(0.5),
		},
		{
			name:      "Rate error",
			callError: errors.New("nope"),
			givenName: "rate",
			runAndValidate: func(t *testing.T, p *Player) {
				err := p.SetRate(2)
				assert.Equal(t, "failed to set property \"org.mpris.MediaPlayer2.Player.Rate\": nope", fmt.Sprint(err))
			},
			expectedDest:     "rate",
			expectedPath:     "/org/mpris/MediaPlayer2",
			expectedProperty: "org.mpris.MediaPlayer2.Player.Rate",
			expectedValue:    dbus.MakeVariant(float64(2)),
		},
		{
			name:      "Shuffle",
			givenName: "shuffle",
			runAndValidate: func(t *testing.T, p *Player) {
				err := p.SetShuffle(true)
				assert.NoError(t, err)
			},
			expectedDest:     "shuffle",
			expectedPath:     "/org/mpris/MediaPlayer2",
			expectedProperty: "org.mpris.MediaPlayer2.Player.Shuffle",
			expectedValue:    dbus.MakeVariant(true),
		},
		{
			name:      "Shuffle error",
			callError: errors.New("nope"),
			givenName: "shuffle",
			runAndValidate: func(t *testing.T, p *Player) {
				err := p.SetShuffle(false)
				assert.Equal(t, "failed to set property \"org.mpris.MediaPlayer2.Player.Shuffle\": nope", fmt.Sprint(err))
			},
			expectedDest:     "shuffle",
			expectedPath:     "/org/mpris/MediaPlayer2",
			expectedProperty: "org.mpris.MediaPlayer2.Player.Shuffle",
			expectedValue:    dbus.MakeVariant(false),
		},
		{
			name:      "Volume",
			givenName: "volume",
			runAndValidate: func(t *testing.T, p *Player) {
				err := p.SetVolume(1)
				assert.NoError(t, err)
			},
			expectedDest:     "volume",
			expectedPath:     "/org/mpris/MediaPlayer2",
			expectedProperty: "org.mpris.MediaPlayer2.Player.Volume",
			expectedValue:    dbus.MakeVariant(float64(1)),
		},
		{
			name:      "Volume error",
			callError: errors.New("nope"),
			givenName: "volume",
			runAndValidate: func(t *testing.T, p *Player) {
				err := p.SetVolume(0.5)
				assert.Equal(t, "failed to set property \"org.mpris.MediaPlayer2.Player.Volume\": nope", fmt.Sprint(err))
			},
			expectedDest:     "volume",
			expectedPath:     "/org/mpris/MediaPlayer2",
			expectedProperty: "org.mpris.MediaPlayer2.Player.Volume",
			expectedValue:    dbus.MakeVariant(0.5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var calledDest string
			var calledPath dbus.ObjectPath
			var calledProperty string
			var calledValue interface{}

			m := &dbusConnMock{
				ObjectFunc: func(dest string, path dbus.ObjectPath) dbusBusObject {
					calledDest = dest
					calledPath = path
					return &dbusBusObjectMock{
						SetPropertyFunc: func(p string, v interface{}) error {
							calledProperty = p
							calledValue = v

							return tt.callError
						},
					}
				},
			}

			tt.runAndValidate(t, &Player{name: tt.givenName, connection: m})

			assert.Equal(t, tt.expectedDest, calledDest, "called dest is not as expected")
			assert.Equal(t, tt.expectedPath, calledPath, "called path is not as expected")
			assert.Equal(t, tt.expectedProperty, calledProperty, "called property is not as expected")
			assert.Equal(t, tt.expectedValue, calledValue, "called value is not as expected")
		})
	}
}

func TestPlayer_Close(t *testing.T) {
	tests := []struct {
		name        string
		closeErr    error
		expectedErr string
	}{
		{
			name: "Happycase",
		}, {
			name:        "Close error",
			closeErr:    errors.New("unexpected error"),
			expectedErr: "failed to close dbus connection: unexpected error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := dbusConnMock{
				CloseFunc: func() error {
					return tt.closeErr
				},
			}

			err := Player{
				connection: &mock,
			}.Close()
			require.Equal(t, msgOrEmpty(err), tt.expectedErr)
			assert.Equal(t, 1, len(mock.CloseCalls()))
		})
	}
}

func TestPlayer_Seeked(t *testing.T) {
	tests := []struct {
		name string

		givenSignals []*dbus.Signal

		addMatchSignalErr               error
		expectedAddMatchSignalCallCount int
		expectedSignalCallCount         int

		expectedErr       string
		expectedPositions []int
	}{
		{
			name:                            "happycase",
			expectedAddMatchSignalCallCount: 1,
			expectedSignalCallCount:         1,
			givenSignals: []*dbus.Signal{
				{
					Name: "org.mpris.MediaPlayer2.Player.Seeked",
					Body: []interface{}{int64(1111)},
				}, {
					Name: "unknown name",
					Body: []interface{}{int64(22222)},
				}, { // multiple body infos
					Name: "org.mpris.MediaPlayer2.Player.Seeked",
					Body: []interface{}{int64(333111), int64(333222)},
				}, { // invalid body type
					Name: "org.mpris.MediaPlayer2.Player.Seeked",
					Body: []interface{}{"4444444"},
				}, {
					Name: "org.mpris.MediaPlayer2.Player.Seeked",
					Body: []interface{}{int64(55555555)},
				},
			},
			expectedPositions: []int{
				1111,
				55555555,
			},
		}, {
			name:                            "add match signal error",
			addMatchSignalErr:               errors.New("unexpected error"),
			expectedAddMatchSignalCallCount: 1,
			expectedSignalCallCount:         0,
			expectedErr:                     "failed to add signal match option: unexpected error",
			givenSignals: []*dbus.Signal{
				{
					Name: "org.mpris.MediaPlayer2.Player.Seeked",
					Body: []interface{}{int64(1111)},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			mock := &dbusConnMock{
				AddMatchSignalFunc: func(_ ...dbus.MatchOption) error {
					return tt.addMatchSignalErr
				},
				SignalFunc: func(ch chan<- *dbus.Signal) {
					go func() {
						for _, sig := range tt.givenSignals {
							ch <- sig
						}
					}()
				},
			}

			poss, err := Player{
				connection: mock,
			}.Seeked(testCtx)
			require.Equal(t, tt.expectedErr, msgOrEmpty(err))

			var collectedPoss []int
			if poss != nil {
				for {
					p, ok := <-poss
					if !ok {
						break
					}
					collectedPoss = append(collectedPoss, p)
				}
			}
			assert.EqualValues(t, tt.expectedPositions, collectedPoss)
			assert.Equal(t, tt.expectedAddMatchSignalCallCount, len(mock.AddMatchSignalCalls()))
			assert.Equal(t, tt.expectedSignalCallCount, len(mock.SignalCalls()))
		})
	}
}

func TestNewPlayer(t *testing.T) {
	oldDbusSessionBus := dbusSessionBus
	defer func() {
		dbusSessionBus = oldDbusSessionBus
	}()

	dbusConn := &dbus.Conn{}
	dbusSessionBus = func() (conn *dbus.Conn, err error) {
		return dbusConn, nil
	}

	p, err := NewPlayer("test")

	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, &dbusConnWrapper{dbusConn}, p.connection)
}

func TestNewPlayer_Error(t *testing.T) {
	oldDbusSessionBus := dbusSessionBus
	defer func() {
		dbusSessionBus = oldDbusSessionBus
	}()

	dbusSessionBusErr := errors.New("nope")
	dbusSessionBus = func() (conn *dbus.Conn, err error) {
		return nil, dbusSessionBusErr
	}

	expectedError := errors.New("failed to connect to session-bus: nope")
	_, err := NewPlayer("test")
	if fmt.Sprint(err) == fmt.Sprint() {
		t.Fatalf("unexpected error. Given: %q, Expected: %q", err, expectedError)
	}
}

func TestNewPlayerWithConnection(t *testing.T) {
	dbusConn := &dbus.Conn{}

	p := NewPlayerWithConnection("test", dbusConn)
	assert.Equal(t, &dbusConnWrapper{dbusConn}, p.connection)
}
