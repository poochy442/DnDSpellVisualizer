package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"DnDSpellVisualizer/spell_visualizer"
	"DnDSpellVisualizer/spell_visualizer/shared"
	"DnDSpellVisualizer/spell_visualizer/visualizer"
)

func GetSpellByName(name string) *shared.Spell {
	spell := spell_visualizer.GetSpellByName(name)
	if spell == nil {
		return nil
	}
	return spell
}

func validateConfig(config *shared.VisualizationConfig) error {
	if config.DrawStyle != "" && config.DrawStyle != "Classic" && config.DrawStyle != "Curved" && config.DrawStyle != "Linear" {
		return errors.New("invalid style option")
	}
	if config.ColorStyle != "" && config.ColorStyle != "Classic" && config.ColorStyle != "Complex" {
		return errors.New("invalid color option")
	}
	return nil
}

func GenerateSpellHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	spellName := r.URL.Query().Get("spellName")
	if spellName == "" {
		http.Error(w, "Missing spellName parameter", http.StatusBadRequest)
		return
	}
	spell := GetSpellByName(spellName)
	if spell == nil {
		http.Error(w, "Spell not found", http.StatusNotFound)
		return
	}
	var config shared.VisualizationConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, "Invalid config: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := validateConfig(&config); err != nil {
		http.Error(w, "Invalid config: "+err.Error(), http.StatusBadRequest)
		return
	}
	svgContent, err := visualizer.VisualizeSpell(spell, config)
	if err != nil {
		http.Error(w, "Error visualizing Spell: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(svgContent))
}
