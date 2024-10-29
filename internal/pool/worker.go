package pool

type Worker struct {
	taskQueue chan Task
	stop      chan struct{}
}

// 初始化协程
func NewWorker() *Worker {
	return &Worker{
		taskQueue: make(chan Task),
		stop:      make(chan struct{}),
	}
}

// 启动协程
func (w *Worker) Start(pool *Pool) {
	// 防止阻塞，开启协程处理
	go func() {
		for {
			// 注册协程加入到协程池
			pool.workers <- w
			select {
			case <-w.stop: // 如果接收到停止信号，退出循环
				return
			case task := <-w.taskQueue: // 从任务队列中取出任务
				task.Do() // 执行任务
				pool.wg.Done()
			}
		}
	}()
}

// 停止协程
func (w *Worker) Stop() {
	close(w.stop)
}
