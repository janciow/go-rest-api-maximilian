package handlers

import (
	"go-test/internal/models"
	"go-test/internal/services"
	apperrors "go-test/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service *services.EventService
}

func NewEventHandler(service *services.EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) GetEvents(c *gin.Context) {
	events, err := h.service.GetAllEvents()
	if err != nil {
		appErr := apperrors.NewInternalError(err)
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func (h *EventHandler) GetEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		appErr := apperrors.NewBadRequestError("Invalid event ID")
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	event, err := h.service.GetEventByID(id)
	if err != nil {
		appErr := apperrors.NewInternalError(err)
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	if event == nil {
		appErr := apperrors.NewNotFoundError("Event not found")
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		appErr := apperrors.NewBadRequestError(err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	event.UserID = "user_1"
	if err := h.service.CreateEvent(&event); err != nil {
		appErr := apperrors.NewBadRequestError(err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created", "id": event.ID})
}

func (h *EventHandler) UpdateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		appErr := apperrors.NewBadRequestError("Invalid event ID")
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	var updatedEvent models.Event
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		appErr := apperrors.NewBadRequestError(err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	updatedEvent.ID = id
	if err := h.service.UpdateEvent(&updatedEvent); err != nil {
		appErr := apperrors.NewBadRequestError(err.Error())
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func (h *EventHandler) DeleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		appErr := apperrors.NewBadRequestError("Invalid event ID")
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	if err := h.service.DeleteEvent(id); err != nil {
		appErr := apperrors.NewInternalError(err)
		c.JSON(appErr.Code, gin.H{"error": appErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
