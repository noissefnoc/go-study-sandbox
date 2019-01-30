package apiclient

import (
	"io/ioutil"
	"net/http"
)

type API struct {
	Client *http.Client
	baseURL string
}

func (api *API) DoStuff() ([]byte, error) {
	resp, err := api.Client.Get(api.baseURL + "/some/path")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, err
}
