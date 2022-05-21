package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"userBar/back/models"
)

func (h *Handler) BarmanList (c *gin.Context){

	barId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	list, err := h.services.Barman.GetList(barId)

	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) Create (c *gin.Context)  {
	var input models.Drinks

	barId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Barman.CreateDrink(barId, input.Item, input.Promille, input.Cost)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})
}