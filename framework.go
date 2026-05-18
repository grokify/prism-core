package core

// Framework constants for compliance and methodology mappings.
const (
	// NIST Frameworks
	FrameworkNISTCSF    = "NIST_CSF"    // NIST Cybersecurity Framework 1.1
	FrameworkNISTCSF2   = "NIST_CSF_2"  // NIST Cybersecurity Framework 2.0
	FrameworkNIST80053  = "NIST_800_53" // NIST 800-53 Security Controls
	FrameworkNIST800171 = "NIST_800_171"
	FrameworkNISTRMF    = "NIST_RMF"    // NIST Risk Management Framework
	FrameworkNISTAIRMF  = "NIST_AI_RMF" // NIST AI Risk Management Framework

	// FedRAMP Baselines
	FrameworkFedRAMPHigh = "FEDRAMP_HIGH"
	FrameworkFedRAMPMod  = "FEDRAMP_MOD"
	FrameworkFedRAMPLow  = "FEDRAMP_LOW"

	// Industry Standards
	FrameworkISO27001    = "ISO_27001"
	FrameworkSOC2        = "SOC_2"
	FrameworkCISControls = "CIS_CONTROLS"
	FrameworkPCIDSS      = "PCI_DSS"
	FrameworkHIPAA       = "HIPAA"
	FrameworkGDPR        = "GDPR"

	// Security Frameworks
	FrameworkMITREATTACK = "MITRE_ATTACK"
	FrameworkOWASP       = "OWASP"

	// Operations Frameworks
	FrameworkDORA = "DORA" // DevOps Research and Assessment
	FrameworkSRE  = "SRE"  // Site Reliability Engineering
)

// NIST CSF Function constants.
const (
	NISTCSFGovern   = "govern"
	NISTCSFIdentify = "identify"
	NISTCSFProtect  = "protect"
	NISTCSFDetect   = "detect"
	NISTCSFRespond  = "respond"
	NISTCSFRecover  = "recover"
)

// FrameworkMapping represents a mapping to an external framework.
type FrameworkMapping struct {
	Framework   string   `json:"framework"`
	Reference   string   `json:"reference,omitempty"`
	Controls    []string `json:"controls,omitempty"`
	Description string   `json:"description,omitempty"`
}

// AllFrameworks returns all framework constants.
func AllFrameworks() []string {
	return []string{
		FrameworkNISTCSF,
		FrameworkNISTCSF2,
		FrameworkNIST80053,
		FrameworkNIST800171,
		FrameworkNISTRMF,
		FrameworkNISTAIRMF,
		FrameworkFedRAMPHigh,
		FrameworkFedRAMPMod,
		FrameworkFedRAMPLow,
		FrameworkISO27001,
		FrameworkSOC2,
		FrameworkCISControls,
		FrameworkPCIDSS,
		FrameworkHIPAA,
		FrameworkGDPR,
		FrameworkMITREATTACK,
		FrameworkOWASP,
		FrameworkDORA,
		FrameworkSRE,
	}
}

// NISTFrameworks returns NIST-related frameworks.
func NISTFrameworks() []string {
	return []string{
		FrameworkNISTCSF,
		FrameworkNISTCSF2,
		FrameworkNIST80053,
		FrameworkNIST800171,
		FrameworkNISTRMF,
		FrameworkNISTAIRMF,
	}
}

// ComplianceFrameworks returns compliance-focused frameworks.
func ComplianceFrameworks() []string {
	return []string{
		FrameworkISO27001,
		FrameworkSOC2,
		FrameworkPCIDSS,
		FrameworkHIPAA,
		FrameworkGDPR,
		FrameworkFedRAMPHigh,
		FrameworkFedRAMPMod,
		FrameworkFedRAMPLow,
	}
}

// SecurityFrameworks returns security-focused frameworks.
func SecurityFrameworks() []string {
	return []string{
		FrameworkNISTCSF,
		FrameworkNISTCSF2,
		FrameworkMITREATTACK,
		FrameworkOWASP,
		FrameworkCISControls,
	}
}

// OperationsFrameworks returns operations-focused frameworks.
func OperationsFrameworks() []string {
	return []string{
		FrameworkDORA,
		FrameworkSRE,
	}
}

// NISTCSFFunctions returns NIST CSF functions in canonical order.
func NISTCSFFunctions() []string {
	return []string{
		NISTCSFGovern,
		NISTCSFIdentify,
		NISTCSFProtect,
		NISTCSFDetect,
		NISTCSFRespond,
		NISTCSFRecover,
	}
}

// FrameworkDisplayName returns a human-readable name for a framework.
func FrameworkDisplayName(framework string) string {
	names := map[string]string{
		FrameworkNISTCSF:     "NIST CSF 1.1",
		FrameworkNISTCSF2:    "NIST CSF 2.0",
		FrameworkNIST80053:   "NIST 800-53",
		FrameworkNIST800171:  "NIST 800-171",
		FrameworkNISTRMF:     "NIST RMF",
		FrameworkNISTAIRMF:   "NIST AI RMF",
		FrameworkFedRAMPHigh: "FedRAMP High",
		FrameworkFedRAMPMod:  "FedRAMP Moderate",
		FrameworkFedRAMPLow:  "FedRAMP Low",
		FrameworkISO27001:    "ISO 27001",
		FrameworkSOC2:        "SOC 2",
		FrameworkCISControls: "CIS Controls",
		FrameworkPCIDSS:      "PCI DSS",
		FrameworkHIPAA:       "HIPAA",
		FrameworkGDPR:        "GDPR",
		FrameworkMITREATTACK: "MITRE ATT&CK",
		FrameworkOWASP:       "OWASP",
		FrameworkDORA:        "DORA",
		FrameworkSRE:         "SRE",
	}
	if name, ok := names[framework]; ok {
		return name
	}
	return framework
}

// ValidFramework checks if a framework value is valid.
func ValidFramework(framework string) bool {
	for _, f := range AllFrameworks() {
		if f == framework {
			return true
		}
	}
	return false
}

// NISTCSFFunctionSortWeight returns a sort weight for NIST CSF functions.
// Returns the canonical order: govern=1, identify=2, protect=3, detect=4, respond=5, recover=6.
func NISTCSFFunctionSortWeight(function string) int {
	weights := map[string]int{
		NISTCSFGovern:   1,
		NISTCSFIdentify: 2,
		NISTCSFProtect:  3,
		NISTCSFDetect:   4,
		NISTCSFRespond:  5,
		NISTCSFRecover:  6,
	}
	if w, ok := weights[function]; ok {
		return w
	}
	return 99
}
