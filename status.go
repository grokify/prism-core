package core

// Status constants for document lifecycle.
const (
	StatusDraft      = "draft"
	StatusInReview   = "in_review"
	StatusApproved   = "approved"
	StatusActive     = "active"
	StatusCompleted  = "completed"
	StatusDeprecated = "deprecated"
	StatusArchived   = "archived"
)

// HealthStatus constants for metric health indicators.
const (
	HealthGreen  = "green"
	HealthYellow = "yellow"
	HealthRed    = "red"
)

// CapabilityStatus constants for capability lifecycle.
const (
	CapabilityStatusPlanned     = "planned"
	CapabilityStatusInProgress  = "in_progress"
	CapabilityStatusImplemented = "implemented"
	CapabilityStatusOperational = "operational"
	CapabilityStatusDeprecated  = "deprecated"
)

// InitiativeStatus constants for initiative tracking.
const (
	InitiativeStatusPlanned    = "planned"
	InitiativeStatusInProgress = "in_progress"
	InitiativeStatusCompleted  = "completed"
	InitiativeStatusCancelled  = "cancelled"
	InitiativeStatusOnHold     = "on_hold"
)

// ValidHealthStatus checks if a health status value is valid.
func ValidHealthStatus(status string) bool {
	switch status {
	case HealthGreen, HealthYellow, HealthRed, "":
		return true
	default:
		return false
	}
}
