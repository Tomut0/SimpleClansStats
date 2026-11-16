package application

import (
	"context"
	"unicode/utf8"

	"github.com/Tomut0/SimpleClansStats/internal/domain"
	"github.com/Tomut0/SimpleClansStats/internal/domain/ports"
)

type SearchService struct{ repository ports.SearchRepository }

func NewSearchService(repository ports.SearchRepository) *SearchService {
	return &SearchService{repository: repository}
}

func (s *SearchService) Search(ctx context.Context, serverID, query string, limit int) ([]domain.SearchResult, error) {
	if utf8.RuneCountInString(query) > 100 {
		query = string([]rune(query)[:100])
	}
	if limit < 1 || limit > 50 {
		limit = 10
	}
	data, err := s.repository.Search(ctx, serverID, query, limit)
	if err != nil {
		return nil, err
	}
	return domain.ResolveSearchResults(data.Matches, data.Clans, data.Players)
}
