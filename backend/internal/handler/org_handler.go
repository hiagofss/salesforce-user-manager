package handler

import (
	"net/http"
	"salesforce-user-manager/internal/domain"
	"salesforce-user-manager/internal/usecase"

	"github.com/gin-gonic/gin"
)

type OrgHandler struct {
	usecase *usecase.OrgUsecase
}

func NewOrgHandler(usecase *usecase.OrgUsecase) *OrgHandler {
	return &OrgHandler{usecase: usecase}
}

func (h *OrgHandler) GetOrg(c *gin.Context) {
	id := c.Param("id")

	org, err := h.usecase.GetOrg(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Org not found"})
		return
	}
	c.JSON(http.StatusOK, org)
}

func (h *OrgHandler) GetAllOrgs(c *gin.Context) {
	orgs, err := h.usecase.GetAllOrgs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orgs"})
		return
	}
	c.JSON(http.StatusOK, orgs)
}

func (h *OrgHandler) CreateOrg(c *gin.Context) {
	var org domain.Org
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdOrg, err := h.usecase.CreateOrg(&org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create org"})
		return
	}
	c.JSON(http.StatusCreated, createdOrg)
}

func (h *OrgHandler) UpdateOrg(c *gin.Context) {
	id := c.Param("id")
	var org domain.Org
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	org.Id = id // Ensure the ID from the URL is used
	if err := h.usecase.UpdateOrg(&org); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update org"})
		return
	}
	c.JSON(http.StatusOK, org)
}

func (h *OrgHandler) DeleteOrg(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteOrg(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete org"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *OrgHandler) SyncUsers(c *gin.Context) {
	id := c.Param("id")
	err := h.usecase.SyncUsersFromSalesforce(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Users synced successfully"})
}
