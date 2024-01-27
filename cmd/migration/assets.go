package main

import "github.com/eduardoths/finance-api/modules/ledger/structs"

var assets = []structs.Asset{
	{
		Model:     structs.Model{ID: 1},
		AssetType: structs.CURRENCY,
		AssetCode: "BRL",
		AssetName: "Real",
	},
	{
		Model:     structs.Model{ID: 2},
		AssetType: structs.CURRENCY,
		AssetCode: "EUR",
		AssetName: "Euro",
	},
	{
		Model:     structs.Model{ID: 3},
		AssetType: structs.CURRENCY,
		AssetCode: "USD",
		AssetName: "Dollar",
	},
	{
		Model:     structs.Model{ID: 4},
		AssetType: structs.CURRENCY,
		AssetCode: "CZK",
		AssetName: "Czech Koruna",
	},
	{
		Model:     structs.Model{ID: 5},
		AssetType: structs.CURRENCY,
		AssetCode: "HUF",
		AssetName: "Hungarian Forint",
	},
}
