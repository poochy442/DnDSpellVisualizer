package main

import (
	"DnDSpellVisualizer/spell_visualizer"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/draw-spell", spell_visualizer.GenerateSpellHandler)

	port := "8080"
	fmt.Printf("Server is listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
