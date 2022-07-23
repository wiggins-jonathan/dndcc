package data

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

const BaseURL = `https://raw.githubusercontent.com/kinkofer/FightClub5eXML/master`

type CoreOnly struct {
	XMLName xml.Name `xml:"collection"`
	Text    string   `xml:",chardata"`
	Doc     []struct {
		Text string `xml:",chardata"`
		Href string `xml:"href,attr"`
	} `xml:"doc"`
}

// Query list of data sources at CoreOnly.xml & filter, returning a list of URLs
// containing data
func getURLs(filter string) ([]string, error) {
	c, err := genericUnmarshal(&CoreOnly{}, BaseURL+`/Collections/CoreOnly.xml`)
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

// Given urls to a data source, unmarshal the data source to a generic, empty
// data structure & return it full of data
func genericUnmarshal[T any](data *T, endpoint string) (*T, error) {
	// Read data source
	resp, err := http.Get(endpoint)
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
	if err := xml.Unmarshal(body, data); err != nil {
		return nil, err
	}

	return data, nil
}

// Remove duplicates from a slice
func unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, item := range s {
		if _, ok := inResult[item]; !ok {
			inResult[item] = true
			result = append(result, item)
		}
	}
	return result
}
