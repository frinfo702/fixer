package usecase

import "github.com/frinfo702/rustysearch/domain/model"

type Searcher interface {
	Search(query model.SearchQuery) ([]model.SearchResult, error)
}
