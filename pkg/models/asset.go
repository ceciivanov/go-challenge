package models

// Define constants for Asset Types
const (
	AssetTypeChart    string = "chart"
	AssetTypeInsight  string = "insight"
	AssetTypeAudience string = "audience"
)

// Asset interface to be implemented by all asset types
type Asset interface {
	GetID() string
	GetType() string
	GetDescription() string
}
