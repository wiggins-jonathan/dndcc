package data

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"sort"
	"sync"
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

// Concurrently retrieve races from data sources & unmarshal
func GetRaces() ([]Races, error) {
	filter := "races-"
	raceFiles, err := GetURLs(filter)
	if err != nil {
		return nil, err
	}

	var races []Races
	var wg sync.WaitGroup
	var mt sync.Mutex
	for _, raceFile := range raceFiles {
		wg.Add(1)
		go func(raceFile string) ([]Races, error) {
			defer wg.Done()
			// Read data source
			resp, err := http.Get(BaseURL + raceFile)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			// Convert to bytes
			mt.Lock()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			// Unmarshal bytes to struct
			var r Races
			if err := xml.Unmarshal(body, &r); err != nil {
				return nil, err
			}

			races = append(races, r)
			mt.Unlock()

			return races, nil
		}(raceFile)
	}
	wg.Wait()

	return races, nil
}

// Return a []string of all race names
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
