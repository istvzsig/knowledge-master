package faq

import (
	"github.com/istvzsig/knowledge-master/db"
	"github.com/istvzsig/knowledge-master/internal/types"
)

type Service interface {
	FetchFAQs() ([]types.FAQ, error)
	CreateFAQ(faq types.FAQ) (string, error)
	DeleteAllFAQs() error
	DeleteFAQByID(id string) error
}

type faqService struct {
	km *types.KnowledgeMaster
}

func NewFAQService(km *types.KnowledgeMaster) *faqService {
	return &faqService{km: km}
}

// FetchFAQs fetches all FAQs from the database.
func (s *faqService) FetchFAQs() ([]types.FAQ, error) {
	return db.GetFAQs()
}

// CreateFAQ creates a new FAQ in the database.
func (s *faqService) CreateFAQ(faq types.FAQ) (string, error) {
	return db.CreateFAQ(faq)
}

// DeleteAllFAQs deletes all FAQs from the database.
func (s *faqService) DeleteAllFAQs() error {
	return db.DeleteAllFAQs()
}

// DeleteFAQByID deletes a single FAQ by ID from the database.
func (s *faqService) DeleteFAQByID(id string) error {
	return db.DeleteFAQByID(id)
}
