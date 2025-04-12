package types

type KnowledgeMaster struct {
	FAQs []FAQ
}

func NewKnowledgeMaster() *KnowledgeMaster {
	return &KnowledgeMaster{
		FAQs: nil,
	}
}
