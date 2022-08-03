package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	FindCampaign(UserID int) ([]Campaign, error)
	FindCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaign(UserID int) ([]Campaign, error) {
	if UserID != 0 {
		campaign, err := s.repository.FindByUserID(UserID)
		if err != nil {
			return campaign, err
		}
		return campaign, nil
	}
	campaign, err := s.repository.FindAll()
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) FindCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindCampaignByID(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID
	slugtext := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	//proses pembuatan slug menggunakan library slug dari github
	campaign.Slug = slug.Make(slugtext)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}
