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
