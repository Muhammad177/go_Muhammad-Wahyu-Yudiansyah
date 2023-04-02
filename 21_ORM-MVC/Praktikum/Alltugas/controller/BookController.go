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

var books []models.Book

// get all books
func GetBooksController(c echo.Context) error {
	err := database.DB.Find(&books).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all books",
		"users":    books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid Book ID")
	}

	if err := database.DB.First(&books, bookId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get book by id",
		"Buku":     books,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	var books models.Book
	if err := database.DB.Where("id = ?", id).First(&books).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := database.DB.Delete(&books).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book by id",
		"Buku":    books,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	id := c.Param("id")

	var book models.Book
	if err := database.DB.Model(&models.Book{}).Where("id = ?", id).First(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Book not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := database.DB.Model(&models.Book{}).Updates(book).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book updated successfully",
		"user":    book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	books := models.Book{}
	c.Bind(&books)

	if err := database.DB.Save(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"Buku":    books,
	})
}
