package timer

import "sync"

type Reply struct {
    Code int64 `json:"code"`
    Msg string `json:"msg"`
    Err error   `json:"err"`
    Uuid string `json:"uuid"`
}

type TaskScheduler struct {
    tasks   *sync.Map
    swap    *sync.Map
    running *sync.Map
    add     chan TaskInterface
    remove  chan string
    stop    chan struct{}
    Logger  TaskLogInterface
    lock	bool
}


func GetSuccessResult(msg string)  Reply {
    if msg != "" {
        successResult.Msg = msg
    }
    return successResult
}

func GetErrorResult(code int64, msg string, err error) Reply  {
    return Reply{
        Code:code,
        Msg:msg,
        Err:err,
    }
}