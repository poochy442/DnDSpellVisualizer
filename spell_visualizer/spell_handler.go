package spell_visualizer

import (
	"net/http"
)

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

	svgContent := VisualizeSpell(spell)

	w.Header().Set("Content-Type", "image/svg+xml")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(svgContent))
}
