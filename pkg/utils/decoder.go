package utils

import (
	"encoding/json"
	"fmt"

	"github.com/ceci/go-challenge/pkg/models"
)

// DecodeAsset decodes JSON into the correct asset type
func DecodeAsset(data []byte) (models.Asset, error) {
	var base struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(data, &base); err != nil {
		return nil, err
	}

	switch base.Type {
	case models.AssetTypeChart:
		var chart models.Chart
		if err := json.Unmarshal(data, &chart); err != nil {
			return nil, err
		}
		return &chart, nil
	case models.AssetTypeInsight:
		var insight models.Insight
		if err := json.Unmarshal(data, &insight); err != nil {
			return nil, err
		}
		return &insight, nil
	case models.AssetTypeAudience:
		var audience models.Audience
		if err := json.Unmarshal(data, &audience); err != nil {
			return nil, err
		}
		return &audience, nil
	default:
		return nil, fmt.Errorf("invalid asset type")
	}
}
