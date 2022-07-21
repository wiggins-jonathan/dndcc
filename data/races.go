package data

import (
    "net/http"
    "io/ioutil"
    "encoding/xml"
)

// Races in the player's hand book
const phbraces = "https://raw.githubusercontent.com/kinkofer/FightClub5eXML/master/FightClub5eXML/Sources/PlayersHandbook/races-phb.xml"

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
    // Read data source
    resp, err := http.Get(phbraces)
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

    var races []string
    for _, race := range r.Race {
        races = append(races, race.Name)
    }

    return races, nil
}
