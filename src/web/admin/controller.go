package admin

import (
	"database/sql"
	"fmt"
	"otumian-empire/go-ecom/src/handlers"
	"otumian-empire/go-ecom/src/model"
	"otumian-empire/go-ecom/src/utils"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	dao Dao
}

func NewController(db *sql.DB) Controller {
	return Controller{
		dao: NewDao(db),
	}
}
		}

		// validate request body
		if !utils.IsValidEmail(payload.Email) {
			context.JSON(handlers.FailureMessageResponse("Enter a valid email"))
			return
		}

		if !utils.HasValidSize(payload.Password, 5) {
			context.JSON(handlers.FailureMessageResponse("Enter a password of at least 5 characters"))
			return
		}

		// authenticate the credentials
		row, err := controller.dao.FindOneByEmail(payload.Email)

		if err != nil {
			context.JSON(handlers.FailureMessageResponse(fmt.Sprintf("Invalid credentials: %v", err.Error())))
			return
		}

		if isValidPassword, err := utils.ComparePassword(payload.Password, row.Password); err != nil || !isValidPassword {
			fmt.Println(err)
			context.JSON(handlers.FailureMessageResponse("Invalid credentials: password do not match " + err.Error()))
			return
		}

		token, err := utils.GenerateJwt(row.Id)
		if err != nil {
			context.JSON(handlers.FailureMessageResponse("Invalid credentials: password do not match: " + err.Error()))
			return
		}

		context.JSON(handlers.SuccessResponse("Login successful", LoginResponseMapper(row, token)))

	}
}

func (controller *Controller) CreatedAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		// the admins role must be check to make sure that the admin
		// is a super admin before they can add a new another admin

		user, isUser := context.MustGet("user").(model.Admin)
		if !isUser || user.Role != SUPER_ADMIN {
			context.JSON(handlers.FailureMessageResponse("Not authorized"))
			return
		}

		// get the request body
		var payload CreatedAdminDto

		if err := context.BindJSON(&payload); err != nil {
			context.JSON(handlers.FailureMessageResponse("Error occurred while parsing body"))
			return
		}

		// validate the request body
		if !utils.IsValidEmail(payload.Email) {
			context.JSON(handlers.FailureMessageResponse("Enter a valid email"))
			return
		}

		if !utils.HasValidSize(payload.FullName, 5) {
			context.JSON(handlers.FailureMessageResponse("Enter a full name of at least 5 characters"))
			return
		}

		// check that the role is one of super_admin|mod
		if payload.Role != SUPER_ADMIN && payload.Role != MODERATOR {
			context.JSON(handlers.FailureMessageResponse(fmt.Sprintf("Role can be one of these: %v, %v", SUPER_ADMIN, MODERATOR)))
			return
		}

		// authenticate the email to make sure that the email doesn't
		// already exist
		if _, err := controller.dao.FindOneByEmail(payload.Email); err == nil {
			context.JSON(handlers.FailureMessageResponse("Email already taken"))
			return
		}

		// create admin
		if err := controller.dao.Create(payload); err != nil {
			context.JSON(handlers.FailureMessageResponse(err.Error()))
			return
		}

		// TODO: send admin email with default password to reset the password

		context.JSON(handlers.SuccessMessageResponse("Admin created successfully"))
	}
}

// here the admin can update their own full name
func (controller *Controller) UpdateProfile() gin.HandlerFunc {
	return func(context *gin.Context) {
		// get the full name passed in the body
		var payload UpdateProfileDto

		if err := context.BindJSON(&payload); err != nil {
			context.JSON(handlers.FailureMessageResponse(fmt.Sprintf("Error occurred parsing body, %v", err)))
			return
		}

		// validate the full name
		if !utils.HasValidSize(payload.FullName, 5) {
			context.JSON(handlers.FailureMessageResponse("Full name must be at least 5 characters long"))
			return
		}

		// fetch the user by id (from the user object on request or session)
		userId := user.Id

		if _, err := controller.dao.FindOneById(userId); err != nil {
			context.JSON(handlers.FailureMessageResponse(fmt.Sprintf("An error occurred reading user, %v", err)))
			return
		}

		// update the admins full name
		if err := controller.dao.Update(userId, payload.FullName); err != nil {
			context.JSON(handlers.FailureMessageResponse(fmt.Sprintf("An error occurred updating user full name, %v", err)))
			return
		}

		context.JSON(handlers.SuccessMessageResponse("Full name updated successfully"))
	}
}

func (controller *Controller) UpdateAdminRole() gin.HandlerFunc {
	return func(context *gin.Context) {

		// // get the full name passed in the body
		// var payload UpdateProfileDto

		// if err := context.BindJSON(&payload); err != nil {
		// 	context.JSON(handlers.FailureMessageResponse(fmt.Sprintf("Error occurred parsing body, %v", err)))
		// 	return
		// }

		// // validate the full name
		// if !HasValidSize(payload.FullName, 5) {
		// 	context.JSON(handlers.FailureMessageResponse("Full name must be at least 5 characters long"))
		// 	return
		// }

		// // fetch the user by id (from the user object on request or session)
		// // TODO: Get the userId from the session/req
		// const userId = 1

		// if _, err := controller.dao.FindOneById(userId); err != nil {
		// 	context.JSON(handlers.FailureMessageResponse(fmt.Sprintf("An error occurred reading user, %v", err)))
		// 	return
		// }

		// // update the admins full name
		// if err := controller.dao.Update(userId, payload.FullName); err != nil {
		// 	context.JSON(handlers.FailureMessageResponse(fmt.Sprintf("An error occurred updating user full name, %v", err)))
		// 	return
		// }

		context.JSON(handlers.SuccessMessageResponse("Full name updated successfully"))
	}
}

func (controller *Controller) ForgetPassword() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"ping": "pong"})
	}
}

func (controller *Controller) ResetPassword() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"ping": "pong"})
	}
}

func (controller *Controller) BlockUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"ping": "pong"})
	}
}

func (controller *Controller) BlockAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"ping": "pong"})
	}
}
