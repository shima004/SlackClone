package delivery

// import (
// 	"github.com/labstack/echo/v4"
// 	"github.com/shima004/slackclone/entities"
// 	"github.com/shima004/slackclone/usecases/inputport"
// )

// type UserHandler struct {
// 	UserUsecase inputport.UserUsecase
// }

// func NewUserHandler(g *echo.Group, userUsecase inputport.UserUsecase) *UserHandler {
// 	handler := &UserHandler{UserUsecase: userUsecase}
// 	g.POST("/users", handler.CreateUser)
// 	g.DELETE("/users/:id", handler.DeleteUser)
// 	g.POST("/login", handler.Login)
// 	return handler
// }

// func (uh *UserHandler) CreateUser(c echo.Context) error {
// 	var user entities.User
// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}
// 	if ok, err := isRequestValid(&user); !ok {
// 		return err
// 	}
// 	if err := uh.UserUsecase.CreateUser(c.Request().Context(), user); err != nil {
// 		return err
// 	}
// 	return c.JSON(201, user)
// }

// func (uh *UserHandler) DeleteUser(c echo.Context) error {
// 	sUserID := c.Param("id")
// 	userID, err := StringToUint(sUserID)
// 	if err != nil {
// 		return err
// 	}
// 	if err := uh.UserUsecase.DeleteUser(c.Request().Context(), userID); err != nil {
// 		return err
// 	}
// 	return c.NoContent(204)
// }

// func (uh *UserHandler) Login(c echo.Context) error {
// 	var user entities.User
// 	if err := c.Bind(&user); err != nil {
// 		return err
// 	}
// 	if ok, err := isRequestValid(&user); !ok {
// 		return err
// 	}
// 	token, err := uh.UserUsecase.Login(c.Request().Context(), user.Email, user.Password)
// 	if err != nil {
// 		return err
// 	}
// 	return c.JSON(200, token)
// }
