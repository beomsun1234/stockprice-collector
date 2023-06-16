package mocks

type MockScheduler struct {
	task func()
}

func (m *MockScheduler) AddFunc(spec string, cmd func()) error {
	m.task = cmd
	return nil
}

func (m *MockScheduler) Start() {
	// Do nothing in mock implementation
}

func (m *MockScheduler) Stop() {
	// Do nothing in mock implementation
}
