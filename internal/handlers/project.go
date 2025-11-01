package handlers

import (
	"net/http"

	"taskboardapi/internal/database"
	"taskboardapi/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateProjectInput struct {
	Title		string		`json:"title"  binding:"required"`
	Description	string		`json:"description"`
}

type UpdateProjectInput struct {
	Title		string		`json:"title"  binding:"required"`
	Description	string		`json:"description"`
}

func CreateProject(c *gin.Context) {
	var input CreateProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	project := models.Project{
		Title:       input.Title,
		Description: input.Description,
		UserID:      userID,
	}

	if err := database.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func GetProjects(c *gin.Context) {
	userID := c.GetUint("user_id")

	var projects []models.Project
	if err := database.DB.Where("user_id = ?", userID).Preload("Tasks").Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func GetProject(c *gin.Context) {
	projectID := c.Param("id")
	userID := c.GetUint("user_id")

	var project models.Project
	if err := database.DB.Where("id = ? AND user_id = ?", projectID, userID).Preload("Tasks").First(&project).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	c.JSON(http.StatusOK, project)
}

func UpdateProject(c *gin.Context) {
	projectID := c.Param("id")
	userID := c.GetUint("user_id")

	var project models.Project
	if err := database.DB.Where("id = ? AND user_id = ?", projectID, userID).First(&project).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	var input UpdateProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project.Title = input.Title
	project.Description = input.Description

	if err := database.DB.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}
	c.JSON(http.StatusOK, project)
}

func DeleteProject(c *gin.Context) {
	projectID := c.Param("id")
	userID := c.GetUint("user_id")

	var project models.Project
	if err := database.DB.Where("id = ? AND user_id = ?", projectID, userID).First(&project).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	if err := database.DB.Delete(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}