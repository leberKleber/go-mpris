package mpris

import (
	"fmt"
	"github.com/godbus/dbus/v5"
)

const (
	mediaPlayerObjectPath  = "/org/mpris/MediaPlayer2"
	mediaPlayerInterface   = "org.mpris.MediaPlayer2"
	mediaPlayerRaiseMethod = mediaPlayerInterface + ".Raise"
	mediaPlayerQuitMethod  = mediaPlayerInterface + ".Quit"
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
func NewMediaPlayer(name string) (MediaPlayer, error) {
	connection, err := dbusSessionBus()
	if err != nil {
		return MediaPlayer{}, fmt.Errorf("failed to connect to session-bus: %w", err)
	}

	return NewMediaPlayerWithConnection(name, connection), nil
}

// NewMediaPlayerWithConnection returns a new MediaPlayer with the given name and connection.
// Deprecated: NewMediaPlayerWithConnection will be removed in the future.
// Plain Struct initialization should be used instead.
// Private fields will be public.
func NewMediaPlayerWithConnection(name string, connection *dbus.Conn) MediaPlayer {
	return MediaPlayer{
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

// Raise brings the media player's user interface to the front using any appropriate mechanism available.
// The media player may be unable to control how its user interface is displayed, or it may not have a graphical user interface at all.
// In this case, the CanRaise property is false and this method does nothing.
// see: https://specifications.freedesktop.org/mpris-spec/2.2/Media_Player.html#Method:Raise
func (mp MediaPlayer) Raise() {
	mp.connection.Object(mp.name, mediaPlayerObjectPath).Call(mediaPlayerRaiseMethod, 0)
}

// Quit causes the media player to stop running.
// The media player may refuse to allow clients to shut it down. In this case, the CanQuit property is false and this method does nothing.
// Note: Media players which can be D-Bus activated, or for which there is no sensibly easy way to terminate a running instance (via the main interface or a notification area icon for example) should allow clients to use this method.
// Otherwise, it should not be needed.
// If the media player does not have a UI, this should be implemented.
// see: https://specifications.freedesktop.org/mpris-spec/2.2/Media_Player.html#Method:Quit
func (mp MediaPlayer) Quit() {
	mp.connection.Object(mp.name, mediaPlayerObjectPath).Call(mediaPlayerQuitMethod, 0)
}
