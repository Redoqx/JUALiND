package schema

import "database/sql"

type JSONResponse struct {
	Message string      `json:"msg"`
	Body    interface{} `json:"body,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type OrderStruct struct {
	BuyerName        string         `json:"buyer_name" csv:"buyer_name"`
	ProductName      string         `json:"product_name" csv:"product_name"`
	ProductPrice     int            `json:"product_price" csv:"product_price"`
	Amount           int            `json:"amount" csv:"amount"`
	Date             string         `json:"date" csv:"date"`
	ConfirmationLink sql.NullString `json:"confirmation_link" csv:"-"`
}
