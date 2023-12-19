package controllers

import (
	"go-todo/config"
	"go-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	var tasks []models.Task
	if err := models.GetAllTasks(&tasks); err != nil {
		config.WarningLogger.Printf("Failed to get tasks: %v\n", err)
		c.JSON(http.StatusNotFound, "Tasks not found")
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

func CreateNewTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		config.InfoLogger.Printf("Failed to bind request to JSON: %v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := models.CreateNewTask(&task); err != nil {
		config.WarningLogger.Printf("Failed to create task: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, task)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := models.GetTask(&task, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, task)
	}
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := models.GetTask(&task, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := c.BindJSON(&task); err != nil {
		config.InfoLogger.Printf("Failed to bind request to JSON: %v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := models.UpdateTask(&task); err != nil {
		config.WarningLogger.Printf("Failed to update task: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, task)
	}
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := models.GetTask(&task, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := models.DeleteTask(&task); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, task)
}
