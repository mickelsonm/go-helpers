/**
	This is a helper that interacts with Dweet.io

	The limitation is that under the free version you're restricted to:
	- 500 dweets
	- Persist for only 24 hours
	- Alerts only work with locked dweet "things"

	See their website on more specifics of these limitations.
**/
package dweetio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	DWEET_SET_API    = "https://dweet.io/dweet/for/%s"
	DWEET_GET_API    = "https://dweet.io/get/dweets/for/%s"
	DWEET_LATEST_API = "https://dweet.io/get/latest/dweet/for/%s"
)

type Dweet struct {
	This    string `json:"this"`
	Because string `json:"because,omitempty"`
	By      string `json:"by"`
	The     string `json:"the"`
	With    []With `json:"with"`
}

type With struct {
	Thing   string                 `json:"thing"`
	Created time.Time              `json:"created"`
	Content map[string]interface{} `json:"content"`
}

//TODO: implement Add dweet logic

func GetDweets(thingName string, latest bool) (dweets Dweet, err error) {
	if thingName == "" {
		err = errors.New("Must provide a dweet name.")
		return
	}
	var buf []byte
	var resp *http.Response
	var endpoint string

	if latest { //this grabs a single dweet
		endpoint = fmt.Sprintf(DWEET_LATEST_API, thingName)
	} else { //this grabs all dweets under this thing's name
		endpoint = fmt.Sprintf(DWEET_GET_API, thingName)
	}

	if resp, err = http.Get(endpoint); err != nil {
		return
	}
	defer resp.Body.Close()

	if buf, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(buf, &dweets); err != nil {
		return
	}

	return
}
