package mpris

import (
	"fmt"
	"github.com/godbus/dbus/v5"
)

const (
	playerObjectPath      = "/org/mpris/MediaPlayer2"
	playerInterface       = "org.mpris.MediaPlayer2.Player"
	playerNextMethod      = playerInterface + ".Next"
	playerPreviousMethod  = playerInterface + ".Previous"
	playerPauseMethod     = playerInterface + ".Pause"
	playerPlayPauseMethod = playerInterface + ".PlayPause"
	playerStopMethod      = playerInterface + ".Stop"
	playerPlayMethod      = playerInterface + ".Play"
)

type Player struct {
	Name       string
	Connection *dbus.Conn
}

//NewPlayer returns a new Player which is already connected to session-bus via dbus.SessionBus. Create your own Player
//instance if you want to use an other bus
func NewPlayer(name string) (Player, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return Player{}, fmt.Errorf("failed to connect to session-bus")
	}

	return Player{
		Name:       name,
		Connection: conn,
	}, nil
}

//Next skips to the next track in the tracklist.
//If there is no next track (and endless playback and track repeat are both off), stop playback.
//If playback is paused or stopped, it remains that way.
//If CanGoNext is false, attempting to call this method should have no effect.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:Next
func (p Player) Next() {
	p.Connection.Object(p.Name, playerObjectPath).Call(playerNextMethod, 0)
}

//Previous skips to the previous track in the tracklist.
//If there is no previous track (and endless playback and track repeat are both off), stop playback.
//If playback is paused or stopped, it remains that way.
//If CanGoPrevious is false, attempting to call this method should have no effect.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:Previous
func (p Player) Previous() {
	p.Connection.Object(p.Name, playerObjectPath).Call(playerPreviousMethod, 0)
}

//Pause pauses playback.
//If playback is already paused, this has no effect.
//Calling Play after this should cause playback to start again from the same position.
//If CanPause is false, attempting to call this method should have no effect.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:Pause
func (p Player) Pause() {
	p.Connection.Object(p.Name, playerObjectPath).Call(playerPauseMethod, 0)
}

//PlayPause pauses playback.
//If playback is already paused, resumes playback.
//If playback is stopped, starts playback.
//If CanPause is false, attempting to call this method should have no effect and raise an error.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:PlayPause
func (p Player) PlayPause() {
	p.Connection.Object(p.Name, playerObjectPath).Call(playerPlayPauseMethod, 0)
}

//Stop stops playback.
//If playback is already stopped, this has no effect.
//Calling Play after this should cause playback to start again from the beginning of the track.
//If CanControl is false, attempting to call this method should have no effect and raise an error.
//see: https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html#Method:Stop
func (p Player) Stop() {
	p.Connection.Object(p.Name, playerObjectPath).Call(playerStopMethod, 0)
}

//Play starts or resumes playback.
//If already playing, this has no effect.
//If paused, playback resumes from the current position.
//If there is no track to play, this has no effect.
//If CanPlay is false, attempting to call this method should have no effect.
func (p Player) Play() {
	p.Connection.Object(p.Name, playerObjectPath).Call(playerPlayMethod, 0)
}
