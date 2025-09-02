package routes

import (
	"net/http"
	"strconv"

	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

// registerForEvent godoc
// @Summary Register a user for an event
// @Description Registers the authenticated user for the specified event
// @Tags Events
// @Accept  json
// @Produce  json
// @Param   id   path     int  true  "Event ID"
// @Success 201  {object} map[string]string
// @Failure 400  {object} map[string]string
// @Failure 500  {object} map[string]string
// @Security BearerAuth
// @Router /events/{id}/register [post]
func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
	}


	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user registered!"})
}

// cancelRegistration godoc
// @Summary Cancel user registration
// @Description Cancels the authenticated user's registration for the specified event
// @Tags Events
// @Accept  json
// @Produce  json
// @Param   id   path     int  true  "Event ID"
// @Success 200  {object} map[string]string
// @Failure 400  {object} map[string]string
// @Failure 500  {object} map[string]string
// @Security BearerAuth
// @Router /events/{id}/cancel [delete]
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "registration canceled!"})

}