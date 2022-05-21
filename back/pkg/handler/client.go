package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"userBar/back/models"
)

func (h *Handler) Me (c *gin.Context){
	userId, err := getUserId(c)
	fmt.Println(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}
	fmt.Println(userId)
	list, err := h.services.Client.GetById(userId)


	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) ClientList (c *gin.Context)  {
	userId, err := getUserId(c)
	fmt.Println(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}
	list, err := h.services.Client.GetList(userId)


	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) Buy (c *gin.Context)  {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "id is not found")
		return
	}
	//fmt.Println(userId)
	//list, err := h.services.Client.GetById(userId)
	var drinkId models.DrinkId
	if err := c.BindJSON(&drinkId); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.Atoi(drinkId.Id)
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
	}
	if err := h.services.UpdateUser(userId, id); err != nil{
		newErrorResponse(c, http.StatusOK, err.Error())
	}
}