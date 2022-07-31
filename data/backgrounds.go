package data

import (
	"encoding/xml"
	"sort"
)

type Backgrounds struct {
	XMLName    xml.Name `xml:"compendium"`
	Text       string   `xml:",chardata"`
	Version    string   `xml:"version,attr"`
	AutoIndent string   `xml:"auto_indent,attr"`
	Background []struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name"`
		Proficiency string `xml:"proficiency"`
		Trait       []struct {
			Chardata string `xml:",chardata"`
			Name     string `xml:"name"`
			Text     string `xml:"text"`
		} `xml:"trait"`
	} `xml:"background"`
}

// Retrieve races from all data sources
func NewBackgrounds() ([]Backgrounds, error) {
	backgrounds, err := getData(Backgrounds{}, "backgrounds-")
	if err != nil {
		return nil, err
	}

	return backgrounds, nil
}

// Return a []string containing all background names sorted & de-duped
func ListBackgroundNames() ([]string, error) {
	b, err := NewBackgrounds()
	if err != nil || len(b) < 1 {
		return nil, err
	}

	var list []string
	for _, backgrounds := range b {
		for _, background := range backgrounds.Background {
			list = append(list, background.Name)
		}
	}

	list = unique(list)
	sort.Strings(list)
	return list, nil
}
