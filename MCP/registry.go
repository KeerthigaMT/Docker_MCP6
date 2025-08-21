package main

import (
	"github.com/dvp-data-api/mcp-server/config"
	"github.com/dvp-data-api/mcp-server/models"
	tools_authentication "github.com/dvp-data-api/mcp-server/tools/authentication"
	tools_discovery "github.com/dvp-data-api/mcp-server/tools/discovery"
	tools_namespaces "github.com/dvp-data-api/mcp-server/tools/namespaces"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_authentication.CreatePostusers2faloginTool(cfg),
		tools_authentication.CreatePostusersloginTool(cfg),
		tools_discovery.CreateGetnamespacesTool(cfg),
		tools_discovery.CreateGetnamespaceTool(cfg),
		tools_namespaces.CreateGetnamespaceyearsTool(cfg),
		tools_namespaces.CreateGetnamespacetimespansTool(cfg),
		tools_namespaces.CreateGetnamespacetimespanmetadataTool(cfg),
		tools_namespaces.CreateGetnamespacedatabytimespanTool(cfg),
	}
}
