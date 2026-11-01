package core

import (
	"errors"
	"fmt"
)

// ProductLine is the canonical, shared product-line entity — a top-level product
// concept (peer of Initiative), referenced by roadmap items and execution work
// alike. It lives in prism-core because both planning tools and execution tools
// such as visionstudio reference it; omniroadmap-core carries only the ID refs.
//
// Source tools name this concept differently (Aha! "product", Productboard
// "component", Jira Product Discovery "project"); an integration such as
// omniroadmap maps each to ID, Name, and Kind.
//
// ParentID and MarketPosition are hand-curated strategic judgments: a person
// sets them, and a sync from a source tool never recomputes or overwrites them.
type ProductLine struct {
	// ID is a canonical, stable identifier (e.g. "analytics").
	ID string `json:"id"`
	// Name is the human-facing name (e.g. "Analytics").
	Name string `json:"name"`
	// Kind is the portfolio-axis classification.
	Kind ProductLineKind `json:"kind"`
	// ParentID is the ID of a higher product-line grouping (product group or
	// business unit), or "" for a top-level line. Lines form a hierarchy so a
	// market position can be assigned at whichever level owns it.
	ParentID string `json:"parentId,omitempty"`
	// MarketPosition is the growth-share classification. It is assigned at the
	// level that has a distinct market position — usually this line, but it may
	// be set on a parent group or business unit and inherited down. "" means
	// unclassified.
	MarketPosition MarketPosition `json:"marketPosition,omitempty"`
}

// Validate reports the first problem with the product line: a missing ID or
// name, an unknown kind or market position, or a market position on a market
// category, which is not a product line and so has none.
func (p ProductLine) Validate() error {
	if p.ID == "" {
		return errors.New("product line: id is required")
	}
	if p.Name == "" {
		return fmt.Errorf("product line %q: name is required", p.ID)
	}
	if !ValidProductLineKind(p.Kind) {
		return fmt.Errorf("product line %q: unknown kind %q", p.ID, p.Kind)
	}
	if p.MarketPosition != "" && !ValidMarketPosition(p.MarketPosition) {
		return fmt.Errorf("product line %q: unknown market position %q", p.ID, p.MarketPosition)
	}
	if p.Kind == ProductLineKindMarket && p.MarketPosition != "" {
		return fmt.Errorf("product line %q: a market category has no market position", p.ID)
	}
	return nil
}

// ProductLineKind classifies a product line on the portfolio axis (distinct from
// the Initiative > Epic > Feature hierarchy).
type ProductLineKind string

const (
	// ProductLineKindVertical is a line with its own market position (e.g.
	// Analytics, Payments, Messaging).
	ProductLineKindVertical ProductLineKind = "vertical"
	// ProductLineKindPlatform is a shared enabler ("Platform"): leverage across
	// lines, not its own market share. Its market position is
	// MarketPositionEnabler. (Some teams say "horizontal"; we say "platform" to
	// avoid clashing with a cross-line initiative.)
	ProductLineKindPlatform ProductLineKind = "platform"
	// ProductLineKindMarket is a go-to-market category ("Partner Channels"). It
	// is not a product line and has no market position.
	ProductLineKindMarket ProductLineKind = "market"
)

// AllProductLineKinds returns all product line kinds.
func AllProductLineKinds() []ProductLineKind {
	return []ProductLineKind{ProductLineKindVertical, ProductLineKindPlatform, ProductLineKindMarket}
}

// ValidProductLineKind reports whether k is a known product line kind.
func ValidProductLineKind(k ProductLineKind) bool {
	switch k {
	case ProductLineKindVertical, ProductLineKindPlatform, ProductLineKindMarket:
		return true
	default:
		return false
	}
}

// MarketPosition classifies a product line (or a group or business unit it
// rolls up to) by market growth and market share. The quadrants are inspired by
// the Boston Consulting Group's growth-share matrix, with Enabler added for a
// shared platform. It is hand-curated; prism-roadmap's MarketPositionDimension
// is an optional evidence-backed rubric for choosing it.
type MarketPosition string

const (
	MarketPositionStar         MarketPosition = "star"          // high growth, high share → invest
	MarketPositionCashCow      MarketPosition = "cash_cow"      // low growth, high share → harvest
	MarketPositionQuestionMark MarketPosition = "question_mark" // high growth, low share → selective bet
	MarketPositionDog          MarketPosition = "dog"           // low growth, low share → divest
	MarketPositionEnabler      MarketPosition = "enabler"       // shared platform → sustain for leverage
)

// AllMarketPositions returns all market positions.
func AllMarketPositions() []MarketPosition {
	return []MarketPosition{
		MarketPositionStar,
		MarketPositionCashCow,
		MarketPositionQuestionMark,
		MarketPositionDog,
		MarketPositionEnabler,
	}
}

// ValidMarketPosition reports whether m is a known market position.
func ValidMarketPosition(m MarketPosition) bool {
	switch m {
	case MarketPositionStar, MarketPositionCashCow, MarketPositionQuestionMark,
		MarketPositionDog, MarketPositionEnabler:
		return true
	default:
		return false
	}
}

// ProductLineScope is derived from how many lines an initiative references, not
// stored: an initiative may be line-specific (one) or cross-line/"horizontal"
// (many, or all — e.g. accessibility, localization, reliability, security).
// ScopeFor derives it. The zero value is ScopeUnknown, so an unset scope never
// reads as a real answer.
type ProductLineScope int

const (
	ScopeUnknown      ProductLineScope = iota // no line referenced
	ScopeLineSpecific                         // exactly one line
	ScopeCrossLine                            // a named subset of lines
	ScopeAllLines                             // portfolio-wide
)

// ScopeFor derives the scope of an initiative that references `referenced` of
// the portfolio's `total` product lines.
func ScopeFor(referenced, total int) ProductLineScope {
	switch {
	case referenced <= 0:
		return ScopeUnknown
	case referenced == 1:
		return ScopeLineSpecific
	case referenced >= total:
		return ScopeAllLines
	default:
		return ScopeCrossLine
	}
}
