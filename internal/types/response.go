package types

import T "github.com/istvzsig/knowledge-master/pkg/types"

type CreateFAQResponse struct {
	Key string
	Err error
}

type GetFAQsResponse struct {
	FAQs *T.Collection[FAQ]
	Err  error
}

type DeleteAllFAQsResponse struct {
	FAQs *T.Collection[FAQ]
	Err  error
}
