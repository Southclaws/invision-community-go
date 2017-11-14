package ips

// Group represents a forum member group
type Group struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FormattedName string `json:"formattedName"`
}
