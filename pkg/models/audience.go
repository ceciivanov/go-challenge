package models

// Audience represents an audience asset
type Audience struct {
	ID                string `json:"id"`
	Type              string `json:"type"`
	Description       string `json:"description"`
	Age               int    `json:"age"`
	AgeGroup          string `json:"ageGroup"`
	Gender            string `json:"gender"`
	BirthCountry      string `json:"birthCountry"`
	HoursSpentOnMedia int    `json:"hoursSpentOnMedia"`
	NumberOfPurchases int    `json:"numberOfPurchases"`
}

// Implement the Asset interface for Audience
func (a Audience) GetID() string {
	return a.ID
}

func (a Audience) GetType() string {
	return a.Type
}

func (a Audience) GetDescription() string {
	return a.Description
}
