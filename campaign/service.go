package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	FindCampaign(UserID int) ([]Campaign, error)
	FindCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	Update(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
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

func (s *service) Update(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindCampaignByID(inputID.ID)
	if err != nil {
		return campaign, err
	}
	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount
	// slugtext := fmt.Sprintf("%s %d", inputData.Name, inputData.User.ID)
	// //proses pembuatan slug menggunakan library slug dari github
	// campaign.Slug = slug.Make(slugtext)

	//jika user yg update bukan si pemilik campaign
	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("Not an owner of the campaign")
	}

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}
	return updatedCampaign, nil

}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkImagesNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}

	}
	campaign, err := s.repository.FindCampaignByID(input.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {

		return CampaignImage{}, errors.New("Not an owner of the campaign")
	}
	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID

	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	campImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return campImage, err
	}
	return campImage, nil
}
