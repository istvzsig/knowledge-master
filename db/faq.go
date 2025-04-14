package db

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"maps"

	"github.com/google/uuid"
	F "github.com/istvzsig/knowledge-master/internal/types"
	T "github.com/istvzsig/knowledge-master/pkg/types"
)

const RESPONSE_TIMOUT = 10000
const BUFF_SIZE = 1000

var mu sync.Mutex
var faqCollection = T.NewCollection[F.FAQ]()

func GetFAQs() (*T.Collection[F.FAQ], error) {
	return responseWithTimeout(RESPONSE_TIMOUT, BUFF_SIZE, func(ctx context.Context) (*T.Collection[F.FAQ], error) {
		ref := FirestoreClient.NewRef("faqs")

		if err := ref.Get(ctx, &faqCollection.Items); err != nil {
			return nil, fmt.Errorf("failed to get FAQs: %w", err)
		}

		mu.Lock()
		maps.Copy(faqCollection.Items, faqCollection.Items)
		mu.Unlock()

		return faqCollection, nil
	})
}

func CreateFAQ(faq F.FAQ) (string, error) {
	return responseWithTimeout(RESPONSE_TIMOUT, BUFF_SIZE, func(ctx context.Context) (string, error) {
		ref := FirestoreClient.NewRef("faqs")

		faq.CreatedAt = time.Now().Unix()
		faq.ID = uuid.New().String()

		newRef, err := ref.Push(ctx, faq)
		if err != nil {
			return "", fmt.Errorf("failed to create FAQ: %w", err)
		}

		return newRef.Key, nil
	})
}

func DeleteAllFAQs() (any, error) {
	return responseWithTimeout(RESPONSE_TIMOUT, BUFF_SIZE, func(ctx context.Context) (any, error) {
		ref := FirestoreClient.NewRef("faqs")
		if err := ref.Set(ctx, nil); err != nil {
			return nil, fmt.Errorf("failed to delete all FAQs: %w", err)
		}
		return nil, nil
	})
}

func DeleteFAQByID(id string) error {
	_, err := responseWithTimeout(RESPONSE_TIMOUT, BUFF_SIZE, func(ctx context.Context) (struct{}, error) {
		fRef := FirestoreClient.NewRef("faqs")
		ref := fRef.Child(id)

		if err := ref.Delete(ctx); err != nil {
			return struct{}{}, fmt.Errorf("failed to delete FAQ with ID %s: %w", id, err)
		}

		return struct{}{}, nil
	})

	if err != nil {
		log.Printf("Error deleting FAQ with ID %s: %v", id, err)
	}

	return err
}

func responseWithTimeout[T any](timeout time.Duration, buf int, fn func(ctx context.Context) (T, error)) (T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*timeout)
	defer cancel()

	resultChan := make(chan struct {
		val T
		err error
	}, buf)

	go func() {
		val, err := fn(ctx)
		resultChan <- struct {
			val T
			err error
		}{val, err}
	}()

	select {
	case res := <-resultChan:
		return res.val, res.err
	case <-ctx.Done():
		var zero T
		return zero, fmt.Errorf("operation timed out after %dms", timeout)
	}
}
