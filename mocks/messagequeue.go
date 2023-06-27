package mocks

import "github/beomsun1234/stockprice-collector/infra/messagequeue"

type MockMessagequeue struct {
}

func NewMockMessagequeue() messagequeue.Messagequeue {
	return &MockMessagequeue{}
}

func (m *MockMessagequeue) SendMessage(msg []byte) error {
	return nil
}
