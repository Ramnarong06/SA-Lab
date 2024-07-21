package genders


import (

   "net/http"


   "example.com/week02/config"

   "example.com/week02/models"

   "github.com/gin-gonic/gin"

)


func GetAll(c *gin.Context) {


   db := config.DB()


   var genders []models.Genders

   db.Find(&genders)


   c.JSON(http.StatusOK, &genders)


}