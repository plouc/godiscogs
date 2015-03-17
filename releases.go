package godiscogs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Release struct {
	Title             string `json:"title"`
	Status            string `json:"status"`
	ReleasedFormatted string `json:"released_formatted`
	EstimatedWeight   int    `json:"estimated_weight"`
	MasterUrl         string `json:"master_url"`
	Released          string `json:"released"`
	DateAdded         string `json:"date_added"`
	Country           string `json:"country"`
	Notes             string `json:"notes"`
}

func (d *Discogs) Release(id int) {
	var req *http.Request
	var err error

	url := fmt.Sprintf("https://api.discogs.com/releases/%d", id)

	req, err = http.NewRequest("GET", url, nil)

	resp, err := d.Client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	if resp.StatusCode >= 400 {
		err = errors.New("request failed: " + resp.Status)
		return
	}

	fmt.Printf("%##v", string(contents))
}
