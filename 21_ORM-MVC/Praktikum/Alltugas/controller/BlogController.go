package controller

import (
	"go_Muhammad-Wahyu-Yudiansyah/21_ORM-MVC/Praktikum/Alltugas/database"
	"go_Muhammad-Wahyu-Yudiansyah/21_ORM-MVC/Praktikum/Alltugas/models"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

var blogs []models.Blog

// get all blogs
func GetBlogsController(c echo.Context) error {
	err := database.DB.Preload("User").Find(&blogs).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all blogs",
		"users":    blogs,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	blogId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid Blog ID")
	}

	if err := database.DB.Preload("User").First(&blogs, blogId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get blog by id",
		"Buku":     blogs,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	id := c.Param("id")
	var blogs models.Blog
	if err := database.DB.Where("id = ?", id).First(&blogs).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Delete(&blogs).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog by id",
		"Buku":    blogs,
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	id := c.Param("id")

	var blog models.Blog
	if err := database.DB.Model(&models.Blog{}).Where("id = ?", id).First(&blog).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if err := c.Bind(&blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := database.DB.Model(&models.Blog{}).Updates(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Blog updated successfully",
		"user":    blog,
	})
}

// create new blog
func CreateBlogController(c echo.Context) error {
	blogs := models.Blog{}
	c.Bind(&blogs)

	if err := database.DB.Save(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"Buku":    blogs,
	})
}
