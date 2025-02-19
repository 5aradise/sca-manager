package breeds

import (
	"net/http"
	"time"

	"github.com/5aradise/sca-manager/pkg/types"
	"github.com/bytedance/sonic"
)

const ApiUrl = "https://api.thecatapi.com/v1/breeds"

type BreedsResponse []struct {
	Name string `json:"name"`
}

func listBreeds(apiKey string, reqTimeout time.Duration) (types.Set[string], error) {
	req, err := http.NewRequest(http.MethodGet, ApiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", apiKey)

	client := http.Client{Timeout: reqTimeout}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var breedsRes BreedsResponse
	err = sonic.ConfigDefault.NewDecoder(res.Body).Decode(&breedsRes)
	if err != nil {
		return nil, err
	}

	set := make(types.Set[string], len(breedsRes))
	for _, breed := range breedsRes {
		set.Store(breed.Name)
	}
	return set, nil
}
