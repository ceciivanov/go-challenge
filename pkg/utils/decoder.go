package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ceciivanov/go-challenge/pkg/models"
)

// ValidateChart validates that all required fields for Chart are present
func ValidateChart(chart *models.Chart) error {
	if chart.Title == "" {
		return errors.New("missing or invalid 'title' field for Chart")
	}
	if chart.XAxesTitle == "" {
		return errors.New("missing or invalid 'xAxesTitle' field for Chart")
	}
	if chart.YAxesTitle == "" {
		return errors.New("missing or invalid 'yAxesTitle' field for Chart")
	}
	if len(chart.DataPoints) == 0 {
		return errors.New("missing or invalid 'dataPoints' field for Chart")
	}

	return nil
}

// ValidateInsight validates that all required fields for Insight are present
func ValidateInsight(insight *models.Insight) error {
	if insight.Text == "" {
		return errors.New("missing or invalid 'text' field for Insight")
	}

	return nil
}

// ValidateAudience validates that all required fields for Audience are present
func ValidateAudience(audience *models.Audience) error {
	if audience.Age == 0 {
		return errors.New("missing or invalid 'age' field for Audience")
	}
	if audience.AgeGroup == "" {
		return errors.New("missing or invalid 'ageGroup' field for Audience")
	}
	if audience.Gender == "" {
		return errors.New("missing or invalid 'gender' field for Audience")
	}
	if audience.BirthCountry == "" {
		return errors.New("missing or invalid 'birthCountry' field for Audience")
	}
	if audience.HoursSpentOnMedia == 0 {
		return errors.New("missing or invalid 'hoursSpentOnMedia' field for Audience")
	}
	if audience.NumberOfPurchases == 0 {
		return errors.New("missing or invalid 'numberOfPurchases' field for Audience")
	}

	return nil
}

// DecodeAsset decodes JSON into the correct asset type
func DecodeAsset(data []byte) (models.Asset, error) {
	// var base struct {
	// 	Type models.AssetType `json:"type"`
	// }

	// if err := json.Unmarshal(data, &base); err != nil {
	// 	return nil, err
	// }

	var baseAsset struct {
		ID          int              `json:"id"`
		Type        models.AssetType `json:"type"`
		Description string           `json:"description"`
	}

	if err := json.Unmarshal(data, &baseAsset); err != nil {
		return nil, err
	}

	// Validate base fields
	if baseAsset.ID == 0 {
		return nil, errors.New("missing or invalid 'id' field")
	}
	if baseAsset.Type == "" {
		return nil, errors.New("missing or invalid 'type' field")
	}
	if baseAsset.Description == "" {
		return nil, errors.New("missing or invalid 'description' field")
	}

	switch baseAsset.Type {
	case models.ChartType:
		var chart models.Chart
		if err := json.Unmarshal(data, &chart); err != nil {
			return nil, err
		}
		if err := ValidateChart(&chart); err != nil {
			return nil, err
		}
		return &chart, nil
	case models.InsightType:
		var insight models.Insight
		if err := json.Unmarshal(data, &insight); err != nil {
			return nil, err
		}
		if err := ValidateInsight(&insight); err != nil {
			return nil, err
		}
		return &insight, nil
	case models.AudienceType:
		var audience models.Audience
		if err := json.Unmarshal(data, &audience); err != nil {
			return nil, err
		}
		if err := ValidateAudience(&audience); err != nil {
			return nil, err
		}
		return &audience, nil
	default:
		return nil, fmt.Errorf("invalid asset type")
	}
}
