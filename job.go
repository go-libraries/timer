package timer

import "fmt"

//default Job
type TaskJob struct {
    Fn func()
}


func (j *TaskJob) OnStart(f func(reply Reply))  {
    fmt.Println("start")
    f(Reply{})
}

func (j *TaskJob) OnStop(f func(reply Reply))  {
    fmt.Println("start")
    f(Reply{})
}

func (j *TaskJob) OnFinish(f func(reply Reply))  {
    fmt.Println("start")
    f(Reply{})
}

func (j *TaskJob) OnSuccess(f func(reply Reply))  {
    fmt.Println("start")
    f(Reply{})
}

func (j *TaskJob) OnError(f func(reply Reply))  {
    fmt.Println("error")
    f(Reply{
        Msg:"error",
    })
}

func (j *TaskJob) Run(){
    j.Fn()
}