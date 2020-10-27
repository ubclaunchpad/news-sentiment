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
func insertNewsPiece(new NewsPiece, associatedUser User) error {
	return nil
}

// fetch User from mongo
func fetchUser(findUser User) (error, User) {
	return nil
}

// fetch NewsPiece associated with User from mongo
func fetchNewsPiece(findNews NewsPiece, associatedUser User) (error, User, NewsPiece) {
	return nil
}
