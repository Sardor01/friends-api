package domain

import "strings"

// validIncludes
var validIncludes = []string{
	"firstAppearance",
	"lastAppearance",
	"numberOfEpisodes",
	"portrayedBy",
	"name",
	"nicknames",
	"gender",
	"dateOfBirth",
	"occupations",
	"relatives",
	"lines",
}

// Person struct
type Person struct {
	FirstAppearance  string          `json:"firstAppearance,omitempty" bson:"firstAppearance"`
	LastAppearance   string          `json:"lastAppearance,omitempty" bson:"lastAppearance"`
	NumberOfEpisodes int32           `json:"numberOfEpisodes,omitempty" bson:"numberOfEpisodes"`
	PotrayedBy       []string        `json:"portrayedBy,omitempty" bson:"portrayedBy"`
	Name             string          `json:"name,omitempty" bson:"name"`
	Nicknames        []string        `json:"nicknames,omitempty" bson:"nicknames"`
	Gender           string          `json:"gender,omitempty" bson:"gender"`
	DateOfBirth      string          `json:"dateOfBirth,omitempty" bson:"dateOfBirth"`
	Occupations      []string        `json:"occupations,omitempty" bson:"occupations"`
	Relatives        []Relative      `json:"relatives,omitempty" bson:"relatives"`
	Lines            []CharacterLine `json:"lines,omitempty" bson:"lines"`
}

// Relative struct
type Relative struct {
	Name         string `json:"name,omitempty" bson:"name"`
	Relationship string `json:"relationship,omitempty" bson:"relationship"`
}

// CharacterLine struct
type CharacterLine struct {
	Said          string `json:"said,omitempty" bson:"said"`
	SeasonNumber  int    `json:"seasonNumber,omitempty" bson:"seasonNumber"`
	EpisodeNumber int    `json:"episodeNumber,omitempty" bson:"episodeNumber"`
}

// PersonQuery struct
type PersonQuery struct {
	Find    map[string]interface{}
	Limit   int
	OrderBy []string
	Include map[string]interface{}
}

// AddLimit function
func (pq *PersonQuery) AddLimit(limitQuery string) bool {
	limit, ok := isLimitParamValid(limitQuery)
	pq.Limit = limit
	return ok
}

// AddOrderBy function
func (pq *PersonQuery) AddOrderBy(orderByQuery string) {
	orderBy := isOrderByParamValid(strings.Split(orderByQuery, ","),
		[]string{"-numberOfEpisodes", "numberOfEpisodes"})
	pq.OrderBy = orderBy
}

// AddFind function
func (pq *PersonQuery) AddFind(key string, value interface{}) {
	if pq.Find == nil {
		pq.Find = map[string]interface{}{}
	}
	pq.Find[key] = value
}

// AddInclude function
func (pq *PersonQuery) AddInclude(includeQuery string) {
	include := map[string]interface{}{
		"firstAppearance":  1,
		"lastAppearance":   1,
		"numberOfEpisodes": 1,
		"name":             1,
	}
	pq.Include = isIncludeParamValid(strings.Split(includeQuery, ","), validIncludes, include)
}
