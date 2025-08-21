package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dvp-data-api/mcp-server/config"
	"github.com/dvp-data-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetnamespacetimespanmetadataHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		namespaceVal, ok := args["namespace"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: namespace"), nil
		}
		namespace, ok := namespaceVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: namespace"), nil
		}
		yearVal, ok := args["year"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: year"), nil
		}
		year, ok := yearVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: year"), nil
		}
		timespantypeVal, ok := args["timespantype"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: timespantype"), nil
		}
		timespantype, ok := timespantypeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: timespantype"), nil
		}
		timespanVal, ok := args["timespan"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: timespan"), nil
		}
		timespan, ok := timespanVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: timespan"), nil
		}
		url := fmt.Sprintf("%s/namespaces/%s/pulls/exports/years/%s/%s/%s", cfg.BaseURL, namespace, year, timespantype, timespan)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.TimespanModel
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetnamespacetimespanmetadataTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_namespaces_namespace_pulls_exports_years_year_timespantype_timespan",
		mcp.WithDescription("Get namespace metadata for timespan"),
		mcp.WithString("namespace", mcp.Required(), mcp.Description("Namespace to fetch data for")),
		mcp.WithNumber("year", mcp.Required(), mcp.Description("Year to fetch data for")),
		mcp.WithString("timespantype", mcp.Required(), mcp.Description("Type of timespan to fetch data for")),
		mcp.WithNumber("timespan", mcp.Required(), mcp.Description("Timespan to fetch data for")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetnamespacetimespanmetadataHandler(cfg),
	}
}
