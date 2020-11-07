package db

// User represents the data that a user would have
type User struct {
	ID         string
	Name       string
	Email      string
	Preference string
	Votes      []Vote
}

// NewsPiece represents a News Article
type Article struct {
	Source string
	Title  string
	URL    string
	Votes  []Vote
}

// Vote represents a vote
type Vote struct {
	UserId    string
	ArticleId string
	VoteValue int32
}
