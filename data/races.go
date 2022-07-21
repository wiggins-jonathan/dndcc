package data

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"sort"
	"sync"
)

// Data source
const raceBaseURL = "https://raw.githubusercontent.com/kinkofer/FightClub5eXML/master"

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

// Concurrently retrieve races from data source
func GetRaces() ([]string, error) {
	filter := "races-"
	raceFiles, err := GetURLs(filter)
	if err != nil {
		return nil, err
	}

	var races []string
	var wg sync.WaitGroup
	var mt sync.Mutex
	for _, raceFile := range raceFiles {
		wg.Add(1)
		go func(raceFile string) ([]string, error) {
			defer wg.Done()
			// Read data source
			resp, err := http.Get(raceBaseURL + raceFile)
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

			for _, race := range r.Race {
				races = append(races, race.Name)
			}
			mt.Unlock()

			return races, nil
		}(raceFile)
	}
	wg.Wait()

	sort.Strings(races)
	return races, nil
}
