package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/istvzsig/knowledge-master/internal/types"
)

var mu sync.Mutex
var faqsMap = make(map[string]types.FAQ)

func GetFAQs() ([]types.FAQ, error) {
	ctx, cancel := getContextWithTimeout(50000)
	defer cancel()

	respChan := make(chan struct {
		faqs []types.FAQ
		err  error
	})

	go func() {
		ref := FirestoreClient.NewRef("faqs")

		if err := ref.Get(ctx, &faqsMap); err != nil {
			respChan <- struct {
				faqs []types.FAQ
				err  error
			}{nil, fmt.Errorf("failed to get FAQs: %w", err)}
			return
		}

		mu.Lock()
		for key, faq := range faqsMap {
			faq.ID = key
			faqsMap[key] = faq
		}
		mu.Unlock()

		var faqList []types.FAQ
		for _, faq := range faqsMap {
			faqList = append(faqList, faq)
		}

		respChan <- struct {
			faqs []types.FAQ
			err  error
		}{faqList, nil}
	}()

	select {
	case res := <-respChan:
		if res.err != nil {
			log.Printf("Error fetching FAQs: %v", res.err)
		}
		if len(res.faqs) > 0 {
			return res.faqs, nil
		} else {
			return nil, fmt.Errorf("No FAQs found.")
		}
	case <-ctx.Done():
		return nil, fmt.Errorf("get FAQs operation timed out")
	}
}

func CreateFAQ(faq types.FAQ) (string, error) {
	ctx, cancel := getContextWithTimeout(2000)
	defer cancel()

	respChan := make(chan struct {
		key string
		err error
	})

	go func() {
		faq.CreatedAt = time.Now().Unix()
		ref := FirestoreClient.NewRef("faqs")
		newRef, err := ref.Push(ctx, faq)
		if err != nil {
			respChan <- struct {
				key string
				err error
			}{"", fmt.Errorf("failed to create FAQ: %w", err)}
			return
		}
		respChan <- struct {
			key string
			err error
		}{newRef.Key, nil}
	}()

	select {
	case res := <-respChan:
		if res.err != nil {
			log.Printf("Error pushing FAQ to Firestore: %v", res.err)
		}
		return res.key, res.err
	case <-ctx.Done():
		return "", fmt.Errorf("create FAQ operation timed out")
	}
}

func DeleteAllFAQs() error {
	ctx, cancel := getContextWithTimeout(10)
	defer cancel()

	errCh := make(chan error)

	go func() {
		ref := FirestoreClient.NewRef("faqs")
		errCh <- ref.Set(ctx, nil)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		return fmt.Errorf("delete all FAQs operation timed out")
	}
}

func DeleteFAQByID(id string) error {
	ctx, cancel := getContextWithTimeout(500)
	defer cancel()

	errCh := make(chan error)

	go func() {
		fRef := FirestoreClient.NewRef("faqs")
		ref := fRef.Child(id)
		errCh <- ref.Delete(ctx)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			log.Printf("Error deleting FAQ with ID %s: %v", id, err)
			return fmt.Errorf("failed to delete FAQ with ID %s: %w", id, err)
		}
		return nil
	case <-ctx.Done():
		return fmt.Errorf("delete FAQ by ID operation timed out")
	}
}

func getContextWithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Millisecond*timeout)
}
