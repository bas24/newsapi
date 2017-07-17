package newsapi

const (
	URL = "https://newsapi.org/v1/articles?"
)

type Article struct {
	Autor       string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}

type Response struct {
	Status   string    `json:"status"`
	Source   string    `json:"source"`
	SortBy   string    `json:"sortBy"`
	Articles []Article `json:"articles"`
}
