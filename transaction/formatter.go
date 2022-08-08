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
