package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{}
	formatter.ID = campaign.ID
	formatter.UserID = campaign.UserID
	formatter.Name = campaign.Name
	formatter.ShortDescription = campaign.ShortDescription
	formatter.GoalAmount = campaign.GoalAmount
	formatter.CurrentAmount = campaign.CurrentAmount
	formatter.Slug = campaign.Slug
	formatter.ImageURL = ""
	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	//cara lama jika data kosong

	// if len(campaigns) == 0 {
	// 	return []CampaignFormatter{}
	// }
	// var campaignsFormatters []CampaignFormatter

	//cara baru jika data kosong
	campaignsFormatters := []CampaignFormatter{}
	for _, c := range campaigns {
		campaignFormatter := FormatCampaign(c)
		campaignsFormatters = append(campaignsFormatters, campaignFormatter)
	}
	return campaignsFormatters
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	BackerCount      int                      `json:"backer_count"`
	CurrentAmount    int                      `json:"current_amount"`
	UserID           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}

	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.BackerCount = campaign.BackerCount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.ImageURL = ""
	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk)) //trimspace menghilangkan space di depan
	}

	campaignDetailFormatter.Perks = perks

	//mengisi objecj json User
	campaignUserFormatter := CampaignUserFormatter{}
	user := campaign.User
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName

	campaignDetailFormatter.User = campaignUserFormatter

	//mengisi object json Images
	images := []CampaignImageFormatter{}
	for _, img := range campaign.CampaignImages { //CampaignImages refer ke entity Campaign struct
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = img.FileName

		isPrimary := false
		if img.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImageFormatter.IsPrimary = isPrimary
		images = append(images, campaignImageFormatter)

	}
	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}
