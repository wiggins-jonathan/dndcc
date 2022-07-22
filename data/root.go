package data

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

const BaseURL  = `https://raw.githubusercontent.com/kinkofer/FightClub5eXML/master`

type CoreOnly struct {
	XMLName xml.Name `xml:"collection"`
	Text    string   `xml:",chardata"`
	Doc     []struct {
		Text string `xml:",chardata"`
		Href string `xml:"href,attr"`
	} `xml:"doc"`
}

// Parse an xml formatted list pointing to urls containing data on all the
// 'core' 5e books
func GetCoreData() (*CoreOnly, error) {
	resp, err := http.Get(BaseURL + `/Collections/CoreOnly.xml`)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var c CoreOnly
	if err := xml.Unmarshal(body, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

// Filter out specific list of data urls based on filter
func GetURLs(filter string) ([]string, error) {
	c, err := GetCoreData()
	if err != nil {
		return nil, err
	}

	var coreURLs []string
	for _, coreURL := range c.Doc {
		if ok := strings.Contains(coreURL.Href, filter); ok {
			coreURL.Href = strings.TrimLeft(coreURL.Href, "..")
			coreURLs = append(coreURLs, coreURL.Href)
		}
	}

	return coreURLs, nil
}
