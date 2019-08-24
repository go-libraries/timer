package timer

import (
    "github.com/pkg/errors"
    "runtime"
)

//default Job
type TaskJob struct {
    Fn func()
    start   chan  bool
    err     chan  error
    done    chan  bool
    stop    chan  bool
    replies map[string]func(reply Reply)
}


func (j *TaskJob) OnStart(f func(reply Reply))  {
    j.replies["start"] = f
}

func (j *TaskJob) OnStop(f func(reply Reply))  {
    j.replies["stop"] = f
}

func (j *TaskJob) OnFinish(f func(reply Reply))  {
    j.replies["finish"] = f
}

func (j *TaskJob) OnError(f func(reply Reply))  {
    j.replies["error"] = f
}

func (j *TaskJob) run() {
    go j.Run()


    for {
        select {
             case e := <-j.err:
                 if f,ok := j.replies["error"]; ok {
                     reply := Reply{
                         Code:500,
                         Msg:e.Error(),
                         Err:e,
                     }
                     f(reply)
                 }
                 return

                 case <-j.start:
                     if f,ok := j.replies["start"]; ok {
                         f(Reply{})
                     }
                 case <-j.done:
                     if f,ok := j.replies["finish"]; ok {
                         f(successResult)
                     }
                     return
                 case <-j.stop:
                     //todo:
                     runtime.Goexit()
        }
    }
}

func (j *TaskJob) Stop(){
    j.stop <- true
}

func (j *TaskJob) Run(){
    j.start <- true
    isPanic := false
    defer func() {
        if x := recover(); x != nil {
            err := errors.Errorf("job error with panic:%v", x)
            j.err <- err
            isPanic = true
            return
        }
    }()

    j.Fn()

    defer func() {
        if !isPanic {
            j.done <- true
        }
    }()
}