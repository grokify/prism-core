package core

import "time"

// Metadata contains common document metadata fields.
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

// NewMetadata creates a new Metadata with the given name and author.
func NewMetadata(name string, author Person) *Metadata {
	now := time.Now()
	return &Metadata{
		Name:      name,
		Status:    StatusDraft,
		Authors:   []Person{author},
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Touch updates the UpdatedAt timestamp.
func (m *Metadata) Touch() {
	m.UpdatedAt = time.Now()
}

// AddAuthor adds an author if not already present.
func (m *Metadata) AddAuthor(author Person) {
	for _, a := range m.Authors {
		if a.Email == author.Email {
			return
		}
	}
	m.Authors = append(m.Authors, author)
}

// HasTag checks if the metadata has a specific tag.
func (m *Metadata) HasTag(tag string) bool {
	for _, t := range m.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// AddTag adds a tag if not already present.
func (m *Metadata) AddTag(tag string) {
	if !m.HasTag(tag) {
		m.Tags = append(m.Tags, tag)
	}
}
