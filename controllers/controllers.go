package controllers

import (
	"api-donasi/config"
	"api-donasi/middleware"
	"api-donasi/models"
	"api-donasi/responses"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get All Users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID to Get User",
		})
	}
	// Cek id di user
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"error": "User not found",
		})
	}
	// Simpan perubahan ke database
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to update user",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get User",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create New User",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID to Delete User",
		})
	}
	// Cek id di user
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"error": "User Not Found",
		})
	}

	// Soft Delete user
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to Soft Delete User",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Soft Delete User",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}
	// Cek id di user
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"error": "User Not Found",
		})
	}
	// Binding datanya
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid Data",
		})
	}
	// Simpan
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update User",
		"user":    user,
	})
}

// Login user
func LoginUserController(c echo.Context) error {
	// Binding data
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password =?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to Login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to Login",
			"error":   err.Error(),
		})
	}

	userResponJWT := models.UserResponseJWT{user.ID, user.Name, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create User",
		"user":    userResponJWT,
	})
}

// Membuat kampanye baru
func CreateCampaign(c echo.Context) error {
	// Mendapatkan data kampanye yang baru dibuat
	newCampaign := &models.Campaign{}
	if err := c.Bind(newCampaign); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Campaign Data Not Valid"})
	}

	// Periksa apakah pengguna dengan ID yang sesuai ada
	var user models.User
	if err := config.DB.Where("ID = ?", newCampaign.UserID).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID Not Valid"})
	}

	// Menyimpan kampanye ke database
	if err := config.DB.Create(newCampaign).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error Creating Campaign"})
	}

	// Kemudian, mengambil data kampanye dengan Preload
	if err := config.DB.Preload("User").First(newCampaign).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// Kasus ini terjadi jika data kampanye tidak ditemukan.
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Campaign Not Found"})
		} else {
			// Kesalahan lain yang mungkin terjadi selain RecordNotFoundError.
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error Geting Campaign Data"})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success Create New Campaign",
		"campaign": newCampaign,
	})
}

// Mengambil Semua Campaign
func GetCampaigns(c echo.Context) error {
	var campaigns []models.Campaign

	if err := config.DB.Preload("User").Find(&campaigns).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Membuat struktur data baru untuk respons dengan key mapping dalam huruf kecil
	// var response struct {
	// 	Message   string            `json:"message"`
	// 	Campaigns []responses.ResponseCampaign `json:"campaigns"`
	// }

	// response.Message = "Success Get All Campaign"
	// response.Campaigns = campaigns

	response := responses.GetCampaignResponse(campaigns)

	return c.JSON(http.StatusOK, response)
}

// Mengambil kampanye berdasar id
func GetCampaign(c echo.Context) error {
	campaignID := c.Param("id")

	var campaign []models.Campaign
	if err := config.DB.Preload("User").Where("id = ?", campaignID).First(&campaign).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Kampanye tidak ditemukan"})
	}

	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"message":  "Success Get Campaign",
	// 	"campaign": campaign,
	// })
	response := responses.GetCampaignResponse(campaign)

	return c.JSON(http.StatusOK, response)
}

// Mendapatkan donasi yang baru dibuat
func CreateDonation(c echo.Context) error {
	newDonation := &models.Donation{}
	if err := c.Bind(newDonation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Data Not Valid"})
	}

	// Mendapatkan kampanye yang sesuai dari database
	campaign := &models.Campaign{}
	if err := config.DB.Where("ID = ?", newDonation.CampaignID).First(campaign).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Campaign Not Found"})
	}

	// Menambahkan nilai Amount ke TotalCollected
	campaign.TotalCollected += newDonation.Amount

	// Menyimpan perubahan ke database
	if err := config.DB.Save(campaign).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error Saving Campaign"})
	}

	// Menyimpan donasi ke database
	if err := config.DB.Create(newDonation).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating Donation"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success Create Donation",
		"donation": newDonation,
	})
}

// Mengambil semua data donasi
func GetDonations(c echo.Context) error {
	var donations []models.Donation
	if err := config.DB.Find(&donations).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error Getting Donation"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success Get Donations",
		"donation": donations,
	})
}

// Mengambil data donasi berdasar id
func GetDonationByID(c echo.Context) error {
	// Mendapatkan ID donasi dari parameter URL
	donationID := c.Param("id")

	// Membuat objek Donasi untuk menampung hasil
	var donation models.Donation

	// Mengambil donasi berdasarkan ID
	if err := config.DB.First(&donation, donationID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// Kasus donasi tidak ditemukan
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Donation Not Found"})
		} else {
			// Kesalahan lain yang mungkin terjadi
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error Getting Data Donation"})
		}
	}

	// Jika donasi ditemukan, kembalikan respons JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success Get Donation By Id",
		"donation": donation,
	})
}

// Mengambil semua data donasi berdasarkan ID pengguna
func GetDonationsByUserID(c echo.Context) error {
	// Mendapatkan ID pengguna dari parameter URL
	userID := c.Param("id")

	// Membuat slice untuk menampung donasi oleh pengguna
	var donations []models.Donation

	// Mengambil donasi berdasarkan ID pengguna
	if err := config.DB.Where("user_id = ?", userID).Find(&donations).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error getting Data Donation"})
	}

	// Kembalikan daftar donasi dalam format JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success Get Donation By User Id",
		"donation": donations,
	})
}
