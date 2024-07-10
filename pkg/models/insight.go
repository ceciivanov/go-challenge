package models

// Insight represents an insight asset
type Insight struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Text        string `json:"text"`
}

// Implement the Asset interface for Insight
func (i Insight) GetID() string {
	return i.ID
}

func (i Insight) GetType() string {
	return i.Type
}

func (i Insight) GetDescription() string {
	return i.Description
}
