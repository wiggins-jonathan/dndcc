package data

import (
	"encoding/xml"
	"sort"
)

type Feats struct {
	XMLName    xml.Name `xml:"compendium"`
	Text       string   `xml:",chardata"`
	Version    string   `xml:"version,attr"`
	AutoIndent string   `xml:"auto_indent,attr"`
	Feat       []struct {
		Chardata     string `xml:",chardata"`
		Name         string `xml:"name"`
		Prerequisite string `xml:"prerequisite"`
		Text         string `xml:"text"`
		Modifier     struct {
			Text     string `xml:",chardata"`
			Category string `xml:"category,attr"`
		} `xml:"modifier"`
	} `xml:"feat"`
}

// Retrieve feats from all data sources
func NewFeats() ([]Feats, error) {
	feats, err := getData(Feats{}, "feats-")
	if err != nil {
		return nil, err
	}

	return feats, nil
}

// Return a []string containing all feat names sorted & de-duped
func ListFeatNames(f []Feats) []string {
	var list []string
	for _, feats := range f {
		for _, feat := range feats.Feat {
			list = append(list, feat.Name)
		}
	}

	list = unique(list)
	sort.Strings(list)
	return list
}
