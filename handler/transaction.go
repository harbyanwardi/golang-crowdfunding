package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentuser").(user.User)
	input.User = currentUser
	trans, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse("Failed to get campaign transactions", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign Transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(trans))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentuser").(user.User)
	userID := currentUser.ID

	trans, err := h.service.GetTransactionByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("User Transaction", http.StatusOK, "success", transaction.FormatUserTransactions(trans))
	c.JSON(http.StatusOK, response)
}
