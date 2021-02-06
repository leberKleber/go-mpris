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
	p, err := mpris.NewPlayer("org.mpris.MediaPlayer2.mpv")
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
			return fmt.Errorf("failed to read offset: %w", err)
		}

		_, err = strconv.ParseInt(strings.Trim(offsetAsString, "\n "), 10, 64)
		if err != nil {
			fmt.Println("input must be a number")
			return nil
		}

		p.Seek(10000000)
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
}
