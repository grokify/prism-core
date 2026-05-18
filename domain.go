// Package core provides shared primitives for the PRISM ecosystem.
//
// The PRISM ecosystem consists of three modules:
//   - prism-capability: What capabilities exist (capability stacks, layers)
//   - prism-intelligence: How we measure maturity (SLIs, SLOs, maturity models)
//   - prism-execution: How we improve (OKRs, V2MOM, roadmaps)
//
// This package provides shared types used across all three modules.
package core

// Domain constants represent functional areas with their own standards.
const (
	// Primary domains
	DomainSecurity   = "security"   // Application and infrastructure security
	DomainOperations = "operations" // Reliability, performance, efficiency
	DomainQuality    = "quality"    // Testing, code quality, defects

	// Extended domains
	DomainPlatform       = "platform"       // Platform engineering
	DomainAI             = "ai"             // AI/ML capabilities
	DomainData           = "data"           // Data management
	DomainObservability  = "observability"  // Monitoring and observability
	DomainInfrastructure = "infrastructure" // Infrastructure management
	DomainProduct        = "product"        // Product management
	DomainCompliance     = "compliance"     // Compliance and governance
)

// AllDomains returns all available domain constants.
func AllDomains() []string {
	return []string{
		DomainSecurity,
		DomainOperations,
		DomainQuality,
		DomainPlatform,
		DomainAI,
		DomainData,
		DomainObservability,
		DomainInfrastructure,
		DomainProduct,
		DomainCompliance,
	}
}

// PrimaryDomains returns the three primary domains.
func PrimaryDomains() []string {
	return []string{
		DomainSecurity,
		DomainOperations,
		DomainQuality,
	}
}

// ValidDomain checks if a domain value is valid.
func ValidDomain(domain string) bool {
	switch domain {
	case DomainSecurity, DomainOperations, DomainQuality,
		DomainPlatform, DomainAI, DomainData,
		DomainObservability, DomainInfrastructure,
		DomainProduct, DomainCompliance:
		return true
	default:
		return false
	}
}

// DomainDisplayName returns a human-readable name for a domain.
func DomainDisplayName(domain string) string {
	names := map[string]string{
		DomainSecurity:       "Security",
		DomainOperations:     "Operations",
		DomainQuality:        "Quality",
		DomainPlatform:       "Platform",
		DomainAI:             "AI/ML",
		DomainData:           "Data",
		DomainObservability:  "Observability",
		DomainInfrastructure: "Infrastructure",
		DomainProduct:        "Product",
		DomainCompliance:     "Compliance",
	}
	if name, ok := names[domain]; ok {
		return name
	}
	return domain
}
