package structs

type Asset struct {
	Model
	Type AssetType `json:"type"`
	Code string    `json:"code"`
	Name string    `json:"name"`
}

type AssetType string

const (
	CURRENCY   AssetType = "CURRENCY"
	INVESTMENT AssetType = "INVESTMENT"
)
