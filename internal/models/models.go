package models

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Topic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Editorial struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Year      int        `json:"year"`
	Language  string     `json:"language"`
	ISBN      string     `json:"isbn"`
	Edition   string     `json:"edition"`
	CoverURL  string     `json:"cover_url"`
	Editorial *Editorial `json:"editorial,omitempty"`
	Pages     int        `json:"pages"`
	Location  string     `json:"location"`
	Authors   []Author   `json:"authors"`
	Topics    []Topic    `json:"topics"`
}
