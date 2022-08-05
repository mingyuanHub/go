package main

import "fmt"

//参考文档：http://www.javashuo.com/article/p-pdxcczvd-gk.html

type workerPool struct {
	workerChan chan *Worker
	workerList []*Worker
}

var WorkerPool *workerPool

func (wp *workerPool) StopJobs(jobs []int64) {
	for _, id := range jobs {
		for _, w := range wp.workerList {
			if w.SM.JobID == id {
				w.SM.Stop(id)
			}
		}
	}
}

type Worker struct {
	ID      int
	RepJobs chan int64
	SM      *SM
	quit    chan bool
}


func (w *Worker) Start() {
	go func() {
		for {
			WorkerPool.workerChan <- w

			select {
			case jobId := <-w.RepJobs:
				w.handleRepJob(jobId)
			case q := <-w.quit:
				if q {
					return
				}
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func (w *Worker) handleRepJob(jobId int64) {
	err := w.SM.Reset(jobId)
	if err != nil {
		// 更新任务状态：初始化失败
		return
	}

	//参数及其他校验，如果失败则任务取消，更新任务状态：任务不可用

	//如果一切正常， 执行任务，传入其他参数
	w.SM.Start(jobId)
}


func NewWorker(id int) *Worker {
	w := &Worker{
		ID:      id,
		RepJobs: make(chan int64),
		quit:    make(chan bool),
		SM:      &SM{},
	}
	w.SM.Init()
	return w
}

func InitWorkerPool()  {

	//配置文件获取工人数量
	maxJobWorker := 10

	WorkerPool = &workerPool{
		workerChan: make(chan *Worker, maxJobWorker),
		workerList: make([]*Worker, maxJobWorker),
	}

	for i := 0; i < maxJobWorker; i ++ {
		worker := NewWorker(i)
		WorkerPool.workerList = append(WorkerPool.workerList, worker)
		worker.Start()
		fmt.Println("worker start", worker.ID)
	}
}

var jobQueue chan int64

func Dispatch()  {
	for  {
		select {
		case job := <- jobQueue:
			go func(jobId int64) {
				worker := <-WorkerPool.workerChan
				worker.RepJobs <- jobId
			}(job)
		}
	}
}

type SM struct {
	JobID int64
}

func (sm *SM) Init() error {
	fmt.Println("sm init")
	return nil
}

func (sm *SM) Stop(id int64)  {
	fmt.Println("stop job id", id)
}

func (sm *SM) Reset(id int64) error {
	fmt.Println("reset job id", id)
	return nil
}

func (sm *SM) Start(id int64) error {
	fmt.Println("start job id", id)
	return nil
}