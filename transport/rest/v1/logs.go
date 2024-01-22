package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"userLoggingProject/pkg/logger"
	"userLoggingProject/transport/rest/v1/requests"
)

func (h *Handler) initLogRoutes(api *gin.RouterGroup) {
	logs := api.Group("/logs")
	{
		logs.GET("/:id", h.getAll)
		logs.POST("/getById/", h.getById)
		logs.DELETE("/:id", h.removeAll)
		logs.POST("/", h.add)
	}
}

// @Summary Get all logs
// @Tags logs
// @Description  get all user logs
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {array}  LogEntry.LogEntry
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /logs/{id} [get]
func (h *Handler) getAll(c *gin.Context) {
	logs, err := h.services.Logs.GetAll(c.GetString("id"))
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: logs})
}

// @Summary Get Log By id
// @Tags logs
// @Description  get log by id
// @Accept  json
// @Produce  json
// @Param body body requests.GetByIdRequest true "id"
// @Success 200 {object}  LogEntry.LogEntry
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /logs/getById [post]
func (h *Handler) getById(c *gin.Context) {
	var request requests.GetByIdRequest
	if err := c.BindJSON(&request); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}
	log, err := h.services.Logs.GetById(request.UserId, request.LogId)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: log})
}

// @Summary Remove All User Logs
// @Tags logs
// @Description  remove all user logs
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /logs/{id} [delete]
func (h *Handler) removeAll(c *gin.Context) {
	err := h.services.Logs.ClearLogs(c.GetString("id"))
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{})
}

// @Summary Add Log
// @Tags logs
// @Description  Add log
// @Accept  json
// @Produce  json
// @Param body body requests.AddLogRequest true "log"
// @Success 200 {object} int
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /logs/ [post]
func (h *Handler) add(c *gin.Context) {
	var request requests.AddLogRequest
	if err := c.BindJSON(&request); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		logger.Error(err)
		return
	}
	logId, err := h.services.Logs.AddLog(request.UserId, &request.Log)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, dataResponse{Data: logId})
}
