package search_test

import (
	"testing"

	"github.com/bragamat/studyspace/domain/search"
	"github.com/bragamat/studyspace/domain/stack"
)

type testExpectations struct {
	Name         string
	IsValid      bool
	SearchBuiler *search.Search
}

func TestIsValid(t *testing.T) {
	mockSearches := []testExpectations{
		{
			Name:    "Language must be valid",
			IsValid: false,
			SearchBuiler: &search.Search{
				Owner: "fakeUser",
			},
		},
		{
			Name:    "Valid",
			IsValid: true,
			SearchBuiler: &search.Search{
				Owner:    "fakeUser",
				Language: stack.Golang,
			},
		},
	}

	for _, tt := range mockSearches {
		t.Run(tt.Name, func(t *testing.T) {
			IsValid, _ := tt.SearchBuiler.IsValid()

			if IsValid != tt.IsValid {
				t.Errorf("Search Builder validation result is %v and it must be %v", IsValid, tt.IsValid)
			}
		})
	}
}
