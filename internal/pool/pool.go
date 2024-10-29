package pool

import "sync"

type Pool struct {
	Size      int       // 协程池最大worker数量，即goroutine数量
	TaskQueue chan Task // 任务队列
	workers   chan *Worker
	wg        sync.WaitGroup
}

// 初始化协程池
func NewPool(size int) *Pool {
	return &Pool{
		Size:      size,
		TaskQueue: make(chan Task),
		workers:   make(chan *Worker, size),
		wg:        sync.WaitGroup{},
	}
}

// 启动协程池
func (p *Pool) Start() {
	// 启动协程，并且加入池子中
	for i := 0; i < p.Size; i++ {
		worker := NewWorker()
		go worker.Start(p)
	}

	// 分配任务
	go func() {
		for task := range p.TaskQueue {
			p.wg.Add(1)
			worker := <-p.workers
			worker.taskQueue <- task
		}
	}()
}

// 添加任务
func (p *Pool) AddTask(task Task) {
	p.TaskQueue <- task
}

// 停止协程池
func (p *Pool) Stop() {
	p.wg.Wait()        // 等待所有任务完成
	close(p.TaskQueue) // 关闭任务队列
	close(p.workers)   // 关闭工作者通道
}
