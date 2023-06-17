package mocks

import (
	"github/beomsun1234/stockprice-collector/domain"
	"github/beomsun1234/stockprice-collector/repository"
)

type MockKisAccessTokenRepository struct {
	ResToken *domain.Token
}

func NewMockKisAccessTokenRepository() repository.KisAccessTokenRepositoryInterface {
	return &MockKisAccessTokenRepository{}
}

func (m *MockKisAccessTokenRepository) GetKisAccessToken() (*domain.Token, error) {
	return m.ResToken, nil
}
func (m *MockKisAccessTokenRepository) DeleteKisAccessToken() {

}
func (m *MockKisAccessTokenRepository) InsertKisAccessToken(token *domain.Token) error {
	return nil

}
