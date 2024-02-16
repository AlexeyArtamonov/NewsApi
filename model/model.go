package model

//go:generate reform

// News represents a row in news table.
//
//reform:news
type News struct {
	ID         int64  `reform:"id,pk"`
	Title      string `reform:"title"`
	Content    string `reform:"content"`
	Categories []int64
}

//go:generate reform

// Newscategories represents a row in newscategories table.
//
//reform:newscategories
type Newscategories struct {
	Newsid     int64 `reform:"newsid"`
	Categoryid int64 `reform:"categoryid"`
}

type JSONNews struct {
	Id         int64   `json:"Id"`
	Title      string  `json:"Title"`
	Content    string  `json:"Content"`
	Categories []int64 `json:"Categories"`
}
