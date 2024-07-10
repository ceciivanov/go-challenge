package models

// Insight represents an insight asset
type Insight struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Text        string `json:"text"`
}

// Implement the Asset interface for Insight
func (i Insight) GetID() int {
	return i.ID
}

func (i Insight) GetType() string {
	return i.Type
}

func (i Insight) GetDescription() string {
	return i.Description
}
