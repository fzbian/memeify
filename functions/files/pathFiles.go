package files

type Paths struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Path    string  `json:"path"`
	Example Example `json:"example"`
}

type Example struct {
	Text []string `json:"text"`
	URL  string   `json:"url"`
}

var (
	urlHOST  = "localhost"
	portHOST = ":1323"
)

var GlobalPaths = []Paths{
	{
		ID:   0,
		Name: "gopher",
		Path: "memes/gopher.jpg",
		Example: Example{
			Text: []string{"go", "for%20life"},
			URL:  urlHOST + portHOST + "/api/gopher/go/for%20life",
		},
	},
	{
		ID:   1,
		Name: "trump",
		Path: "memes/trump.jpg",
		Example: Example{
			Text: []string{"top%20text", "bottom%20text"},
			URL:  urlHOST + portHOST + "/api/trump/top%20text/bottom%20text",
		},
	},
	{
		ID:   2,
		Name: "mrbean",
		Path: "memes/mr-bean.jpg",
		Example: Example{
			Text: []string{"top%20text", "bottom%20text"},
			URL:  urlHOST + portHOST + "/api/mrbean/top%20text/bottom%20text",
		},
	},
}
