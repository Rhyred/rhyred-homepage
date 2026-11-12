package models

// Project represents a portfolio project.
type Project struct {
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Status      string   `json:"status"`
	Featured    bool     `json:"featured"`
}

// Profile represents the site owner's profile.
type Profile struct {
	Name        string         `json:"name"`
	Tagline     string         `json:"tagline"`
	Description string         `json:"description"`
	Roles       []string       `json:"roles"`
	About       AboutSection   `json:"about"`
	Contact     Contact        `json:"contact"`
	TechStack   TechStack      `json:"tech_stack"`
	LabAreas    []LabArea      `json:"lab_areas"`
	Journal     []JournalEntry `json:"journal_entries"`
}

// AboutSection holds about page content.
type AboutSection struct {
	Headline     string `json:"headline"`
	Summary      string `json:"summary"`
	CurrentFocus string `json:"current_focus"`
	Background   string `json:"background"`
}

// Contact holds contact information.
type Contact struct {
	GitHub string            `json:"github"`
	Email  string            `json:"email"`
	Social map[string]string `json:"social"`
}

// TechStack categorizes technologies.
type TechStack struct {
	ProgrammingLanguages []string `json:"programming_languages"`
	Backend              []string `json:"backend"`
	Frontend             []string `json:"frontend"`
	AI                   []string `json:"ai"`
	Networking           []string `json:"networking"`
	Tools                []string `json:"tools"`
}

// LabArea represents a digital lab exploration area.
type LabArea struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// JournalEntry represents a journal/blog entry.
type JournalEntry struct {
	Slug     string `json:"slug"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Category string `json:"category"`
	Summary  string `json:"summary"`
	IsSample bool   `json:"is_sample"`
}
