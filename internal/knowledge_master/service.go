package knowledge_master

import (
	"github.com/istvzsig/knowledge-master/db"
	"github.com/istvzsig/knowledge-master/internal/types"
)

func FetchFAQs() ([]types.FAQ, error) {
	return db.GetFAQs()
}

func CreateFAQ(faq types.FAQ) (string, error) {
	return db.CreateFAQ(faq)
}

func DeleteAllFAQs() ([]types.FAQ, error) {
	return db.DeleteAllFAQs()
}

func DeleteFAQByID(id string) error {
	return db.DeleteFAQByID(id)
}
