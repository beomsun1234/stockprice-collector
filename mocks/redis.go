package mocks

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

type MockRedisDatabase struct {
	Redis     *redis.Client
	RedisMock redismock.ClientMock
}

func NewMockRedisDatabase() *MockRedisDatabase {
	return &MockRedisDatabase{}
}

func (m *MockRedisDatabase) Connect() {
	m.Redis, m.RedisMock = redismock.NewClientMock()
}

func (m *MockRedisDatabase) Client() *redis.Client {
	return m.Redis
}
