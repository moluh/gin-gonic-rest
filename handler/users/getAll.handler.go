package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/moluh/ginrest/service/user"
	util "github.com/moluh/ginrest/util"
)

type handlerGetAll struct {
	service services.ServiceGetAll
}

func NewHandlerGetAllUsers(service services.ServiceGetAll) *handlerGetAll {
	return &handlerGetAll{service: service}
}

func (h *handlerGetAll) GetAllUsersHandler(ctx *gin.Context) {

	res, err := h.service.GetAllUsersService()

	switch err.Type {
	case "error_02":
		util.APIResponse(ctx, "Users data is not exists", err.Code, http.MethodPost, nil)
	default:
		util.APIResponse(ctx, "Results Users data successfully", http.StatusOK, http.MethodPost, res)
	}
}
