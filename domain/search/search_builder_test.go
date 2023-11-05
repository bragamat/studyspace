package search_test

import (
	"testing"

	"github.com/bragamat/studyspace/domain/search"
)

type testExpectations struct {
	Name         string
	IsValid      bool
	SearchBuiler search.Search
}

func TestIsValid(t *testing.T) {
	mockSearches := []testExpectations{
		{
			Name:    "Language must be valid",
			IsValid: false,
			SearchBuiler: search.Search{
				Owner: "fakeUser",
			},
		},
	}

	for _, tt := range mockSearches {
		t.Run(tt.Name, func(t *testing.T) {
			IsValid, _ := tt.SearchBuiler.IsValid()

			if IsValid != tt.IsValid {
				t.Error("Search Builder is Valid and it SHOULD NOT be")
			}
		})
	}
}
