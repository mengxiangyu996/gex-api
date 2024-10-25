package pool

import "sync"

var (
	once sync.Once
	pool *Pool
)

type Task struct {
	Func func() error
}

// 协程池
type Pool struct {
	EntryChan chan *Task    // 任务入口，接收任务的作用
	workerNum int           // 协程池最大工作数量
	taskChan  chan *Task    // 任务队列
	stop      chan struct{} // 停止信号
	sync.WaitGroup
}

// 初始化异步协程池
func GetInstancePool(workerNum int) *Pool {

	once.Do(func() {
		pool = &Pool{
			EntryChan: make(chan *Task),
			workerNum: workerNum,
			taskChan:  make(chan *Task, workerNum),
			stop:      make(chan struct{}),
		}
	})

	return pool
}
