package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	F "github.com/istvzsig/knowledge-master/internal/types"
	T "github.com/istvzsig/knowledge-master/pkg/types"
)

var mu sync.Mutex
var faqCollection = T.NewCollection[F.FAQ]()

func GetFAQs() (*T.Collection[F.FAQ], error) {
	ctx, cancel := getContextWithTimeout(10000)
	defer cancel()

	respChan := make(chan F.GetFAQsResponse)

	go func() {
		ref := FirestoreClient.NewRef("faqs")

		if err := ref.Get(ctx, &faqCollection.Items); err != nil {
			respChan <- F.GetFAQsResponse{
				FAQs: nil,
				Err:  fmt.Errorf("failed to get FAQs: %w", err),
			}
			return
		}

		mu.Lock()
		for key, faq := range faqCollection.Items {
			fmt.Println(faq)
			faqCollection.Items[key] = faq

		}
		mu.Unlock()

		respChan <- F.GetFAQsResponse{
			FAQs: faqCollection,
			Err:  nil,
		}
	}()

	select {
	case res := <-respChan:
		if res.Err != nil {
			log.Printf("Error fetching FAQs: %v", res.Err)
		}
		return res.FAQs, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("Getting FAQs timed out.")
	}
}

func CreateFAQ(faq F.FAQ) (string, error) {
	ctx, cancel := getContextWithTimeout(2000)
	defer cancel()

	respChan := make(chan F.CreateFAQResponse)

	go func() {
		ref := FirestoreClient.NewRef("faqs")

		faq.CreatedAt = time.Now().Unix()
		faq.ID = uuid.New().String()

		newRef, err := ref.Push(ctx, faq)
		if err != nil {
			respChan <- F.CreateFAQResponse{
				Key: "",
				Err: fmt.Errorf("failed to create FAQ: %w", err)}
			return
		}

		respChan <- F.CreateFAQResponse{
			Key: newRef.Key,
			Err: nil,
		}
	}()

	select {
	case res := <-respChan:
		if res.Err != nil {
			log.Printf("Error pushing FAQ to Firestore: %v", res.Err)
		}
		return res.Key, res.Err
	case <-ctx.Done():
		return "", fmt.Errorf("create FAQ operation timed out")
	}
}

func DeleteAllFAQs() (any, error) {
	ctx, cancel := getContextWithTimeout(2000)
	defer cancel()

	respChan := make(chan F.DeleteAllFAQsResponse)

	go func() {
		ref := FirestoreClient.NewRef("faqs")
		if err := ref.Set(ctx, nil); err != nil {
			respChan <- F.DeleteAllFAQsResponse{
				FAQs: nil,
				Err:  fmt.Errorf("failed to get FAQs: %w", err),
			}
			return
		}
	}()
	select {
	case resp := <-respChan:
		return resp.FAQs, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("delete all FAQs operation timed out")
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
