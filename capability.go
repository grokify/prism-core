package core

// CapabilityRelation values describe how one entity relates to a
// capability, shared across the PRISM ecosystem (prism-capability,
// prism-roadmap, prism-maturity) so each module references capabilities
// the same way instead of inventing its own relation vocabulary.
const (
	// CapabilityRelationEnables means the referencing entity creates a
	// capability that did not exist before.
	CapabilityRelationEnables = "enables"

	// CapabilityRelationImproves means the referencing entity strengthens
	// an existing capability without creating a new one.
	CapabilityRelationImproves = "improves"

	// CapabilityRelationDependsOn means the referencing entity requires an
	// existing capability to be in place before it can proceed.
	CapabilityRelationDependsOn = "dependsOn"
)

// AllCapabilityRelations returns all capability relation constants.
func AllCapabilityRelations() []string {
	return []string{
		CapabilityRelationEnables,
		CapabilityRelationImproves,
		CapabilityRelationDependsOn,
	}
}

// ValidCapabilityRelation checks if a capability relation value is valid.
func ValidCapabilityRelation(relation string) bool {
	switch relation {
	case CapabilityRelationEnables, CapabilityRelationImproves, CapabilityRelationDependsOn:
		return true
	default:
		return false
	}
}

// CapabilityRelationDisplayName returns a human-readable name for a
// capability relation.
func CapabilityRelationDisplayName(relation string) string {
	names := map[string]string{
		CapabilityRelationEnables:   "Enables",
		CapabilityRelationImproves:  "Improves",
		CapabilityRelationDependsOn: "Depends On",
	}
	if name, ok := names[relation]; ok {
		return name
	}
	return relation
}

// CapabilityRef references one capability by ID with the relation an
// entity has to it — a thin, ecosystem-shared shape any PRISM module can
// use to reference a capability without depending on prism-capability's
// full domain model (CapabilityStack, Layer, Category, and so on).
type CapabilityRef struct {
	CapabilityID string `json:"capabilityId"`
	Relation     string `json:"relation"`
	Rationale    string `json:"rationale,omitempty"`
}
