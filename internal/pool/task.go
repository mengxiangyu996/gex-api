package pool

type Task interface {
	Do() error
}
