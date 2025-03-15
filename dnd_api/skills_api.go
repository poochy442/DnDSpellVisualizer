package dnd_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://www.dnd5eapi.co"

type APIResponse struct {
	Count   int `json:"count"`
	Results []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Level int    `json:"level"`
		URL   string `json:"url"`
	} `json:"results"`
}

type APISpell struct {
	Index         string   `json:"index"`
	Name          string   `json:"name"`
	Description   []string `json:"desc"`
	HigherLevel   []string `json:"higher_level"`
	Range         string   `json:"range"`
	Components    []string `json:"components"`
	Material      string   `json:"material"`
	Ritual        bool     `json:"ritual"`
	Duration      string   `json:"duration"`
	Concentration bool     `json:"concentration"`
	CastingTime   string   `json:"casting_time"`
	Level         int      `json:"level"`
	School        struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"school"`
	Damage struct {
		DamageType struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"damage_type"`
		DamageAtSlotLevel map[string]string `json:"damage_at_slot_level"`
	} `json:"damage"`
	DC struct {
		DCType struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"dc_type"`
		DCSuccess string `json:"dc_success"`
	} `json:"dc"`
	AreaOfEffect struct {
		Type string `json:"type"`
		Size int    `json:"size"`
	} `json:"area_of_effect"`
	Classes []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"classes"`
	Subclasses []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"subclasses"`
	URL       string `json:"url"`
	UpdatedAt string `json:"updated_at"`
}

func GetSpells() ([]*APISpell, error) {
	apiResponse, err := getAllSpellsFromAPI()
	if err != nil {
		return nil, err
	}

	var spells []*APISpell
	for _, result := range apiResponse.Results {
		spell, err := getSpellFromAPI(result.URL)
		if err != nil {
			return nil, fmt.Errorf("failed to get spell details for %s: %v", result.Name, err)
		}
		spells = append(spells, spell)
	}

	return spells, nil
}

func getAllSpellsFromAPI() (*APIResponse, error) {
	resp, err := http.Get(baseURL + "/api/spells")
	if err != nil {
		return nil, fmt.Errorf("failed to get spells: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get spells: status code %d", resp.StatusCode)
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %v", err)
	}

	return &apiResponse, nil
}

func getSpellFromAPI(url string) (*APISpell, error) {
	resp, err := http.Get(baseURL + url)
	if err != nil {
		return nil, fmt.Errorf("failed to get spell details: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get spell details: status code %d", resp.StatusCode)
	}

	var apiSpell APISpell
	if err := json.NewDecoder(resp.Body).Decode(&apiSpell); err != nil {
		return nil, fmt.Errorf("failed to decode spell details: %v", err)
	}

	return &apiSpell, nil
}
