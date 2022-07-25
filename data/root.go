// Explain that first we have to hit the /Collections dir in FightClub5eXML to
// get the location of the data we want, CoreOnly. Then explain that file points
// to all the other data we want
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
			coreURL.Href = strings.TrimLeft(coreURL.Href, "..") // Remove dots
			coreURLs = append(coreURLs, coreURL.Href)
		}
	}

	return coreURLs, nil
}

// Given url to a data source, unmarshal the data source to a generic, empty
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

// Given a generic data type & filter, unmarshal the generic data type into a
// slice of that type containing our data
func getData[T any](data T, filter string) ([]T, error) {
	// Get list of URLs filtered (by classes, races, spells, etc)
	files, err := getURLs(filter)
	if err != nil {
		return nil, err
	}

	// Unmarshal all of the URLs pointing to xml & put in slice of generic type
	var d []T
	for _, file := range files {
		func(file string) ([]T, error) {
			datum, err := genericUnmarshal(&data, BaseURL+file)
			if err != nil {
				return nil, err
			}

			d = append(d, *datum)
			return d, nil
		}(file)
	}

	return d, nil
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
