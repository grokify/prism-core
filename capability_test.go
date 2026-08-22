package core

import "testing"

func TestValidCapabilityRelation(t *testing.T) {
	tests := []struct {
		relation string
		want     bool
	}{
		{CapabilityRelationEnables, true},
		{CapabilityRelationImproves, true},
		{CapabilityRelationDependsOn, true},
		{"", false},
		{"blocks", false},
	}
	for _, tt := range tests {
		if got := ValidCapabilityRelation(tt.relation); got != tt.want {
			t.Errorf("ValidCapabilityRelation(%q) = %v, want %v", tt.relation, got, tt.want)
		}
	}
}

func TestAllCapabilityRelationsAreValid(t *testing.T) {
	for _, r := range AllCapabilityRelations() {
		if !ValidCapabilityRelation(r) {
			t.Errorf("AllCapabilityRelations() returned %q, which ValidCapabilityRelation rejects", r)
		}
	}
}

func TestCapabilityRelationDisplayName(t *testing.T) {
	if got := CapabilityRelationDisplayName(CapabilityRelationDependsOn); got != "Depends On" {
		t.Errorf("CapabilityRelationDisplayName(dependsOn) = %q, want %q", got, "Depends On")
	}
	if got := CapabilityRelationDisplayName("unknown"); got != "unknown" {
		t.Errorf("CapabilityRelationDisplayName(unknown) = %q, want the input echoed back", got)
	}
}
