package main

type ParsingService struct {
	strategy ParsingStrategy
}

func NewParsingService(strategy ParsingStrategy) *ParsingService {
	return &ParsingService{strategy: strategy}
}

func (s *ParsingService) SetStrategy(strategy ParsingStrategy) {
	s.strategy = strategy
}

func (s *ParsingService) ParseTrack(data []byte) (Track, error) {
	return s.strategy.Parse(data)
}
