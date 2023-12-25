package mpris

import (
	"fmt"
	"github.com/godbus/dbus/v5"
)

// MediaPlayer is an implementation of dbus org.mpris.MediaPlayer2.
// See: https://specifications.freedesktop.org/mpris-spec/2.2/Media_Player.html.
// Use NewMediaPlayer to create a new instance with a connected session-bus via dbus.SessionBus.
// Use NewMediaPlayerWithConnection when you want to use a self-configured dbus.Conn.
type MediaPlayer struct {
	name       string
	connection dbusConn
}

// NewMediaPlayer returns a new MediaPlayer which is already connected to session-bus via dbus.SessionBus.
// Don't forget to MediaPlayer.Close() the MediaPlayer after use.
func NewMediaPlayer(name string) (Player, error) {
	connection, err := dbusSessionBus()
	if err != nil {
		return Player{}, fmt.Errorf("failed to connect to session-bus: %w", err)
	}

	return NewPlayerWithConnection(name, connection), nil
}

// NewMediaPlayerWithConnection returns a new MediaPlayer with the given name and connection.
// Deprecated: NewMediaPlayerWithConnection will be removed in the future.
// Plain Struct initialization should be used instead.
// Private fields will be public.
func NewMediaPlayerWithConnection(name string, connection *dbus.Conn) Player {
	return Player{
		name: name,
		connection: &dbusConnWrapper{
			conn: connection,
		},
	}
}

// Close closes the dbus connection.
func (mp MediaPlayer) Close() error {
	err := mp.connection.Close()
	if err != nil {
		return fmt.Errorf("failed to close dbus connection: %w", err)
	}

	return nil
}
