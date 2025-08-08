package server

import (
	"DnDSpellVisualizer/spell_visualizer/shared"
	"DnDSpellVisualizer/spell_visualizer/visualizer"
	"encoding/json"
	"fmt"
	"net/http"
)

func validateConfig(config *shared.VisualizationConfig) error {
	if config.DrawStyle != "" && config.DrawStyle != shared.StyleClassic && config.DrawStyle != shared.StyleCurved && config.DrawStyle != shared.StyleLinear {
		return fmt.Errorf("invalid draw style: %s", config.DrawStyle)
	}
	if config.ColorStyle != "" && config.ColorStyle != shared.ColorClassic && config.ColorStyle != shared.ColorComplex {
		return fmt.Errorf("invalid color style: %s", config.ColorStyle)
	}
	return nil
}

func setDefaultConfig(config *shared.VisualizationConfig) {
	if config.DrawStyle == "" {
		config.DrawStyle = shared.StyleClassic
	}
	if config.ColorStyle == "" {
		config.ColorStyle = shared.ColorClassic
	}
}

func DrawSpellHandler(spells map[string]shared.Spell) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		spellName := r.URL.Query().Get("spellName")
		if spellName == "" {
			http.Error(w, "Missing spellName parameter", http.StatusBadRequest)
			return
		}
		spell, exists := spells[spellName]
		if !exists {
			http.Error(w, "Spell not found", http.StatusNotFound)
			return
		}
		var config shared.VisualizationConfig
		if r.Body == nil || r.ContentLength == 0 {
			config = shared.VisualizationConfig{}
			setDefaultConfig(&config)
		} else {
			if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
				http.Error(w, "Invalid config: "+err.Error(), http.StatusBadRequest)
				return
			}
			setDefaultConfig(&config)
		}
		if err := validateConfig(&config); err != nil {
			http.Error(w, "Invalid config: "+err.Error(), http.StatusBadRequest)
			return
		}
		svgContent, err := visualizer.VisualizeSpell(&spell, config)
		if err != nil {
			http.Error(w, "Error visualizing Spell: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(svgContent))
	}
}
