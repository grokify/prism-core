package core

import (
	"fmt"
	"strings"
	"time"
)

// Person represents an individual.
type Person struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}

// Approver represents a person who approved something.
type Approver struct {
	Person
	ApprovedAt time.Time `json:"approvedAt,omitzero"`
	Approved   bool      `json:"approved,omitempty"`
	Comments   string    `json:"comments,omitempty"`
}

// Team represents a group of people.
type Team struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type,omitempty"` // stream_aligned, platform, enabling, overlay
	Domain      string   `json:"domain,omitempty"`
	Owner       string   `json:"owner,omitempty"`
	Email       string   `json:"email,omitempty"`
	Slack       string   `json:"slack,omitempty"`
	Members     []Person `json:"members,omitempty"`
}

// Team type constants following Team Topologies.
const (
	TeamTypeStreamAligned = "stream_aligned"
	TeamTypePlatform      = "platform"
	TeamTypeEnabling      = "enabling"
	TeamTypeOverlay       = "overlay"
)

// ValidTeamType checks if a team type is valid.
func ValidTeamType(teamType string) bool {
	switch teamType {
	case TeamTypeStreamAligned, TeamTypePlatform, TeamTypeEnabling, TeamTypeOverlay, "":
		return true
	default:
		return false
	}
}

// FormatPersonMarkdown formats a person for markdown output.
func FormatPersonMarkdown(p Person) string {
	if p.Email != "" {
		if p.Role != "" {
			return fmt.Sprintf("[%s](mailto:%s) (%s)", p.Name, p.Email, p.Role)
		}
		return fmt.Sprintf("[%s](mailto:%s)", p.Name, p.Email)
	}
	if p.Role != "" {
		return fmt.Sprintf("%s (%s)", p.Name, p.Role)
	}
	return p.Name
}

// FormatPeopleMarkdown formats a list of people for markdown output.
func FormatPeopleMarkdown(people []Person) string {
	if len(people) == 0 {
		return ""
	}
	var parts []string
	for _, p := range people {
		parts = append(parts, FormatPersonMarkdown(p))
	}
	return strings.Join(parts, ", ")
}

// NewApprover creates an approver from a person.
func NewApprover(p Person, approved bool, comments string) Approver {
	return Approver{
		Person:     p,
		ApprovedAt: time.Now(),
		Approved:   approved,
		Comments:   comments,
	}
}
