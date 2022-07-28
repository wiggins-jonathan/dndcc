package data

import (
	"encoding/xml"
	"sort"
)

type Classes struct {
	XMLName    xml.Name `xml:"compendium"`
	Text       string   `xml:",chardata"`
	Version    string   `xml:"version,attr"`
	AutoIndent string   `xml:"auto_indent,attr"`
	Class      struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name"`
		Hd          string `xml:"hd"`
		Proficiency string `xml:"proficiency"`
		NumSkills   string `xml:"numSkills"`
		Autolevel   []struct {
			Text             string `xml:",chardata"`
			Level            string `xml:"level,attr"`
			ScoreImprovement string `xml:"scoreImprovement,attr"`
			Feature          struct {
				Chardata string   `xml:",chardata"`
				Optional string   `xml:"optional,attr"`
				Name     string   `xml:"name"`
				Text     []string `xml:"text"`
			} `xml:"feature"`
		} `xml:"autolevel"`
		Armor   string `xml:"armor"`
		Weapons string `xml:"weapons"`
		Tools   string `xml:"tools"`
		Wealth  string `xml:"wealth"`
	} `xml:"class"`
}

// Retrieve classes from data sources
func NewClasses() ([]Classes, error) {
	classes, err := getData(Classes{}, "class-")
	if err != nil {
		return nil, err
	}

	return classes, nil
}

// Return a []string of all class names sorted & de-duped
func ListClassNames() ([]string, error) {
	c, err := NewClasses()
	if err != nil || len(c) < 1 {
		return nil, err
	}

	var list []string
	for _, classes := range c {
		list = append(list, classes.Class.Name)
	}

	list = unique(list)
	sort.Strings(list)
	return list, nil
}
