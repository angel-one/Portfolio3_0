package business

import (
	"context"
	"github.com///constants"
	"github.com///external"
	"github.com///models"
	"github.com///utils/configs"
)

// GetMoxy is used to get the moxy response
func GetMoxy(ctx context.Context) (models.MoxyResponse, error) {
	response := models.MoxyResponse{}

	moxyConfig, err := configs.Get(constants.MoxyConfig)
	if err != nil {
		return response, err
	}

	data, err := external.GetMoxy(ctx, moxyConfig.GetString(constants.URLConfigKey))
	if err != nil {
		return response, err
	}

	response.Data = data

	return response, nil
}
