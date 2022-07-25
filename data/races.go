package data

import (
	"encoding/xml"
	"sort"
)

type Races struct {
	XMLName    xml.Name `xml:"compendium"`
	Text       string   `xml:",chardata"`
	Version    string   `xml:"version,attr"`
	AutoIndent string   `xml:"auto_indent,attr"`
	Race       []struct {
		Text         string `xml:",chardata"`
		Name         string `xml:"name"`
		Size         string `xml:"size"`
		Speed        string `xml:"speed"`
		Ability      string `xml:"ability"`
		Proficiency  string `xml:"proficiency"`
		SpellAbility string `xml:"spellAbility"`
		Trait        []struct {
			Chardata string `xml:",chardata"`
			Name     string `xml:"name"`
			Text     string `xml:"text"`
		} `xml:"trait"`
	} `xml:"race"`
}

// Retrieve races from all data sources
func GetRaces() ([]Races, error) {
	races, err := getData(Races{}, "races-")
	if err != nil {
		return nil, err
	}

	return races, nil
}

// Return a []string containing all race names sorted & de-deuped
func ListRaceNames() ([]string, error) {
	r, err := GetRaces()
	if err != nil || len(r) < 1 {
		return nil, err
	}

	var list []string
	for _, races := range r {
		for _, race := range races.Race {
			list = append(list, race.Name)
		}
	}

	list = unique(list)
	sort.Strings(list)
	return list, nil
}
