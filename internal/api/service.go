package knowledge_master

import (
	"github.com/istvzsig/knowledge-master/db"
	"github.com/istvzsig/knowledge-master/internal/types"
)

func FetchFAQs() (any, error) {
	return db.GetFAQs()
}

func CreateFAQ(faq types.FAQ) (string, error) {
	return db.CreateFAQ(faq)
}

func DeleteAllFAQs() (any, error) {
	return db.DeleteAllFAQs()
}

func DeleteFAQByID(id string) error {
	return db.DeleteFAQByID(id)
}
