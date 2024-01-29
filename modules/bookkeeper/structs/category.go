package structs

type IncomeCategory string

const (
	WAGE          IncomeCategory = "WAGE"
	BANK_TRANSFER IncomeCategory = "BANK_TRANSFER"
	DONATION      IncomeCategory = "DONATION"
	BETS          IncomeCategory = "BETS"
)

type ExpenseCategory string

const (
	PUBLIC_TRANSPORT ExpenseCategory = "PUBLIC_TRANSPORT"
	UBER             ExpenseCategory = "UBER"
	GYM              ExpenseCategory = "GYM"
	EATING_OUT       ExpenseCategory = "EATING_OUT"
	GROCERIES        ExpenseCategory = "GROCERIES"
	ELECTRONICS      ExpenseCategory = "ELECTRONICS"
	PHARMACY         ExpenseCategory = "PHARMACY"
	SPORTS           ExpenseCategory = "SPORTS"
	HOME             ExpenseCategory = "HOME"
	CELLPHONE        ExpenseCategory = "CELLPHONE"
	SUPLEMENTS       ExpenseCategory = "SUPLEMENTS"
	RENT             ExpenseCategory = "RENT"
	LAUNDRY          ExpenseCategory = "LAUNDRY"
	SNACKS           ExpenseCategory = "SNACKS"
	TOURISM          ExpenseCategory = "TOURISM"
	CLOTHES          ExpenseCategory = "CLOTHES"
	BUREAUCRACY      ExpenseCategory = "BUREAUCRACY"
	SOUVENIRS        ExpenseCategory = "SOUVENIRS"
)

type SubCategory string

const (
	COOKIE               SubCategory = "COOKIE"
	YOGURT               SubCategory = "YOGURT"
	MEAT                 SubCategory = "MEAT"
	DAIRY                SubCategory = "DAIRY"
	VEGETABLE            SubCategory = "VEGETABLE"
	BREAD                SubCategory = "BREAD"
	EGG                  SubCategory = "EGG"
	POTATO               SubCategory = "POTATO"
	CANDY                SubCategory = "CANDY"
	FRUITS               SubCategory = "FRUITS"
	SODA                 SubCategory = "SODA"
	BAG                  SubCategory = "BAG"
	SAUCE                SubCategory = "SAUCE"
	SALAD                SubCategory = "SALAD"
	DISCOUNTS            SubCategory = "DISCOUNTS"
	SNACK                SubCategory = "SNACK"
	CEREAL               SubCategory = "CEREAL"
	COOKIE_ACCOMPANIMENT SubCategory = "COOKIE_ACCOMPANIMENT"
	LAUNDRY_SC           SubCategory = "LAUNDRY_SC"
	CLEANING             SubCategory = "CLEANING"
)
