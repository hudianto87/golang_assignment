package entity

import "time"

type Transaction struct {
	ID             int       `json:"id"`
	WalletID       int       `json:"walletid" binding:"required" gorm:"column:walletid"` //ini gorm menentukan kolom name di DB nya, by default selalu snack_case
	TrxType        string    `json:"trxtype" binding:"required" gorm:"column:trxtype"`   //TOPUP, TRANSFER, PAYMENT
	WalletSourceID int       `json:"walletsourceid" gorm:"column:walletsourceid"`
	Description    string    `json:"description" binding:"required"`
	Amount         float32   `json:"amount"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type TransactionResponse struct {
	ID             int       `json:"id"`
	WalletID       int       `json:"walletid" gorm:"column:walletid"`
	WalletName     string    `json:"walletname" gorm:"column:walletname"`
	TrxType        string    `json:"trxtype" gorm:"column:trxtype"` // TOPUP, TRANSFER, PAYMENT
	WalletSourceID int       `json:"walletsourceid" gorm:"column:walletsourceid"`
	Amount         float32   `json:"amount"`
	UserName       string    `json:"username"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
