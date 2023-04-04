package midleware

import (
	"fmt"
	"go_Muhammad-Wahyu-Yudiansyah/22_Midleware/Praktikum/Alltugas/database"
	"go_Muhammad-Wahyu-Yudiansyah/22_Midleware/Praktikum/Alltugas/models"

	"github.com/labstack/echo/v4"
)

func BasicAuthDb(username, password string, c echo.Context) (bool, error) {
	var user models.User
	err := database.DB.Where("email = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return false, fmt.Errorf("invalid username or password")
	}
	return true, nil
}
