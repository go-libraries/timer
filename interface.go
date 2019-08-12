package timer

type Job interface {
    Run()
    OnStart(f func(reply Reply))
    OnStop(f func(reply Reply))
    OnFinish(f func(reply Reply))
    OnSuccess(f func(reply Reply))
    OnError(f func(reply Reply))
}

type TaskInterface interface {
    TaskGetInterface
    TaskSetInterface
}

type TaskSetInterface interface {
    SetJob(job Job) TaskSetInterface
    SetRuntime(runtime int64) TaskSetInterface
    SetUuid(uuid string) TaskSetInterface
    SetSpacing(spacing int64) TaskSetInterface
    SetEndTime(endTime int64) TaskSetInterface
    SetRunNumber(number int) TaskSetInterface
}

type TaskGetInterface interface{
    RunJob()
    GetJob()  Job
    GetUuid() string
    GetRunTime() int64
    GetSpacing() int64
    GetEndTime() int64
    GetRunNumber() int
}

type TaskLogInterface interface {
    Println(v ...interface{})
}
