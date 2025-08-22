package main

import (
	"github.com/phishr-api/mcp-server/config"
	"github.com/phishr-api/mcp-server/models"
	tools_predict "github.com/phishr-api/mcp-server/tools/predict"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_predict.CreatePost_predictTool(cfg),
	}
}
