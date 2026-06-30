package models

type SearchPhotosRequest struct {
	Query       string `json:"query"`
	Page        int    `json:"page"`
	PerPage     int    `json:"per_page"`
	OrderBy     string `json:"order_by"`
	Color       string `json:"color"`
	Orientation string `json:"orientation"`
}

// SearchResponse represents the Unsplash search API response
type SearchResponse struct {
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Results    []Photo `json:"results"`
}

// Photo represents an Unsplash photo
type Photo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Urls        URLs   `json:"urls"`
	User        User   `json:"user"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

// URLs contains various size URLs for a photo
type URLs struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}

// User represents the Unsplash user attached to a photo result
type User struct {
	ID                string       `json:"id"`
	Username          string       `json:"username"`
	Name              string       `json:"name"`
	FirstName         string       `json:"first_name"`
	LastName          string       `json:"last_name"`
	InstagramUsername string       `json:"instagram_username"`
	TwitterUsername   string       `json:"twitter_username"`
	PortfolioURL      string       `json:"portfolio_url"`
	ProfileImage      ProfileImage `json:"profile_image"`
	Links             UserLinks    `json:"links"`
}

// ProfileImage contains the user's avatar URLs
type ProfileImage struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

// UserLinks contains the user's related Unsplash URLs
type UserLinks struct {
	Self      string `json:"self"`
	HTML      string `json:"html"`
	Photos    string `json:"photos"`
	Likes     string `json:"likes"`
	Portfolio string `json:"portfolio"`
}
