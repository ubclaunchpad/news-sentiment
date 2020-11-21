package db

type MockDatabase struct{}

// CreateNewUser adds new user to mongo db
func (md *MockDatabase) CreateNewUser(email string, name string) (string, error) {
	return "", nil
}

func (md *MockDatabase) CreateNewArticle(url string, title string, source string) (string, error) {
	return "", nil
}

func (md *MockDatabase) FindAllArticles() ([]Article, error) {
	return nil, nil
}
