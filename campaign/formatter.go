package campaign

type CampaignFormatter struct {
	ID               int    `json: "id"`
	UserID           int    `json: "user_id"`
	Name             string `json: "name"`
	ShortDescription string `json: "short_description"`
	ImageURL         string `json: "image_url"`
	GoalAmount       int    `json: "goal_amount"`
	CurrentAmount    int    `json: "current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{}
	formatter.ID = campaign.ID
	formatter.UserID = campaign.UserID
	formatter.Name = campaign.Name
	formatter.ShortDescription = campaign.ShortDescription
	formatter.GoalAmount = campaign.GoalAmount
	formatter.CurrentAmount = campaign.CurrentAmount
	formatter.ImageURL = ""
	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	formatters := []CampaignFormatter{}

	for _, campaign := range campaigns {
		formatter := FormatCampaign(campaign)
		formatters = append(formatters, formatter)
	}

	return formatters
}
