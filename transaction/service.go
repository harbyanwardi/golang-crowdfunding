package transaction

import (
	"bwastartup/campaign"
	"errors"
	"fmt"
	"time"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionByUserID(UserID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionsInput) (Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {

	//check campaign user id
	campaign, err := s.campaignRepository.FindCampaignByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}
	transaction, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) GetTransactionByUserID(UserID int) ([]Transaction, error) {
	transaction, err := s.repository.GetByUserID(UserID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) CreateTransaction(input CreateTransactionsInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.Status = "pending"
	transaction.UserID = input.User.ID

	currentTime := time.Now()
	date := currentTime.Format("01022006150405")
	code := fmt.Sprintf("ORDER%d%s", input.User.ID, date)
	transaction.Code = code

	newTrans, err := s.repository.Save(transaction)
	if err != nil {
		return newTrans, err
	}
	return newTrans, nil
}
