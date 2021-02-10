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
	playerLoopStatusProperty     = playerInterface + ".LoopStatus"
	playerRateProperty           = playerInterface + ".Rate"
	playerShuffleProperty        = playerInterface + ".Shuffle"
	playerMetadataProperty       = playerInterface + ".Metadata"
	playerVolumeProperty         = playerInterface + ".Volume"
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
	SetProperty(p string, v interface{}) error
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
//May be "Playing", "Paused" or "Stopped".
//https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html#Property:PlaybackStatus
func (p Player) PlaybackStatus() (string, error) {
	v, err := p.getProperty(playerPlaybackStatusProperty)
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

//LoopStatus returns the current loop / repeat status
//May be:
//"None" if the playback will stop when there are no more tracks to play
//"Track" if the current track will start again from the beginning once it has finished playing
//"Playlist" if the playback loops through a list of tracks
//If CanControl is false, attempting to set this property (SetLoopStatus) should have no effect and raise an error.
//https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html#Property:LoopStatus
func (p Player) LoopStatus() (string, error) {
	v, err := p.getProperty(playerLoopStatusProperty)
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

//SetLoopStatus sets the current loop / repeat status
//May be:
//"None" if the playback will stop when there are no more tracks to play
//"Track" if the current track will start again from the beginning once it has finished playing
//"Playlist" if the playback loops through a list of tracks
//If CanControl is false, attempting to set this property (SetLoopStatus) should have no effect and raise an error.
//see: https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html#Property:LoopStatus
func (p Player) SetLoopStatus(status string) error {
	return p.setProperty(playerLoopStatusProperty, status)
}

//Rate return the current playback rate.
//The value must fall in the range described by MinimumRate and MaximumRate, and must not be 0.0. If playback is paused, the PlaybackStatus property should be used to indicate this. A value of 0.0 should not be set by the client. If it is, the media player should act as though Pause was called.
//If the media player has no ability to play at speeds other than the normal playback rate, this must still be implemented, and must return 1.0. The MinimumRate and MaximumRate properties must also be set to 1.0.
//Not all values may be accepted by the media player. It is left to media player implementations to decide how to deal with values they cannot use; they may either ignore them or pick a "best fit" value. Clients are recommended to only use sensible fractions or multiples of 1 (eg: 0.5, 0.25, 1.5, 2.0, etc).
//see: https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html#Property:Rate
func (p Player) Rate() (float64, error) {
	v, err := p.getProperty(playerRateProperty)
	if err != nil {
		return 0, err
	}
	return v.Value().(float64), nil
}

//SetRate sets the current playback rate.
//The value must fall in the range described by MinimumRate and MaximumRate, and must not be 0.0. If playback is paused, the PlaybackStatus property should be used to indicate this. A value of 0.0 should not be set by the client. If it is, the media player should act as though Pause was called.
//If the media player has no ability to play at speeds other than the normal playback rate, this must still be implemented, and must return 1.0. The MinimumRate and MaximumRate properties must also be set to 1.0.
//Not all values may be accepted by the media player. It is left to media player implementations to decide how to deal with values they cannot use; they may either ignore them or pick a "best fit" value. Clients are recommended to only use sensible fractions or multiples of 1 (eg: 0.5, 0.25, 1.5, 2.0, etc).
//see: https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html#Property:Rate
func (p Player) SetRate(rate float64) error {
	return p.setProperty(playerRateProperty, rate)
}

//Shuffle returns a value of false indicates that playback is progressing linearly through a playlist, while true means playback is progressing through a playlist in some other order.
//If CanControl is false, attempting to set this property should have no effect and raise an error.
//see: https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html#Property:Shuffle
func (p Player) Shuffle() (bool, error) {
	v, err := p.getProperty(playerShuffleProperty)
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

//Shuffle set a value of false indicates that playback is progressing linearly through a playlist, while true means playback is progressing through a playlist in some other order.
//If CanControl is false, attempting to set this property should have no effect and raise an error.
//see: https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html#Property:Shuffle
func (p Player) SetShuffle(shuffle bool) error {
	return p.setProperty(playerShuffleProperty, shuffle)
}

//Metadata metadata of the current element.
//If there is a current track, this must have a "mpris:trackid" entry (of D-Bus type "o") at the very least, which contains a D-Bus path that uniquely identifies this track.
//See the type documentation for more details.
//see: https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html#Property:Metadata
func (p Player) Metadata() (map[string]dbus.Variant, error) {
	v, err := p.getProperty(playerMetadataProperty)
	if err != nil {
		return nil, err
	}
	return v.Value().(map[string]dbus.Variant), nil
}

//Volume returns the volume level.
//When setting, if a negative value is passed, the volume should be set to 0.0.
//If CanControl is false, attempting to set this property should have no effect and raise an error.
//see: https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html#Property:Volume
func (p Player) Volume() (float64, error) {
	v, err := p.getProperty(playerVolumeProperty)
	if err != nil {
		return 0, err
	}
	return v.Value().(float64), nil
}

//SetVolume sets the volume level.
//When setting, if a negative value is passed, the volume should be set to 0.0.
//If CanControl is false, attempting to set this property should have no effect and raise an error.
//see: https://specifications.freedesktop.org/mpris-spec/2.2/Player_Interface.html#Property:Volume
func (p Player) SetVolume(volume float64) error {
	return p.setProperty(playerVolumeProperty, volume)
}

func (p Player) getProperty(property string) (dbus.Variant, error) {
	v, err := p.connection.Object(p.name, playerObjectPath).GetProperty(property)
	if err != nil {
		return dbus.Variant{}, fmt.Errorf("failed to get property %q: %w", property, err)
	}

	return v, nil
}

func (p Player) setProperty(property string, value interface{}) error {
	err := p.connection.Object(p.name, playerObjectPath).SetProperty(property, dbus.MakeVariant(value))
	if err != nil {
		return fmt.Errorf("failed to set property %q: %w", property, err)
	}

	return nil
}
