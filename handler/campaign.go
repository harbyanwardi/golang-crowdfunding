package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	UserID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.FindCampaign(UserID)
	if err != nil {

		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List Campaign", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler) GetCampaignDetail(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	campaignDetail, err := h.service.FindCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Campaign Detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Create Campaign Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//currentuser adalah context yg telah kita set di middleware ketika berhasil validasi token
	currentUser := c.MustGet("currentuser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {

		response := helper.APIResponse("Create Campaign Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Create Campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
	return
}
