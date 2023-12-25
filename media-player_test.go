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
