package types

type FAQ struct {
	ID        string `json:"id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	CreatedAt int64  `json:"createdAt"`
}

type FAQService interface {
	FetchFAQs() ([]FAQ, error)
	CreateFAQ(faq FAQ) (string, error)
	DeleteAllFAQs() ([]FAQ, error)
	DeleteFAQByID(id string) error
}
