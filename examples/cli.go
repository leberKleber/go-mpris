package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/leberKleber/go-mpris"
)

func main() {
	p, err := mpris.NewPlayer("org.mpris.MediaPlayer2.vlc")
	if err != nil {
		fmt.Printf("failed to create gompris.Player: %s\n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("mpris example client with mpv (https://mpv.io) and mpv-mpris (https://github.com/hoyon/mpv-mpris)")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("failed to read input: %s\n", err)
			os.Exit(1)
		}

		handleInput(p, reader, strings.Trim(input, "\n "))
	}
}

func handleSetInput(p mpris.Player, reader *bufio.Reader, input string) error {
	var err error
	switch input {
	case "set-position":
		err = handleSetPosition(p, reader)
	case "set loop-status":
		err = handleSetLoopStatus(p, reader)
	case "set rate":
		err = handleSetRate(p, reader)
	case "set shuffle":
		err = handleSetShuffle(p, reader)
	case "set volume":
		err = handleSetVolume(p, reader)
	case "seek":
		err = handleSeek(p, reader)
	case "open-uri":
		err = handleOpenURI(p, reader)
	}

	return err
}

func handleOpenURI(p mpris.Player, reader *bufio.Reader) error {
	fmt.Println("Open the given uri.")
	fmt.Print("-->")
	uri, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	p.OpenURI(strings.Trim(uri, "\n "))
	return nil
}

func handleSeek(p mpris.Player, reader *bufio.Reader) error {
	fmt.Println("The number of microseconds to seek forward")
	fmt.Print("-->")
	offsetAsString, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	offset, err := strconv.ParseInt(strings.Trim(offsetAsString, "\n "), 10, 64)
	if err != nil {
		fmt.Println("input must be a number")
		return nil
	}

	p.SeekTo(offset)
	return nil
}

func handleSetVolume(p mpris.Player, reader *bufio.Reader) error {
	volumeAsString, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	shuffle, err := strconv.ParseFloat(strings.Trim(volumeAsString, "\n "), 10)
	if err != nil {
		fmt.Println("input must be a float64")
	}

	err = p.SetVolume(shuffle)
	if err != nil {
		fmt.Printf("failed to set volume: %s\n", err)
	}
	return nil
}

func handleSetShuffle(p mpris.Player, reader *bufio.Reader) error {
	shuffleAsString, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	shuffle, err := strconv.ParseBool(strings.Trim(shuffleAsString, "\n "))
	if err != nil {
		fmt.Println("input must be a bool")
	}

	err = p.SetShuffle(shuffle)
	if err != nil {
		fmt.Printf("failed to set shuffle: %s\n", err)
	}
	return nil
}

func handleSetRate(p mpris.Player, reader *bufio.Reader) error {
	rateAsString, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	rate, err := strconv.ParseFloat(strings.Trim(rateAsString, "\n "), 10)
	if err != nil {
		fmt.Println("input must be a float64")
	}

	err = p.SetRate(rate)
	if err != nil {
		fmt.Printf("failed to set rate: %s\n", err)
	}
	return nil
}

func handleSetLoopStatus(p mpris.Player, reader *bufio.Reader) error {
	status, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	err = p.SetLoopStatus(mpris.LoopStatus(strings.Trim(status, "\n ")))
	if err != nil {
		fmt.Printf("failed to set loop status: %s\n", err)
	}
	return nil
}

func handleSetPosition(p mpris.Player, reader *bufio.Reader) error {
	fmt.Println("The track position in microseconds.")
	fmt.Print("-->")
	offsetAsString, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	position, err := strconv.ParseInt(strings.Trim(offsetAsString, "\n "), 10, 64)
	if err != nil {
		fmt.Println("input must be a number")
		return nil
	}

	p.SetPosition("/not/used", position)
	return nil
}

func handleGetInput(p mpris.Player, input string) {
	switch input {
	case "playback-status":
		handleGetPlaybackStatus(p)
	case "loop-status":
		handleGetLoopStatus(p)
	case "rate":
		handleGetRate(p)
	case "shuffle":
		handleGetShuffle(p)
	case "metadata":
		handleGetMetadata(p)
	case "volume":
		handleGetVolume(p)
	case "position":
		handleGetPosition(p)
	case "minimum-rate":
		handleGetMinimumRate(p)
	case "maximum-rate":
		handleExtMaximumRate(p)
	}
}

func handleGetPlaybackStatus(p mpris.Player) {
	s, err := p.PlaybackStatus()
	if err != nil {
		fmt.Printf("failed to get playback status: %s\n", err)
	}
	fmt.Println(s)
}

func handleGetLoopStatus(p mpris.Player) {
	s, err := p.LoopStatus()
	if err != nil {
		fmt.Printf("failed to get loop status: %s\n", err)
	}
	fmt.Println(s)
}

func handleGetRate(p mpris.Player) {
	s, err := p.Rate()
	if err != nil {
		fmt.Printf("failed to get rate: %s\n", err)
	}
	fmt.Println(s)
}

func handleGetShuffle(p mpris.Player) {
	s, err := p.Shuffle()
	if err != nil {
		fmt.Printf("failed to get shuffle: %s\n", err)
	}
	fmt.Println(s)
}

func handleGetMetadata(p mpris.Player) {
	s, err := p.Metadata()
	if err != nil {
		fmt.Printf("failed to get metadata: %s\n", err)
	}
	fmt.Println(s)
}

func handleGetVolume(p mpris.Player) {
	s, err := p.Volume()
	if err != nil {
		fmt.Printf("failed to get volume: %s\n", err)
	}
	fmt.Println(s)
}

func handleGetPosition(p mpris.Player) {
	s, err := p.Position()
	if err != nil {
		fmt.Printf("failed to get position: %s\n", err)
	}
	fmt.Println(s)
}

func handleGetMinimumRate(p mpris.Player) {
	s, err := p.MinimumRate()
	if err != nil {
		fmt.Printf("failed to get minimum-rate: %s\n", err)
	}
	fmt.Println(s)
}

func handleExtMaximumRate(p mpris.Player) {
	s, err := p.MaximumRate()
	if err != nil {
		fmt.Printf("failed to get maximum-rate: %s\n", err)
	}
	fmt.Println(s)
}

func handleCan(p mpris.Player, input string) {
	switch input {
	case "can-go-next":
		s, err := p.CanGoNext()
		if err != nil {
			fmt.Printf("failed to get can-go-next: %s\n", err)
		}
		fmt.Println(s)
	case "can-go-previous":
		s, err := p.CanGoNext()
		if err != nil {
			fmt.Printf("failed to get can-go-previous: %s\n", err)
		}
		fmt.Println(s)
	case "can-play":
		s, err := p.CanPlay()
		if err != nil {
			fmt.Printf("failed to get can-play: %s\n", err)
		}
		fmt.Println(s)
	case "can-pause":
		s, err := p.CanPause()
		if err != nil {
			fmt.Printf("failed to get can-pause: %s\n", err)
		}
		fmt.Println(s)
	case "can-seek":
		s, err := p.CanSeek()
		if err != nil {
			fmt.Printf("failed to get can-seek: %s\n", err)
		}
		fmt.Println(s)
	case "can-control":
		s, err := p.CanControl()
		if err != nil {
			fmt.Printf("failed to get can-control: %s\n", err)
		}
		fmt.Println(s)
	}
}

func handleInput(p mpris.Player, reader *bufio.Reader, input string) {
	handleGetInput(p, input)
	handleCan(p, input)

	err := handleSetInput(p, reader, input)
	if err != nil {
		fmt.Printf("failed to handle set input: %s\n", err)
	}

	switch input {
	case "help":
		printHelp()
	case "quit":
		os.Exit(0)
	case "next":
		p.Next()
	case "previous":
		p.Next()
	case "pause":
		p.Pause()
	case "play-pause":
		p.PlayPause()
	case "stop":
		p.Stop()
	case "play":
		p.Play()
	}
}

func printHelp() {
	fmt.Println("Available commands")
	fmt.Println("- quit (exit this cli)")
	fmt.Println("- next")
	fmt.Println("- previous")
	fmt.Println("- pause")
	fmt.Println("- play-pause")
	fmt.Println("- stop")
	fmt.Println("- play")
	fmt.Println("- seek")
	fmt.Println("- set-position")
	fmt.Println("- open-uri")
	fmt.Println("- playback-status")
	fmt.Println("- loop-status")
	fmt.Println("- set-loop-status")
	fmt.Println("- rate")
	fmt.Println("- set-rate")
	fmt.Println("- metadata")
	fmt.Println("- volume")
	fmt.Println("- set-volume")
	fmt.Println("- position")
	fmt.Println("- minimum-rate")
	fmt.Println("- maximum-rate")
	fmt.Println("- can-go-next")
	fmt.Println("- can-go-previous")
	fmt.Println("- can-play")
	fmt.Println("- can-pause")
	fmt.Println("- can-seek")
	fmt.Println("- can-control")
}
