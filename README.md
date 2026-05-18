# PRISM Core

Shared primitives and types for the PRISM ecosystem.

[![Go Reference](https://pkg.go.dev/badge/github.com/grokify/prism-core.svg)](https://pkg.go.dev/github.com/grokify/prism-core)
[![Go Report Card](https://goreportcard.com/badge/github.com/grokify/prism-core)](https://goreportcard.com/report/github.com/grokify/prism-core)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Overview

PRISM Core provides foundational types and constants used across the PRISM ecosystem:

- **[PRISM Capability](https://github.com/grokify/prism-capability)** - What we need (capability requirements)
- **[PRISM Intelligence](https://github.com/grokify/prism-intelligence)** - How we measure (metrics and observability)
- **[PRISM Execution](https://github.com/grokify/prism-execution)** - How we act (initiatives and roadmaps)

## Installation

```bash
go get github.com/grokify/prism-core
```

## Packages

### Domain Types (`domain.go`)

Domain constants for categorizing capabilities and metrics:

- `DomainSecurity` - Security and compliance
- `DomainOperations` - Operations and reliability
- `DomainQuality` - Quality and testing
- `DomainPlatform` - Platform and infrastructure
- `DomainAI` - AI and machine learning
- `DomainData` - Data management

### Status Types (`status.go`)

Status enumerations for documents and workflows:

- **Document Status**: `StatusDraft`, `StatusApproved`, `StatusActive`, etc.
- **Health Status**: `HealthGreen`, `HealthYellow`, `HealthRed`
- **Priority**: `PriorityCritical`, `PriorityHigh`, `PriorityMedium`, `PriorityLow`
- **MoSCoW**: `MoSCoWMust`, `MoSCoWShould`, `MoSCoWCould`, `MoSCoWWont`

### Metadata (`metadata.go`)

Common metadata structure for documents:

```go
type Metadata struct {
    ID          string    `json:"id,omitempty"`
    Name        string    `json:"name,omitempty"`
    Title       string    `json:"title,omitempty"`
    Description string    `json:"description,omitempty"`
    Version     string    `json:"version,omitempty"`
    Status      string    `json:"status,omitempty"`
    Domain      string    `json:"domain,omitempty"`
    Owner       string    `json:"owner,omitempty"`
    Team        string    `json:"team,omitempty"`
    Authors     []Person  `json:"authors,omitempty"`
    Tags        []string  `json:"tags,omitempty"`
    CreatedAt   time.Time `json:"createdAt,omitzero"`
    UpdatedAt   time.Time `json:"updatedAt,omitzero"`
}
```

### Person & Team (`person.go`)

People and team structures with Team Topologies support:

```go
type Person struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
    Role  string `json:"role,omitempty"`
}

type Team struct {
    ID      string   `json:"id,omitempty"`
    Name    string   `json:"name"`
    Type    string   `json:"type,omitempty"` // stream_aligned, platform, enabling, overlay
    Members []Person `json:"members,omitempty"`
}
```

### Layers & Stages (`layer.go`)

Value stream layers and lifecycle stages:

**Layers:**

- `LayerRequirements` - Product ideation, specs, design
- `LayerCode` - Application code, libraries, dependencies
- `LayerInfra` - Cloud resources, networking, platform
- `LayerRuntime` - Running services, containers, workloads
- `LayerAdoption` - Product analytics, user engagement
- `LayerSupport` - Customer support, incident management

**Stages:**

- `StageDesign` - Architecture, requirements, planning
- `StageBuild` - CI/CD, code quality, dependency management
- `StageTest` - Testing coverage, quality assurance
- `StageRuntime` - Production monitoring, availability
- `StageResponse` - Incident response, remediation

### Maturity Levels (`maturity.go`)

5-level maturity model (M1-M5):

| Level | Name | Description |
|-------|------|-------------|
| M1 | Reactive | Ad-hoc processes, firefighting mode |
| M2 | Basic | Basic controls, individual-dependent |
| M3 | Defined | Standardized, repeatable processes |
| M4 | Managed | Data-driven, measured and controlled |
| M5 | Optimizing | Continuous improvement, automated |

### Frameworks (`framework.go`)

Compliance and security framework constants:

- **NIST**: CSF, CSF 2.0, 800-53, 800-171, RMF, AI RMF
- **FedRAMP**: High, Moderate, Low
- **Industry**: ISO 27001, SOC 2, PCI DSS, HIPAA, GDPR
- **Security**: MITRE ATT&CK, OWASP, CIS Controls
- **Operations**: DORA, SRE

### Risk Types (`risk.go`)

Risk assessment primitives:

```go
type Risk struct {
    ID          string `json:"id,omitempty"`
    Description string `json:"description"`
    Impact      string `json:"impact,omitempty"`      // low, medium, high, critical
    Probability string `json:"probability,omitempty"` // low, medium, high
    Mitigation  string `json:"mitigation,omitempty"`
    Status      string `json:"status,omitempty"`      // open, mitigated, accepted, closed
    Owner       string `json:"owner,omitempty"`
}
```

### SLI Types (`sli.go`)

Service Level Indicator types:

- `SLITypeAvailability` - Service uptime
- `SLITypeLatency` - Response time
- `SLITypeThroughput` - Requests per second
- `SLITypeErrorRate` - Error percentage
- `SLITypeFreshness` - Data staleness
- `SLITypeCorrectness` - Result accuracy
- `SLITypeCoverage` - Response coverage
- `SLITypeDurability` - Data persistence

### Trend Types (`trend.go`)

Trend direction and analysis:

```go
type TrendInfo struct {
    Direction  string  `json:"direction"`  // up, down, flat, unknown
    Magnitude  float64 `json:"magnitude,omitempty"`
    Percentage float64 `json:"percentage,omitempty"`
}
```

## PRISM Ecosystem

```
┌─────────────────────────────────────────────────────────────────┐
│                         PRISM Core                              │
│              (Shared Primitives & Types)                        │
└─────────────────────────────────────────────────────────────────┘
                              │
          ┌───────────────────┼───────────────────┐
          │                   │                   │
          ▼                   ▼                   ▼
┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
│ PRISM Capability│ │PRISM Intelligence│ │ PRISM Execution │
│                 │ │                 │ │                 │
│ "What we need"  │ │"How we measure" │ │  "How we act"   │
│                 │ │                 │ │                 │
│ - Capabilities  │ │ - Metrics       │ │ - Initiatives   │
│ - Requirements  │ │ - SLIs/SLOs     │ │ - Roadmaps      │
│ - Gaps          │ │ - Assessments   │ │ - Tasks         │
└─────────────────┘ └─────────────────┘ └─────────────────┘
```

## License

MIT License - see [LICENSE](LICENSE) for details.
