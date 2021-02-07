package mpris

import (
	"errors"
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/stretchr/testify/assert"
	"testing"
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
			name:      "OpenURI",
			givenName: "open-uri",
			action: func(p *Player) {
				p.OpenURI("file://my/uri")
			},
			expectedDest:   "open-uri",
			expectedPath:   "/org/mpris/MediaPlayer2",
			expectedMethod: "org.mpris.MediaPlayer2.Player.OpenURI",
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

func TestPlayer_Properties(t *testing.T) {
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
				assert.Equal(t, "Paused", s, "status is not as expected")
			},
			expectedDest: "playback-status",
			expectedPath: "/org/mpris/MediaPlayer2",
			expectedKey:  "org.mpris.MediaPlayer2.Player.PlaybackStatus",
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
