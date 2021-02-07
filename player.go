package mpris

import (
	"fmt"
	"github.com/godbus/dbus/v5"
)

const (
	playerObjectPath             = "/org/mpris/MediaPlayer2"
	playerInterface              = "org.mpris.MediaPlayer2.Player"
	playerNextMethod             = playerInterface + ".Next"
	playerPreviousMethod         = playerInterface + ".Previous"
	playerPauseMethod            = playerInterface + ".Pause"
	playerPlayPauseMethod        = playerInterface + ".PlayPause"
	playerStopMethod             = playerInterface + ".Stop"
	playerPlayMethod             = playerInterface + ".Play"
	playerSeekMethod             = playerInterface + ".SeekTo"
	playerSetPositionMethod      = playerInterface + ".SetPosition"
	playerOpenURIMethod          = playerInterface + ".OpenURI"
	playerPlaybackStatusProperty = playerInterface + ".PlaybackStatus"
)

var dbusSessionBus = dbus.SessionBus

//go:generate moq -out dbus-conn_moq_test.go . dbusConn
type dbusConn interface {
	Object(dest string, path dbus.ObjectPath) dbusBusObject
}

//go:generate moq -out dbus-bus-object_moq_test.go . dbusBusObject
type dbusBusObject interface {
	Call(method string, flags dbus.Flags, args ...interface{}) dbusCall
	GetProperty(p string) (v dbus.Variant, e error)
}

//go:generate moq -out dbus-call_moq_test.go . dbusCall
type dbusCall interface {
	Store(retvalues ...interface{}) error
}

//Player is a implementation of dbus org.mpris.MediaPlayer2.Player. see: https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html
//Use NewPlayer to create a new instance with a connected session-bus via dbus.SessionBus.
//Use NewPlayerWithConnection when you want to use a self-configured dbus.Conn
type Player struct {
	name       string
	connection dbusConn
}

//NewPlayer returns a new Player which is already connected to session-bus via dbus.SessionBus.
func NewPlayer(name string) (Player, error) {
	connection, err := dbusSessionBus()
	if err != nil {
		return Player{}, fmt.Errorf("failed to connect to session-bus: %w", err)
	}

	return Player{
		name: name,
		connection: &dbusConnWrapper{
			conn: connection,
		},
	}, nil
}

//NewPlayer returns a new Player with the given name and connection.
func NewPlayerWithConnection(name string, connection *dbus.Conn) Player {
	return Player{
		name: name,
		connection: &dbusConnWrapper{
			conn: connection,
		},
	}
}

//Next skips to the next track in the tracklist.
//If there is no next track (and endless playback and track repeat are both off), stop playback.
//If playback is paused or stopped, it remains that way.
//If CanGoNext is false, attempting to call this method should have no effect.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:Next
func (p Player) Next() {
	p.connection.Object(p.name, playerObjectPath).Call(playerNextMethod, 0)
}

//Previous skips to the previous track in the tracklist.
//If there is no previous track (and endless playback and track repeat are both off), stop playback.
//If playback is paused or stopped, it remains that way.
//If CanGoPrevious is false, attempting to call this method should have no effect.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:Previous
func (p Player) Previous() {
	p.connection.Object(p.name, playerObjectPath).Call(playerPreviousMethod, 0)
}

//Pause pauses playback.
//If playback is already paused, this has no effect.
//Calling Play after this should cause playback to start again from the same position.
//If CanPause is false, attempting to call this method should have no effect.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:Pause
func (p Player) Pause() {
	p.connection.Object(p.name, playerObjectPath).Call(playerPauseMethod, 0)
}

//Play starts or resumes playback.
//If already playing, this has no effect.
//If paused, playback resumes from the current position.
//If there is no track to play, this has no effect.
//If CanPlay is false, attempting to call this method should have no effect.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:Play
func (p Player) Play() {
	p.connection.Object(p.name, playerObjectPath).Call(playerPlayMethod, 0)
}

//PlayPause pauses playback.
//If playback is already paused, resumes playback.
//If playback is stopped, starts playback.
//If CanPause is false, attempting to call this method should have no effect and raise an error.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:PlayPause
func (p Player) PlayPause() {
	p.connection.Object(p.name, playerObjectPath).Call(playerPlayPauseMethod, 0)
}

//Stop stops playback.
//If playback is already stopped, this has no effect.
//Calling Play after this should cause playback to start again from the beginning of the track.
//If CanControl is false, attempting to call this method should have no effect and raise an error.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:Stop
func (p Player) Stop() {
	p.connection.Object(p.name, playerObjectPath).Call(playerStopMethod, 0)
}

//SeekTo seeks forward in the current track by the specified number of microseconds.
//Parameters:
//- offset (The number of microseconds to seek forward.)
//A negative value seeks back. If this would mean seeking back further than the start of the track, the position is set to 0.
//If the value passed in would mean seeking beyond the end of the track, acts like a call to Next.
//If the CanSeek property is false, this has no effect.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:Seek
func (p Player) SeekTo(offset int64) {
	p.connection.Object(p.name, playerObjectPath).Call(playerSeekMethod, 0, offset)
}

//SetPosition Sets the current track position in microseconds.
//Parameters:
//- trackID (The currently playing track's identifier. If this does not match the id of the currently-playing track, the call is ignored as "stale".)
//- position (Track position in microseconds. This must be between 0 and <track_length>.)
//If this does not match the id of the currently-playing track, the call is ignored as "stale".
//If the Position argument is less than 0, do nothing.
//If the Position argument is greater than the track length, do nothing.
//If the CanSeek property is false, this has no effect.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:SetPosition
func (p Player) SetPosition(trackID dbus.ObjectPath, position int64) {
	p.connection.Object(p.name, playerObjectPath).Call(playerSetPositionMethod, 0, trackID, position)
}

//OpenURI opens the Uri given as an argument
//Parameters:
//- uri (Uri of the track to load. Its uri scheme should be an element of the org.mpris.MediaPlayer2.SupportedUriSchemes property and the mime-type should match one of the elements of the org.mpris.MediaPlayer2.SupportedMimeTypes.)
//If the playback is stopped, starts playing
//If the uri scheme or the mime-type of the uri to open is not supported, this method does nothing and may raise an error. In particular, if the list of available uri schemes is empty, this method may not be implemented.
//Clients should not assume that the Uri has been opened as soon as this method returns. They should wait until the mpris:trackid field in the Metadata property changes.
//If the media player implements the TrackList interface, then the opened track should be made part of the tracklist, the org.mpris.MediaPlayer2.TrackList.TrackAdded or org.mpris.MediaPlayer2.TrackList.TrackListReplaced signal should be fired, as well as the org.freedesktop.DBus.Properties.PropertiesChanged signal on the tracklist interface.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:OpenUri
func (p Player) OpenURI(uri string) {
	p.connection.Object(p.name, playerObjectPath).Call(playerOpenURIMethod, 0, uri)
}

//PlaybackStatus returns the current playback status.
//When this property changes, the org.freedesktop.DBus.Properties.PropertiesChanged signal is emitted with the new value.
//May be "Playing", "Paused" or "Stopped".
func (p Player) PlaybackStatus() (string, error) {
	v, err := p.connection.Object(p.name, playerObjectPath).GetProperty(playerPlaybackStatusProperty)
	if err != nil {
		return "", fmt.Errorf("failed to get property %q: %w", playerPlaybackStatusProperty, err)
	}

	return v.Value().(string), nil
}
