# e
为了最大程度的减少 golang 的错误处理代码量，本库提供一些规则和帮助函数。

# Usage
## if err 相应的封装
### 替代 `if (err != nil) return err`
```go
if e.IfErr(err){
    return err
}
```

### 替代 `return errors.WithStack` 和 `WithMessage`
```go
if e.IfErr(err){
    return e.Ws(err) // or e.Wm(err, msg)
}
```

## log 的封装
### 替代 `if (err != nil) log.Trace(err)`
```go
if e.Trace(err){
    return err
}
```
logrus 提供的所有等级都有相应的代理
```go
e.Trace("Something very low level.")
e.Debug("Useful debugging information.")
e.Info("Something noteworthy happened!")
e.Warn("You should probably take a look at this.")
e.Error("Something failed but I'm not quitting.")
// Calls os.Exit(1) after logging
e.Fatal("Bye.")
// Calls panic() after logging
e.Panic("I'm bailing.")
```
### 替代 `if (err != nil) log.Traces("%+v\n", errors.WithStack(err))`
```go
if e.Traces(err){
    return err
}
```
其他等级相同

## Is 和 As
Is 和 As 只是做了透明代理，用法和原本一样
```go
e.Is(err, target)
```
### 替代 `if e.Is() log.Trace`
```go
if e.IsTrace(err, target){
    return
}
```
### 替代 `if e.As() log.Trace`
```go
if e.AsTrace(err, target){
    return
}
```
### ```log.Traces("%+v\n", errors.WithStack(err))```
相应的也有 `IsTraces` 和 `AsTraces`


# 鸣谢
引用了 [logrus](https://github.com/sirupsen/logrus) 提供的日志分级
引用了 [errors](github.com/pkg/errors) 提供的异常栈