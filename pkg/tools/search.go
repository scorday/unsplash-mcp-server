package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/douglarek/unsplash-mcp-server/internal/api"
	"github.com/douglarek/unsplash-mcp-server/internal/config"
	"github.com/douglarek/unsplash-mcp-server/internal/models"
	"github.com/mark3labs/mcp-go/mcp"
)

// NewSearchPhotosTool creates and returns the search photos tool definition
func NewSearchPhotosTool() mcp.Tool {
	return mcp.NewTool("search_photos",
		mcp.WithDescription("Search for Unsplash photos"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Search keyword"),
		),
		mcp.WithNumber("page",
			mcp.Required(),
			mcp.DefaultNumber(1),
			mcp.Description("Page number (1-based)"),
		),
		mcp.WithNumber("per_page",
			mcp.Required(),
			mcp.DefaultNumber(5),
			mcp.Description("Results per page (1-30)"),
		),
		mcp.WithString("order_by",
			mcp.Required(),
			mcp.DefaultString("relevant"),
			mcp.Description("Sort method (relevant or latest)"),
		),
		mcp.WithString("color",
			mcp.Description("Color filter (black_and_white, black, white, yellow, orange, red, purple, magenta, green, teal, blue)"),
		),
		mcp.WithString("orientation",
			mcp.Description("Orientation filter (landscape, portrait, squarish)"),
		),
	)
}

// HandleSearchPhotos returns a handler function for search photos requests
func HandleSearchPhotos(cfg *config.Config) func(context.Context, mcp.CallToolRequest, models.SearchPhotosRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest, args models.SearchPhotosRequest) (*mcp.CallToolResult, error) {
		// Create Unsplash client
		client := api.NewClient(cfg)

		// Extract and validate parameters
		if args.Query == "" {
			return mcp.NewToolResultError("Search keyword must be provided"), nil
		}

		// Build query parameters
		params := buildSearchParams(args)

		// Search photos
		photos, err := client.SearchPhotos(ctx, params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to search photos: %v", err)), nil
		}

		// Format results
		b, _ := json.Marshal(photos)
		return mcp.NewToolResultText(string(b)), nil
	}
}

// buildSearchParams builds URL parameters from request arguments
func buildSearchParams(args models.SearchPhotosRequest) url.Values {
	params := url.Values{}

	// Required parameters
	params.Add("query", args.Query)
	params.Add("page", strconv.Itoa(args.Page))
	params.Add("per_page", strconv.Itoa(args.PerPage))
	params.Add("order_by", args.OrderBy)

	// Optional parameters
	if args.Color != "" {
		params.Add("color", args.Color)
	}

	if args.Orientation != "" {
		params.Add("orientation", args.Orientation)
	}

	return params
}
