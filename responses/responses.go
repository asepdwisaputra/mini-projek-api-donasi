package responses

import "api-donasi/models"

type ResponseUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseCampaign struct {
	ID             int          `json:"id" form:"id"`
	Title          string       `json:"title" form:"title"`
	Description    string       `json:"description" form:"description"`
	Photo          string       `json:"photo" form:"photo"`
	TotalCollected float64      `json:"total_collected" form:"total_collected"`
	User           ResponseUser `json:"user"`
}

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
			}})
	}

	return response
}
