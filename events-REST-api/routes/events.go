package routes

import (
	"net/http"
	"strconv"

	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

// @Summary Get all events
// @Description Returns a list of all events
// @Tags Events
// @Produce json
// @Success 200 {array} models.Event
// @Failure 500 {object} map[string]string
// @Router /events [get]
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events. Try again later!"})
		return
	}
	context.JSON(http.StatusOK, events)
}

// @Summary Get a single event
// @Description Returns details of an event by ID
// @Tags Events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} models.Event
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /events/{id} [get]
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

// @Summary Create a new event
// @Description Creates a new event for the authenticated user
// @Tags Events
// @Accept json
// @Produce json
// @Param event body models.Event true "Event Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /events [post]
func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "couldn't parse the event data"})
		return
	}
	userId := context.GetInt64("userId")
	event.UserId = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event. Try again later!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

// @Summary Update an event
// @Description Updates an existing event (only by owner)
// @Tags Events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param event body models.Event true "Updated Event Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /events/{id} [put]
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "couldn't parse the event data"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}

	if userId != event.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update the event"})
		return
	}
	

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "couldn't parse the request data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated Successfully!"})
}

// @Summary Delete an event
// @Description Deletes an event (only by owner)
// @Tags Events
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /events/{id} [delete]
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "couldn't parse the event data"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}

	if userId != event.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete the event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted Successfully!"})
}