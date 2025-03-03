package base

// Scheduler is an interface for a scheduled task
type Scheduler interface {
	RunTask() *Error
}
