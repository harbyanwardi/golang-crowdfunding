package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	CampaignTransactionFormatters := []CampaignTransactionFormatter{}
	for _, c := range transactions {
		transactionFormatter := FormatCampaignTransaction(c)
		CampaignTransactionFormatters = append(CampaignTransactionFormatters, transactionFormatter)
	}
	return CampaignTransactionFormatters
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}
	formatter.Campaign = campaignFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	UserTransactionFormatters := []UserTransactionFormatter{}
	for _, c := range transactions {
		transactionFormatter := FormatUserTransaction(c)
		UserTransactionFormatters = append(UserTransactionFormatters, transactionFormatter)
	}
	return UserTransactionFormatters
}

type PaymentTransactionFormatter struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	PaymentURL string    `json:"payment_url"`
	Amount     int       `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatPaymentTransaction(transaction Transaction) PaymentTransactionFormatter {
	formatter := PaymentTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.UserID = transaction.UserID
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentURL = transaction.PaymentURL
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}
