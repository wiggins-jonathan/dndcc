package data

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// Races in the player's hand book
const raceBaseURL = "https://raw.githubusercontent.com/kinkofer/FightClub5eXML/master/FightClub5eXML/Sources"

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

func GetRaces() ([]string, error) {
	raceFiles := []string{
		"/PlayersHandbook/races-phb.xml",
		//"/DungeonMastersGuide/races-dmg.xml",
		//"/PrincesOfTheApocalypse/races-eepc.xml",
		//"/MordenkainensTomeOfFoes/races-mtf.xml",
		//"/VolosGuideToMonsters/races-vgm.xml",
		//"/GuildmastersGuideToRavnica/races-ggr.xml",
		//"/EberronRisingFromTheLastWar/races-erlw.xml",
		//"/AcquisitionsIncorporated/races-ai.xml",
		//"/ExplorersGuideToWildemount/races-egw.xml",
		//"/TheTortlePackage/races-ttp.xml",
		//"/OneGrungAbove/races-oga.xml",
		//"/MythicOdysseysOfTheros/races-mot.xml",
		//"/TashasCauldronOfEverything/races-tce.xml",
		//"/LocathahRising/races-lr.xml",
		//"/AdventureWithMuk/races-awm.xml",
	}

	var races []string
	for _, raceFile := range raceFiles {
		// Read data source
		resp, err := http.Get(raceBaseURL + raceFile)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Convert to bytes
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
	}

	return races, nil
}
