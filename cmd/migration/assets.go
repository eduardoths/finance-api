package main

import "github.com/eduardoths/finance-api/modules/ledger/structs"

var assets = []structs.Asset{
	{
		Model:     structs.Model{ID: 1},
		Type: structs.CURRENCY,
		Code: "BRL",
		Name: "Real",
	},
	{
		Model:     structs.Model{ID: 2},
		Type: structs.CURRENCY,
		Code: "EUR",
		Name: "Euro",
	},
	{
		Model:     structs.Model{ID: 3},
		Type: structs.CURRENCY,
		Code: "USD",
		Name: "Dollar",
	},
	{
		Model:     structs.Model{ID: 4},
		Type: structs.CURRENCY,
		Code: "CZK",
		Name: "Czech Koruna",
	},
	{
		Model:     structs.Model{ID: 5},
		Type: structs.CURRENCY,
		Code: "HUF",
		Name: "Hungarian Forint",
	},
}
