package structs

type Asset struct {
	Model
	AssetType AssetType
	AssetCode string
	AssetName string
}

type AssetType string

const (
	CURRENCY   AssetType = "CURRENCY"
	INVESTMENT AssetType = "INVESTMENT"
)
