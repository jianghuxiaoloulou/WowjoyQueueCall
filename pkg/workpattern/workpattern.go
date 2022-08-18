package workpattern

// 任务
type Job interface {
	Do()
}

// worker 工人
type Worker struct {
	// 任务队列
	JobQueue chan Job
	// 停止当前任务
	Quit chan bool
}

// 新建一个worker 通道实例 新建一个工人
func NewWorker() Worker {
	return Worker{
		// 初始化工作队列为null
		JobQueue: make(chan Job),
		Quit:     make(chan bool),
	}
}

/*
整个过程中 每个Worker(工人)都会被运行在一个协程中，
在整个WorkerPool(领导)中就会有num个可空闲的Worker(工人)，
当来一条数据的时候，领导就会小组中取一个空闲的Worker(工人)去执行该Job，
当工作池中没有可用的worker(工人)时，就会阻塞等待一个空闲的worker(工人)。
每读到一个通道参数 运行一个 worker
*/
func (w Worker) Run(wq chan chan Job) {
	// 这是一个独立的协程，循环读取通道内的数据
	// 保存每读取一个通道参数就去工作，没有读到就阻塞
	go func() {
		for {
			// 注册工作通道到线程池
			wq <- w.JobQueue
			select {
			// 获取任务
			case job := <-w.JobQueue:
				job.Do()
			// 终止当前任务
			case <-w.Quit:
				return
			}
		}
	}()
}

// 线程池 领导
type WorkerPool struct {
	// 线程池中worker(工人)的数量
	workerlen int
	// 线程池的job通道
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		// 开始建立 workerlen 个worker(工人)协程
		workerlen: workerlen,
		// 工作队列通道
		JobQueue: make(chan Job),
		// 最大通道参数设为最大协程数workerlen工人的数量最大值
		WorkerQueue: make(chan chan Job, workerlen),
	}
}

// 运行线程池
func (wp *WorkerPool) Run() {
	//初始化时会按照传入的num，启动num个后台协程，然后循环读取Job通道里面的数据，
	//读到一个数据时，再获取一个可用的Worker，并将Job对象传递到该Worker的chan通道
	for i := 0; i < wp.workerlen; i++ {
		//新建 workerlen worker(工人) 协程(并发执行)，每个协程可处理一个请求
		worker := NewWorker()
		worker.Run(wp.WorkerQueue)
	}
	// 循环获取可用的worker,往worker中写job
	// 这是一个单独的协程只负责保证不断获取可用的worker
	go func() {
		for {
			select {
			//读取任务
			case job := <-wp.JobQueue:
				//尝试获取一个可用的worker作业通道
				//这将阻塞，直到一个worker空闲
				worker := <-wp.WorkerQueue
				worker <- job
			}
		}
	}()
}
