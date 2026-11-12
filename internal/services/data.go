package services

import (
	"encoding/json"
	"os"
	"strings"
	"sync"

	"github.com/Rhyred/rhyred-homepage/internal/models"
)

// DataService loads and provides access to JSON data.
type DataService struct {
	dataDir  string
	projects []models.Project
	profile  models.Profile
	mu       sync.RWMutex
	loaded   bool
}

// NewDataService creates a new DataService.
func NewDataService(dataDir string) *DataService {
	return &DataService{dataDir: dataDir}
}

// Load reads data files from disk.
func (ds *DataService) Load() error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if err := ds.loadProjects(); err != nil {
		return err
	}
	if err := ds.loadProfile(); err != nil {
		return err
	}

	ds.loaded = true
	return nil
}

func (ds *DataService) loadProjects() error {
	data, err := os.ReadFile(ds.dataDir + "/projects.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &ds.projects)
}

func (ds *DataService) loadProfile() error {
	data, err := os.ReadFile(ds.dataDir + "/profile.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &ds.profile)
}

// GetProjects returns all projects.
func (ds *DataService) GetProjects() []models.Project {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	result := make([]models.Project, len(ds.projects))
	copy(result, ds.projects)
	return result
}

// GetFeaturedProjects returns featured projects.
func (ds *DataService) GetFeaturedProjects() []models.Project {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	var result []models.Project
	for _, p := range ds.projects {
		if p.Featured {
			result = append(result, p)
		}
	}
	return result
}

// GetProjectBySlug returns a single project by slug.
func (ds *DataService) GetProjectBySlug(slug string) (models.Project, bool) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	for _, p := range ds.projects {
		if p.Slug == slug {
			return p, true
		}
	}
	return models.Project{}, false
}

// GetProjectsByCategory returns projects matching a category.
func (ds *DataService) GetProjectsByCategory(category string) []models.Project {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	if category == "" || category == "all" {
		result := make([]models.Project, len(ds.projects))
		copy(result, ds.projects)
		return result
	}
	var result []models.Project
	for _, p := range ds.projects {
		if strings.EqualFold(p.Category, category) {
			result = append(result, p)
		}
	}
	return result
}

// GetProfile returns the profile data.
func (ds *DataService) GetProfile() models.Profile {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.profile
}

// GetCategories returns unique project categories.
func (ds *DataService) GetCategories() []string {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	seen := make(map[string]bool)
	var categories []string
	for _, p := range ds.projects {
		if !seen[p.Category] {
			seen[p.Category] = true
			categories = append(categories, p.Category)
		}
	}
	return categories
}

