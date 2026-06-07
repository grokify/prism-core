# Release Notes v0.2.0

This release removes Priority and MoSCoW types from prism-core. These are now available in the dedicated [priority-frameworks](https://github.com/grokify/priority-frameworks) library.

## Breaking Changes

The following constants and functions have been removed:

### Priority Constants
- `PriorityCritical`
- `PriorityHigh`
- `PriorityMedium`
- `PriorityLow`

### Priority Functions
- `ValidPriority()`
- `PriorityWeight()`

### MoSCoW Constants
- `MoSCoWMust`
- `MoSCoWShould`
- `MoSCoWCould`
- `MoSCoWWont`

### MoSCoW Functions
- `ValidMoSCoW()`
- `MoSCoWWeight()`

## Migration Guide

Replace prism-core priority types with priority-frameworks:

```go
import pf "github.com/grokify/priority-frameworks"

// Before: prism.PriorityCritical
// After:
level := pf.Severity().Parse("critical")

// Before: prism.ValidPriority(p)
// After:
valid := pf.Severity().Parse(p) != nil

// Before: prism.PriorityWeight(p)
// After:
weight := pf.Normalize(pf.Severity(), p)

// Before: prism.MoSCoWMust
// After:
level := pf.MoSCoW().Parse("must")
```

## Why This Change?

priority-frameworks provides:

- **More frameworks**: Severity, Priority (P#), IETF RFC 2119, MoSCoW, General
- **Cross-framework normalization**: Compare priorities across different systems
- **Score mapping**: Convert CVSS scores to severity levels
- **Level counting**: Track counts per level for dashboards

prism-core now focuses on domain-specific primitives (domains, layers, stages, maturity) while priority-frameworks handles generic prioritization.

## Links

- [priority-frameworks](https://github.com/grokify/priority-frameworks)
- [Changelog](https://github.com/grokify/prism-core/blob/main/CHANGELOG.md)
