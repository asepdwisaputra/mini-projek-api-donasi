package responses

import (
	"api-donasi/models"
	"time"
)

type ResponseUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseCampaign struct {
	ID             int          `json:"id"`
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	Photo          string       `json:"photo"`
	TotalCollected float64      `json:"total_collected"`
	User           ResponseUser `json:"user"`
}

type ResponseDonation struct {
	ID     int       `json:"id"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
	//Status string    `json:"status"`
	User     ResponseUser     `json:"user"`
	Campaign ResponseCampaign `json:"campaign"`
}

// Fungsi Respon untuk Campaign
func GetCampaignResponse(campaigns []models.Campaign) []ResponseCampaign {
	var response []ResponseCampaign

	for _, campaign := range campaigns {
		response = append(response, ResponseCampaign{
			ID:             campaign.ID,
			Title:          campaign.Title,
			Description:    campaign.Description,
			Photo:          campaign.Photo,
			TotalCollected: campaign.TotalCollected,
			User: ResponseUser{
				ID:    campaign.User.ID,
				Name:  campaign.User.Name,
				Email: campaign.User.Email,
			},
		})
	}

	return response
}

// Fungsi Respon untuk Donation
func GetDonationResponse(donations []models.Donation) []ResponseDonation {
	var response []ResponseDonation

	for _, donation := range donations {
		response = append(response, ResponseDonation{
			ID:     donation.ID,
			Amount: donation.Amount,
			Date:   donation.Date,
			User: ResponseUser{
				ID:    donation.User.ID,
				Name:  donation.User.Name,
				Email: donation.User.Email,
			},
			Campaign: ResponseCampaign{
				ID:             donation.Campaign.ID,
				Title:          donation.Campaign.Title,
				Description:    donation.Campaign.Description,
				Photo:          donation.Campaign.Photo,
				TotalCollected: donation.Campaign.TotalCollected,
				User: ResponseUser{
					ID:    donation.User.ID,
					Name:  donation.User.Name,
					Email: donation.User.Email,
				},
			},
		})
	}
	return response
}
