package route

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/moluh/ginrest/handler/users"
	repository "github.com/moluh/ginrest/repository/users"
	services "github.com/moluh/ginrest/service/user"
	"gorm.io/gorm"
)

func InitUsersRoute(db *gorm.DB, route *gin.Engine) {

	// Create
	createRepository := repository.NewRepositoryCreate(db)
	createService := services.NewServiceCreate(createRepository)
	createHandler := handlers.NewHandlerCreateUser(createService)
	// Get All
	getAllRepository := repository.NewRepositoryGetAll(db)
	getAllService := services.NewServiceGetAll(getAllRepository)
	getAllHandler := handlers.NewHandlerGetAllUsers(getAllService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/user", createHandler.CreateUserHandler)
	groupRoute.GET("/user", getAllHandler.GetAllUsersHandler)

}
