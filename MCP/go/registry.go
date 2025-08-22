package main

import (
	"github.com/web-search-client/mcp-server/config"
	"github.com/web-search-client/mcp-server/models"
	tools_websearch "github.com/web-search-client/mcp-server/tools/websearch"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_websearch.CreateWeb_searchTool(cfg),
	}
}
