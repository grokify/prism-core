package core

// Layer constants represent value stream phases.
const (
	LayerRequirements = "requirements" // Product ideation, specs, design
	LayerCode         = "code"         // Application code, libraries, dependencies
	LayerInfra        = "infra"        // Cloud resources, networking, platform
	LayerRuntime      = "runtime"      // Running services, containers, workloads
	LayerAdoption     = "adoption"     // Product analytics, user engagement
	LayerSupport      = "support"      // Customer support, incident management
)

// Lifecycle stage constants for SDLC phases.
const (
	StageDesign   = "design"   // Architecture, requirements, planning
	StageBuild    = "build"    // CI/CD, code quality, dependency management
	StageTest     = "test"     // Testing coverage, quality assurance
	StageRuntime  = "runtime"  // Production monitoring, availability
	StageResponse = "response" // Incident response, remediation, recovery
)

// LayerDef defines a layer with metadata.
type LayerDef struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Order       int     `json:"order,omitempty"`
	Weight      float64 `json:"weight,omitempty"`
}

// AllLayers returns all layer constants in order.
func AllLayers() []string {
	return []string{
		LayerRequirements,
		LayerCode,
		LayerInfra,
		LayerRuntime,
		LayerAdoption,
		LayerSupport,
	}
}

// AllStages returns all lifecycle stage constants in order.
func AllStages() []string {
	return []string{
		StageDesign,
		StageBuild,
		StageTest,
		StageRuntime,
		StageResponse,
	}
}

// ValidLayer checks if a layer value is valid.
func ValidLayer(layer string) bool {
	switch layer {
	case LayerRequirements, LayerCode, LayerInfra, LayerRuntime, LayerAdoption, LayerSupport, "":
		return true
	default:
		return false
	}
}

// ValidStage checks if a lifecycle stage value is valid.
func ValidStage(stage string) bool {
	switch stage {
	case StageDesign, StageBuild, StageTest, StageRuntime, StageResponse, "":
		return true
	default:
		return false
	}
}

// LayerDisplayName returns a human-readable name for a layer.
func LayerDisplayName(layer string) string {
	names := map[string]string{
		LayerRequirements: "Requirements",
		LayerCode:         "Code",
		LayerInfra:        "Infrastructure",
		LayerRuntime:      "Runtime",
		LayerAdoption:     "Adoption",
		LayerSupport:      "Support",
	}
	if name, ok := names[layer]; ok {
		return name
	}
	return layer
}

// StageDisplayName returns a human-readable name for a stage.
func StageDisplayName(stage string) string {
	names := map[string]string{
		StageDesign:   "Design",
		StageBuild:    "Build",
		StageTest:     "Test",
		StageRuntime:  "Runtime",
		StageResponse: "Response",
	}
	if name, ok := names[stage]; ok {
		return name
	}
	return stage
}

// DefaultLayers returns default layer definitions with weights.
func DefaultLayers() []LayerDef {
	return []LayerDef{
		{ID: LayerRequirements, Name: "Requirements", Description: "Product ideation, specifications, and design", Order: 1, Weight: 0.10},
		{ID: LayerCode, Name: "Code", Description: "Application code, libraries, and dependencies", Order: 2, Weight: 0.20},
		{ID: LayerInfra, Name: "Infrastructure", Description: "Cloud resources, networking, and platform services", Order: 3, Weight: 0.20},
		{ID: LayerRuntime, Name: "Runtime", Description: "Running services, containers, and workloads", Order: 4, Weight: 0.25},
		{ID: LayerAdoption, Name: "Adoption", Description: "Product analytics, user engagement, and self-service", Order: 5, Weight: 0.15},
		{ID: LayerSupport, Name: "Support", Description: "Customer support, incident management, and escalations", Order: 6, Weight: 0.10},
	}
}
