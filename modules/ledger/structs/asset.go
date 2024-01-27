package structs

type Asset struct {
	Model
	Type AssetType
	Code string
	Name string
}

type AssetType string

const (
	CURRENCY   AssetType = "CURRENCY"
	INVESTMENT AssetType = "INVESTMENT"
)
