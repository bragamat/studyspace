package search

import (
	"errors"

	"github.com/bragamat/studyspace/domain/repo"
	"github.com/bragamat/studyspace/domain/stack"
)

const UrlSearch = "search?q=owner:bragamat&language:Go&type=repositories"

type Search struct {
	Owner      string
	Language   stack.Stack
	ResultType repo.Repository
	Url        string
}

func (s *Search) IsValid() (bool, []error) {
	var allErrors []error

	if s.Language == "" {
		allErrors = append(allErrors, errors.New("Language must be Valid"))
	}

	if s.Url == "" {
		allErrors = append(allErrors, errors.New("URL must be Valid"))
	}

	if s.Owner == "" {
		allErrors = append(allErrors, errors.New("Owner must be Valid"))
	}

	if len(allErrors) > 0 {
		return false, allErrors
	}

	return true, nil
}
