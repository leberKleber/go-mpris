package main

import (
	"bufio"
	"fmt"
	"github.com/leberKleber/go-mpris"
	"os"
	"strconv"
	"strings"
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

		err = handleInput(p, reader, strings.Trim(input, "\n "))
		if err != nil {
			fmt.Printf("failed to exec command %q: %s\n", input, err)
			os.Exit(1)
		}
	}
}

func handleInput(p mpris.Player, reader *bufio.Reader, input string) error {
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
	case "seek":
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
	case "set-position":
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
	case "open-uri":
		fmt.Println("Open the given uri.")
		fmt.Print("-->")
		uri, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		p.OpenURI(strings.Trim(uri, "\n "))
	case "playback-status":
		s, err := p.PlaybackStatus()
		if err != nil {
			fmt.Printf("failed to get playback status: %s\n", err)
		}
		fmt.Println(s)
	case "loop-status":
		s, err := p.LoopStatus()
		if err != nil {
			fmt.Printf("failed to get loop status: %s\n", err)
		}
		fmt.Println(s)
	case "set loop-status":
		status, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		err = p.SetLoopStatus(strings.Trim(status, "\n "))
		if err != nil {
			fmt.Printf("failed to set loop status: %s\n", err)
		}
	case "rate":
		s, err := p.Rate()
		if err != nil {
			fmt.Printf("failed to get rate: %s\n", err)
		}
		fmt.Println(s)
	case "set rate":
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
	case "shuffle":
		s, err := p.Shuffle()
		if err != nil {
			fmt.Printf("failed to get shuffle: %s\n", err)
		}
		fmt.Println(s)
	case "set shuffle":
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
	case "metadata":
		s, err := p.Metadata()
		if err != nil {
			fmt.Printf("failed to get metadata: %s\n", err)
		}
		fmt.Println(s)

	case "volume":
		s, err := p.Volume()
		if err != nil {
			fmt.Printf("failed to get volume: %s\n", err)
		}
		fmt.Println(s)
	case "set volume":
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
	default:
		fmt.Println("Unknown command.")
		printHelp()
	}

	return nil
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
}
