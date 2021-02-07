package mpris

import (
	"errors"
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayer(t *testing.T) {
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
		}, {
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
		}, {
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
		}, {
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
		}, {
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
		}, {
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
		}, {
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
		}, {
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
		}, {
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
			dbusWrapperMock := &dbusWrapperIMock{}

			var givenDest string
			var givenPath dbus.ObjectPath
			var givenMethod string
			var givenFlags dbus.Flags
			var givenArgs []interface{}

			dbusWrapperMock.CallFunc = func(dest string, path dbus.ObjectPath, method string, flags dbus.Flags, args ...interface{}) {
				givenDest = dest
				givenPath = path
				givenMethod = method
				givenFlags = flags
				givenArgs = args
			}

			tt.action(&Player{
				name:       tt.givenName,
				connection: dbusWrapperMock,
			})

			assert.Equal(t, tt.expectedDest, givenDest, "given dest is not as expected")
			assert.Equal(t, tt.expectedPath, givenPath, "given path is not as expected")
			assert.Equal(t, tt.expectedMethod, givenMethod, "given method is not as expected")
			assert.Equal(t, tt.expectedFlags, givenFlags, "given flags is not as expected")
			assert.EqualValues(t, tt.expectedArgs, givenArgs, "given args is not as expected")
		})
	}
}

func TestNewPlayer(t *testing.T) {
	oldDbusSessionBus := dbusSessionBus
	defer func() {
		dbusSessionBus = oldDbusSessionBus
	}()

	dbConn := &dbus.Conn{}
	dbusSessionBus = func() (conn *dbus.Conn, err error) {
		return dbConn, nil
	}

	p, err := NewPlayer("test")
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, dbusWrapper{conn: dbConn}, p.connection)
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
	dbConn := &dbus.Conn{}

	p, err := NewPlayerWithConnection("test", dbConn)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, dbusWrapper{conn: dbConn}, p.connection)

}
