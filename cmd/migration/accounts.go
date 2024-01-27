package main

import "github.com/eduardoths/finance-api/modules/ledger/structs"

var accounts = []structs.Account{
	{
		Model: structs.Model{ID: 1},
		Name:  "World",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 2},
		Name:  "Wise-BRL",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 3},
		Name:  "Wise-EUR",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 4},
		Name:  "Wise-CZK",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 5},
		Name:  "Wise-HUF",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 6},
		Name:  "Revolut-EUR",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 7},
		Name:  "Revolut-HUF",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 8},
		Name:  "XP-Normal",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 9},
		Name:  "XP-Investments",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 10},
		Name:  "Nomad-Normal",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 11},
		Name:  "Nomad-Investments",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 12},
		Name:  "Santander-PT",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 13},
		Name:  "BTG",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 14},
		Name:  "Nubank",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 15},
		Name:  "Clear",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 16},
		Name:  "Órama",
		Type:  structs.DEBIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 17},
		Name:  "Income",
		Type:  structs.CREDIT_NORMAL,
	},
	{
		Model: structs.Model{ID: 18},
		Name:  "Expenses",
		Type:  structs.DEBIT_NORMAL,
	},
}
