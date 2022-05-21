package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"userBar/back/models"
)



// регистрация
func (h *Handler) Register (c *gin.Context){
	var input models.User

	if err :=c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,

	})

}

type regInput struct {
	Name 		string 	`json:"name" binding:"required"`
	Password 	string 	`json:"password" binding:"required"`
}

func (h *Handler) Login (c *gin.Context)  {
	var input regInput
	
	if err :=c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Name, input.Password)

	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": token,

	})
}