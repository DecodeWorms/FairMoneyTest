package model

type Transactions struct {
	AccountID string  `json:"account_id"`
	Reference string  `json:"reference"`
	Amount    float64 `json:"amount"`
}

type TransactionRecords struct {
	ID        string  `json:"id" bson:"id"`
	AccountID string  `json:"account_id" bson:"account_id"`
	Reference string  `json:"reference" bson:"reference"`
	Amount    float64 `json:"amount" bson:"amount"`
	Balance   float64 `json:"balance" bson:"balance"`
}
