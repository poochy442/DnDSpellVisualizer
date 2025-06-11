package main

import (
	"DnDSpellVisualizer/dnd_api"
	"DnDSpellVisualizer/server"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "update":
		getSkills()
	case "run":
		runServer()
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
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

func getSkills() {
	spells, err := dnd_api.GetSpells()
	if err != nil {
		log.Fatalf("Failed to get spells: %v", err)
	}

	file, err := os.Create("spells.json")
	if err != nil {
		log.Fatalf("Failed to create spells.json: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(spells); err != nil {
		log.Fatalf("Failed to encode spells to JSON: %v", err)
	}

	fmt.Println("Successfully updated spells.json")
}

func runServer() {
	http.HandleFunc("/api/draw-spell", server.GenerateSpellHandler)

	port := "8080"
	fmt.Printf("Server is listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
