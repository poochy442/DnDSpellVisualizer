package main

import (
	"DnDSpellVisualizer/dnd_api"
	"DnDSpellVisualizer/server"
	"DnDSpellVisualizer/spell_visualizer"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		printUsage()
		os.Exit(1)
	}

	const (
		cmdUpdate   = "update"
		cmdRun      = "run"
		cmdHelp     = "help"
		defaultPort = "8080"
	)

	var err error
	switch os.Args[1] {
	case cmdUpdate:
		err = getSkills()
	case cmdRun:
		err = runServer(defaultPort)
	case cmdHelp:
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: main <command>")
	fmt.Println("Commands:")
	fmt.Println("	update	- Update the spells.json file")
	fmt.Println("	run		- Run the HTTP server")
	fmt.Println("	help	- Display this help message")
}

func getSkills() error {
	spells, err := dnd_api.GetSpells()
	if err != nil {
		return fmt.Errorf("failed to get spells: %w", err)
	}

	file, err := os.Create("spells.json")
	if err != nil {
		return fmt.Errorf("failed to create spells.json: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(spells); err != nil {
		return fmt.Errorf("failed to encode spells to JSON: %w", err)
	}

	fmt.Println("Successfully updated spells.json")
	return nil
}

func runServer(port string) error {
	spells, err := spell_visualizer.LoadSpells()
	if err != nil {
		return fmt.Errorf("failed to load spells: %w", err)
	}

	http.HandleFunc("/api/draw-spell", server.DrawSpellHandler(spells))

	fmt.Printf("Server is listening on localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return fmt.Errorf("server error: %w", err)
	}
	return nil
}
