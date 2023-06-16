package scheduler

type Scheduler interface {
	AddFunc(spec string, cmd func()) error
	Start()
	Stop()
}
