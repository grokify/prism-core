# Product Lines

A product line is a top-level product concept, a peer of the initiative.
Roadmap items and execution work both point at it, so `ProductLine` lives in
prism-core, where planning tools and execution tools such as visionstudio can
share one definition.

```go
import prism "github.com/grokify/prism-core"

line := prism.ProductLine{
	ID:             "analytics",
	Name:           "Analytics",
	Kind:           prism.ProductLineKindVertical,
	ParentID:       "data",
	MarketPosition: prism.MarketPositionStar,
}
if err := line.Validate(); err != nil {
	return err
}
```

As JSON:

```json
{
  "id": "analytics",
  "name": "Analytics",
  "kind": "vertical",
  "parentId": "data",
  "marketPosition": "star"
}
```

## Fields

| Field | JSON | Required | Meaning |
|-------|------|----------|---------|
| `ID` | `id` | yes | A canonical, stable identifier |
| `Name` | `name` | yes | The name people see |
| `Kind` | `kind` | yes | Where the line sits on the portfolio axis |
| `ParentID` | `parentId` | no | A higher grouping, such as a product group or business unit |
| `MarketPosition` | `marketPosition` | no | The line's growth-share classification; empty means unclassified |

Source tools name this concept differently: a "product" in Aha!, a "component"
in Productboard, a "project" in Jira Product Discovery. An integration such as
[omniroadmap](https://github.com/grokify/omniroadmap) maps each to `ID`, `Name`,
and `Kind`.

`ParentID` and `MarketPosition` are strategic judgments that a person sets. A
sync from a source tool never recomputes or overwrites them.

## Kinds

| Kind | Constant | Meaning |
|------|----------|---------|
| `vertical` | `ProductLineKindVertical` | A line with its own market position, such as Analytics, Payments, or Messaging |
| `platform` | `ProductLineKindPlatform` | A shared enabler, such as Platform: leverage across lines, not its own market share |
| `market` | `ProductLineKindMarket` | A go-to-market category, such as Partner Channels. It is not a product line and has no market position |

Some teams call a platform line "horizontal". prism-core says "platform" so the
word does not clash with a cross-line initiative.

## Market positions

A market position classifies a line by market growth and market share. The
quadrants are inspired by the Boston Consulting Group's growth-share matrix,
with an enabler position added for a shared platform.

| Position | Constant | Growth | Share | Usual stance |
|----------|----------|--------|-------|--------------|
| `star` | `MarketPositionStar` | high | high | invest |
| `cash_cow` | `MarketPositionCashCow` | low | high | harvest |
| `question_mark` | `MarketPositionQuestionMark` | high | low | selective bet |
| `dog` | `MarketPositionDog` | low | low | divest |
| `enabler` | `MarketPositionEnabler` | — | — | sustain for leverage |

Assign the position at the level that has a distinct market position. Usually
that is the line itself, but it can be set on a parent group or business unit
and inherited by the lines below it.

The position is chosen by hand. To back the choice with evidence, use
prism-roadmap's `MarketPositionDimension`, a rubric whose options use the same
IDs.

## Validation

`Validate` returns the first problem it finds:

- a missing `ID` or `Name`
- a `Kind` that is empty or unknown
- a `MarketPosition` that is not one of the positions above
- a market position on a `market` category

`ValidProductLineKind` and `ValidMarketPosition` check a single value, and
`AllProductLineKinds` and `AllMarketPositions` list them, for building pickers
or schema enums.

## Scope of an initiative

An initiative can reference one product line, several, or all of them. The
scope is derived from those references and not stored:

```go
scope := prism.ScopeFor(len(initiative.ProductLineIDs), len(portfolio.Lines))
```

| Scope | Constant | When |
|-------|----------|------|
| unknown | `ScopeUnknown` | No line referenced. This is the zero value, so an unset scope never reads as an answer |
| line-specific | `ScopeLineSpecific` | Exactly one line |
| cross-line | `ScopeCrossLine` | A named subset of lines |
| all lines | `ScopeAllLines` | Every line, such as accessibility, localization, reliability, or security work |
