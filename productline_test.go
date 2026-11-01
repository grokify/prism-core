package core

import (
	"encoding/json"
	"testing"
)

func TestProductLineValidate(t *testing.T) {
	tests := []struct {
		name    string
		line    ProductLine
		wantErr bool
	}{
		{"vertical with position", ProductLine{ID: "analytics", Name: "Analytics", Kind: ProductLineKindVertical, MarketPosition: MarketPositionStar}, false},
		{"vertical unclassified", ProductLine{ID: "payments", Name: "Payments", Kind: ProductLineKindVertical}, false},
		{"platform enabler", ProductLine{ID: "platform", Name: "Platform", Kind: ProductLineKindPlatform, MarketPosition: MarketPositionEnabler}, false},
		{"market category", ProductLine{ID: "partner-channels", Name: "Partner Channels", Kind: ProductLineKindMarket}, false},
		{"missing id", ProductLine{Name: "Analytics", Kind: ProductLineKindVertical}, true},
		{"missing name", ProductLine{ID: "analytics", Kind: ProductLineKindVertical}, true},
		{"missing kind", ProductLine{ID: "analytics", Name: "Analytics"}, true},
		{"unknown kind", ProductLine{ID: "analytics", Name: "Analytics", Kind: "horizontal"}, true},
		{"unknown position", ProductLine{ID: "analytics", Name: "Analytics", Kind: ProductLineKindVertical, MarketPosition: "unicorn"}, true},
		{"market category with position", ProductLine{ID: "partner-channels", Name: "Partner Channels", Kind: ProductLineKindMarket, MarketPosition: MarketPositionDog}, true},
	}
	for _, tt := range tests {
		err := tt.line.Validate()
		if (err != nil) != tt.wantErr {
			t.Errorf("%s: Validate() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestAllProductLineKindsAreValid(t *testing.T) {
	for _, k := range AllProductLineKinds() {
		if !ValidProductLineKind(k) {
			t.Errorf("AllProductLineKinds() returned %q, which ValidProductLineKind rejects", k)
		}
	}
	if ValidProductLineKind("") {
		t.Error(`ValidProductLineKind("") = true, want false`)
	}
}

func TestAllMarketPositionsAreValid(t *testing.T) {
	for _, m := range AllMarketPositions() {
		if !ValidMarketPosition(m) {
			t.Errorf("AllMarketPositions() returned %q, which ValidMarketPosition rejects", m)
		}
	}
	if ValidMarketPosition("") {
		t.Error(`ValidMarketPosition("") = true, want false`)
	}
}

func TestScopeFor(t *testing.T) {
	tests := []struct {
		referenced, total int
		want              ProductLineScope
	}{
		{0, 5, ScopeUnknown},
		{1, 5, ScopeLineSpecific},
		{1, 1, ScopeLineSpecific},
		{3, 5, ScopeCrossLine},
		{5, 5, ScopeAllLines},
	}
	for _, tt := range tests {
		if got := ScopeFor(tt.referenced, tt.total); got != tt.want {
			t.Errorf("ScopeFor(%d, %d) = %v, want %v", tt.referenced, tt.total, got, tt.want)
		}
	}
	var zero ProductLineScope
	if zero != ScopeUnknown {
		t.Errorf("zero ProductLineScope = %v, want ScopeUnknown", zero)
	}
}

func TestProductLineJSON(t *testing.T) {
	line := ProductLine{ID: "analytics", Name: "Analytics", Kind: ProductLineKindVertical, ParentID: "data", MarketPosition: MarketPositionCashCow}
	b, err := json.Marshal(line)
	if err != nil {
		t.Fatal(err)
	}
	want := `{"id":"analytics","name":"Analytics","kind":"vertical","parentId":"data","marketPosition":"cash_cow"}`
	if string(b) != want {
		t.Errorf("json = %s, want %s", b, want)
	}
	b, err = json.Marshal(ProductLine{ID: "payments", Name: "Payments", Kind: ProductLineKindVertical})
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"id":"payments","name":"Payments","kind":"vertical"}`; string(b) != want {
		t.Errorf("json without optional fields = %s, want %s", b, want)
	}
}
