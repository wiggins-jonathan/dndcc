/*
The source of all our data is located at https://github.com/kinkofer/FightClub5eXML
This repo has an almost exhaustive source in xml format, divided by rule book.
It also divides the data by its canonical value, i.e. - Core Rule books,
homebrew, Unearthed Arcana, etc. To get just the Core only data, parse the
`CoreOnly.xml` for a list of the location of only `core` data.
*/
package data

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

const BaseUrl = `https://raw.githubusercontent.com/kinkofer/FightClub5eXML/master`

type CoreOnly struct {
	XMLName xml.Name `xml:"collection"`
	Text    string   `xml:",chardata"`
	Doc     []struct {
		Text string `xml:",chardata"`
		Href string `xml:"href,attr"`
	} `xml:"doc"`
}

// Query list of all `CoreOnly` data sources at CoreOnly.xml, unmarshal, &
// return in []string in /FightClub5eXML/Sources/<source book>/<data> format
func getUrls() ([]string, error) {
	c, err := genericUnmarshal(&CoreOnly{}, BaseUrl+`/Collections/CoreOnly.xml`)
	if err != nil {
		return nil, err
	}

	coreData := make([]string, len(c.Doc))
	for i, coreHref := range c.Doc {
		coreHref.Href = strings.TrimLeft(coreHref.Href, "..") // Remove dots
		coreData[i] = BaseUrl + coreHref.Href
	}

	return coreData, nil
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
	if resp.StatusCode > 299 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf(string(body))
	}

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
func getData[T any](data T, f string) ([]T, error) {
	// Get list of urls & filter (by classes, races, spells, etc)
	urls, err := getUrls()
	if err != nil {
		return nil, err
	}
	urls = filter(urls, f)

	// Unmarshal all of the URLs pointing to xml & put in slice of generic type
	client := &http.Client{Timeout: 10 * time.Second}
	var wg sync.WaitGroup
	var mt sync.Mutex
	d := make([]T, len(urls))
	for i, url := range urls {
		wg.Add(1)
		go func(j int, url string) ([]T, error) {
			defer wg.Done()
			// Read data source
			resp, err := client.Get(url)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			if resp.StatusCode > 299 {
				body, _ := ioutil.ReadAll(resp.Body)
				return nil, fmt.Errorf(string(body))
			}

			// Convert to bytes
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			mt.Lock()
			// Unmarshal bytes to struct
			if err := xml.Unmarshal(body, &data); err != nil {
				return nil, err
			}

			d[j] = data
			mt.Unlock()

			return d, nil
		}(i, url)
	}
	wg.Wait()

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

// Given a []string from getUrls, filter & return with BaseUrl attached
func filter(urls []string, filter string) []string {
	var filteredUrls []string
	for _, url := range urls {
		if ok := strings.Contains(url, filter); ok {
			filteredUrls = append(filteredUrls, url)
		}
	}

	return filteredUrls
}
