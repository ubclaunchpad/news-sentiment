package db

// global variable to mongo connection

// make connection with mongo
func initSession() error {
	return nil
}

// insert User into mongo
func insertUser(new User) error {
	return nil
}

// insert NewsPiece associated with User into mongo
func insertArticle(article Article) error {
	return nil
}

// fetch User from mongo
func fetchUser(findUser User) (User, error) {
	return User{}, nil
}

// fetch NewsPiece associated with User from mongo
func fetchArticle(findNews Article) (Article, error) {
	return Article{}, nil
}

// Insert a Vote
func insertVote(vote Vote) error {
	return nil
}

// Get votes for article
func fetchVotes(article Article) ([]Vote, error) {
	return nil, nil
}
