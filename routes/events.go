package routes

import (
	"event-planning/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func handleGetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents();
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong", "error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "All events returned", "data": events})
}

func handlePostEvents (ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error saving event", "error": err.Error() })
		return
	}

	userId := ctx.GetInt64("userId")

	event.CreatedBy = userId
	err = event.SaveEvents()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error saving events", "error": err.Error() })
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event saved", "data": event})
}

func handleGetEventByID(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse id"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err.Error() })
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Event returned", "data": event})
}

func handleUpdateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse id"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Event not found"})
		return;
	}

	err = ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error updating event", "error": err.Error() })
		return
	}

	userId := ctx.GetInt64("userId")

	if userId != event.CreatedBy {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update this event"})
		return
	}

	err = event.UpdateEventByID()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Error updating event", "error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Event updated" })
}

func handleDeleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Error parsing event id", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Error getting event", "error": err.Error()})
		return
	}

	userId := ctx.GetInt64("userId")

	if userId != event.CreatedBy {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete this event"})
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Error deleting event", "error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Event deleted" })
}