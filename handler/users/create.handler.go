package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	schemaUser "github.com/moluh/ginrest/schema"
	services "github.com/moluh/ginrest/service/user"
	util "github.com/moluh/ginrest/util"
	// gpc "github.com/go-playground/validator/v10"
)

type handler struct {
	service services.ServiceCreate
}

func NewHandlerCreateUser(service services.ServiceCreate) *handler {
	return &handler{service: service}
}

func (h *handler) CreateUserHandler(ctx *gin.Context) {

	var input schemaUser.SchemaUser
	ctx.ShouldBindJSON(&input)

	// TODO: Validation with go-playground/validator

	_, err := h.service.CreateUserService(&input)
	if err.Type == "error_02" {
		util.APIResponse(ctx, "Create new student account failed", err.Code, http.MethodPost, nil)
	}
	util.APIResponse(ctx, "Create new user account successfully", http.StatusCreated, http.MethodPost, nil)

}
