package ips

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// Member represents a forum user
type Member struct {
	ID                    int                    `json:"id"`                    // ID number
	Name                  string                 `json:"name"`                  // Username
	Title                 string                 `json:"title"`                 // Member title
	Timezone              string                 `json:"timeZone"`              // Member timezone
	FormattedName         string                 `json:"formattedName"`         // Username with group formatting
	PrimaryGroup          Group                  `json:"primaryGroup"`          // Primary group
	SecondaryGroups       []Group                `json:"secondaryGroups"`       // Secondary groups
	Email                 string                 `json:"email"`                 // Email address
	Joined                time.Time              `json:"joined"`                // Registration date
	RegistrationIPAddress string                 `json:"registrationIpAddress"` // IP address when registered
	WarningPoints         int                    `json:"warningPoints"`         // Number of warning points the member has been issued
	ReputationPoints      int                    `json:"reputationPoints"`      // Number of reputation points member has
	PhotoURL              string                 `json:"photoUrl"`              // URL to photo (which will be the site's default if they haven't set one)
	ProfileURL            string                 `json:"profileUrl"`            // URL to profile
	Validating            bool                   `json:"validating"`            // Whether or not the validating flag is set on the member account
	Posts                 int                    `json:"posts"`                 // Number of content item submissions member has made
	LastActivity          time.Time              `json:"lastActivity"`          // Last activity date on the site
	LastVisit             time.Time              `json:"lastVisit"`             // Last distinct visit date on the site
	LastPost              time.Time              `json:"lastPost"`              // Latest content submission date
	ProfileViews          int                    `json:"profileViews"`          // Number of times member's profile has been viewed
	Birthday              string                 `json:"birthday"`              // Member birthday in MM/DD/YYYY format (or MM/DD format if no year has been supplied)
	OriginalCustomFields  map[string]FieldGroup  `json:"customFields"`          // Custom profile fields
	CustomFields          map[string]FieldGroups // some plonker decided the above field should be the worst possible format for such a simple data structure
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

// FieldGroups is a simpler structure for the above types that the official API uses. Those types
// are actually just maps keyed by numbers (y'know, like arrays, but more awkward). So this fix just
// unwraps the nested maps into flat maps and uses the names as keys.
type FieldGroups map[string]string

// GetMember implements GET /core/members/{id} and returns a Member object
// https://invisioncommunity.com/developers/rest-api?endpoint=core/members/GETitem
func (client *Client) GetMember(id string) (member Member, err error) {
	_, err = client.http.R().SetResult(&member).Get("/api/core/members/" + id)
	if err != nil {
		return
	}

	member.CustomFields = make(map[string]FieldGroups)

	for _, fieldGroup := range member.OriginalCustomFields {
		// the first layer is keyed by group name
		member.CustomFields[fieldGroup.Name] = make(map[string]string)
		for _, field := range fieldGroup.Fields {
			// the second layer is keyed by the field name, and it just contains the value
			member.CustomFields[fieldGroup.Name][field.Name] = field.Value
		}
	}

	// much simpler! access via member.CustomFields["group name"]["field name"]

	return
}

// UpdateMember implements POST /core/members/{id}
// https://invisioncommunity.com/developers/rest-api?endpoint=core/members/POSTitem
func (client *Client) UpdateMember(member Member) (err error) {
	member.OriginalCustomFields = make(map[string]FieldGroup)

	fieldGroupNumber := 0
	fieldNumber := 0
	for fieldGroupName, fieldGroupFields := range member.CustomFields {
		member.OriginalCustomFields[fmt.Sprint(fieldGroupNumber+1)] = FieldGroup{
			Name:   fieldGroupName,
			Fields: make(map[string]Field),
		}

		for fieldName, fieldValue := range fieldGroupFields {
			member.OriginalCustomFields[fmt.Sprint(fieldGroupNumber+1)].Fields[fmt.Sprint(fieldNumber+1)] = Field{
				Name:  fieldName,
				Value: fieldValue,
			}
			fieldNumber++
		}

		fieldGroupNumber++
	}

	resp, err := client.http.R().SetBody(member).Post("/api/core/members/" + fmt.Sprint(member.ID))
	if resp.StatusCode() != 200 {
		err = errors.Errorf("status not 200: ", resp.Status())
	}
	return
}
