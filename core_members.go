package ips

import (
	"time"
)

// Member represents a forum user
type Member struct {
	ID                    int                   `json:"id"`                    // ID number
	Name                  string                `json:"name"`                  // Username
	Title                 string                `json:"title"`                 // Member title
	Timezone              string                `json:"timeZone"`              // Member timezone
	FormattedName         string                `json:"formattedName"`         // Username with group formatting
	PrimaryGroup          Group                 `json:"primaryGroup"`          // Primary group
	SecondaryGroups       []Group               `json:"secondaryGroups"`       // Secondary groups
	Email                 string                `json:"email"`                 // Email address
	Joined                time.Time             `json:"joined"`                // Registration date
	RegistrationIPAddress string                `json:"registrationIpAddress"` // IP address when registered
	WarningPoints         int                   `json:"warningPoints"`         // Number of warning points the member has been issued
	ReputationPoints      int                   `json:"reputationPoints"`      // Number of reputation points member has
	PhotoURL              string                `json:"photoUrl"`              // URL to photo (which will be the site's default if they haven't set one)
	ProfileURL            string                `json:"profileUrl"`            // URL to profile
	Validating            bool                  `json:"validating"`            // Whether or not the validating flag is set on the member account
	Posts                 int                   `json:"posts"`                 // Number of content item submissions member has made
	LastActivity          time.Time             `json:"lastActivity"`          // Last activity date on the site
	LastVisit             time.Time             `json:"lastVisit"`             // Last distinct visit date on the site
	LastPost              time.Time             `json:"lastPost"`              // Latest content submission date
	ProfileViews          int                   `json:"profileViews"`          // Number of times member's profile has been viewed
	Birthday              string                `json:"birthday"`              // Member birthday in MM/DD/YYYY format (or MM/DD format if no year has been supplied)
	CustomFields          map[string]FieldGroup `json:"customFields"`          // Custom profile fields

}

// FieldGroup represents a group of custom fields in a member's profile
type FieldGroup struct {
	Name   string           `json:"name"`   // Group name
	Fields map[string]Field `json:"fields"` // Array of field objects
}

// Field represents a custom field in a member's profile
type Field struct {
	Name  string `json:"name"`  // Field name
	Value string `json:"value"` // Field value
}

// GetMember implements /core/members/{id} and returns a Member object
// https://invisioncommunity.com/developers/rest-api?endpoint=core/members/GETitem
func (client *Client) GetMember(id string) (member Member, err error) {
	_, err = client.http.R().SetResult(&member).Get("/api/core/members/" + id)
	if err != nil {
		return
	}

	return
}
