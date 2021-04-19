package adminController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/models"
	"wejh-go/app/services/applistServices"
	"wejh-go/app/utils"
)

type createApplistForm struct {
	Title string `json:"title" binding:"required"`
	Route string `json:"route"`
}
type updateApplistForm struct {
	ID    int64  `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Route string `json:"route"`
}
type deleteApplistForm struct {
	ID int64 `json:"id" binding:"required"`
}

func CreateApplist(c *gin.Context) {
	var postForm createApplistForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}

	err = applistServices.CreateApplist(models.AppList{Title: postForm.Title, Route: postForm.Route})
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
func UpdateApplist(c *gin.Context) {
	var postForm updateApplistForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}

	err = applistServices.UpdateApplist(postForm.ID, models.AppList{
		Title: postForm.Title,
		Route: postForm.Route,
	},
	)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
func DeleteApplist(c *gin.Context) {
	var postForm deleteApplistForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}

	err = applistServices.DeleteApplist(postForm.ID)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
