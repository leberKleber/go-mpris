package mpris

import (
	"errors"
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewMediaPlayer(t *testing.T) {
	oldDbusSessionBus := dbusSessionBus
	defer func() {
		dbusSessionBus = oldDbusSessionBus
	}()

	dbusConn := &dbus.Conn{}
	dbusSessionBus = func() (conn *dbus.Conn, err error) {
		return dbusConn, nil
	}

	p, err := NewMediaPlayer("test")

	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, &dbusConnWrapper{dbusConn}, p.connection)
}

func TestNewMediaPlayer_Error(t *testing.T) {
	oldDbusSessionBus := dbusSessionBus
	defer func() {
		dbusSessionBus = oldDbusSessionBus
	}()

	dbusSessionBusErr := errors.New("nope")
	dbusSessionBus = func() (conn *dbus.Conn, err error) {
		return nil, dbusSessionBusErr
	}

	expectedError := errors.New("failed to connect to session-bus: nope")
	_, err := NewMediaPlayer("test")
	if fmt.Sprint(err) == fmt.Sprint() {
		t.Fatalf("unexpected error. Given: %q, Expected: %q", err, expectedError)
	}
}

func TestNewMediaPlayerWithConnection(t *testing.T) {
	dbusConn := &dbus.Conn{}

	p := NewMediaPlayerWithConnection("test", dbusConn)
	assert.Equal(t, &dbusConnWrapper{dbusConn}, p.connection)
}

func TestMediaPlayer_Close(t *testing.T) {
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

			err := MediaPlayer{
				connection: &mock,
			}.Close()
			require.Equal(t, msgOrEmpty(err), tt.expectedErr)
			assert.Equal(t, 1, len(mock.CloseCalls()))
		})
	}
}

func TestMediaPlayer_Methods(t *testing.T) {
	tests := []struct {
		name           string
		givenName      string
		action         func(p *MediaPlayer)
		expectedDest   string
		expectedPath   dbus.ObjectPath
		expectedMethod string
		expectedFlags  dbus.Flags
		expectedArgs   []interface{}
	}{
		{
			name:      "Raise",
			givenName: "raise",
			action: func(p *MediaPlayer) {
				p.Raise()
			},
			expectedDest:   "raise",
			expectedPath:   "/org/mpris/MediaPlayer2",
			expectedMethod: "org.mpris.MediaPlayer2.Raise",
			expectedFlags:  0,
			expectedArgs:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var givenDest string
			var givenPath dbus.ObjectPath
			var givenMethod string
			var givenFlags dbus.Flags
			var givenArgs []interface{}

			tt.action(&MediaPlayer{
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
