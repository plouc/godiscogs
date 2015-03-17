package godiscogs

import (
	"net/http"
)

type Resource struct {
	Id          int    `json:"id"`
	ResourceUrl string `json:"resource_url`
}

type Label struct {
	Resource
	Name string `json:"name"`
}

type Artist struct {
	Resource
	Name string `json:"name"`
}

type Track struct {
	Title    string `json:"title"`
	Duration string `json:"duration`
	Position string `json:"position"`
}

type Discogs struct {
	BaseUrl string
	Client  *http.Client
}

func NewDiscogs(baseUrl string) *Discogs {

	client := &http.Client{}

	return &Discogs{
		BaseUrl: baseUrl,
		Client:  client,
	}
}
