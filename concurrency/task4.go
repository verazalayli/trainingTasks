package concurrency

//Канал отмены (or-channel)**
//
//Дано несколько `<-chan struct{}`.
//Написать функцию:
//
//```go
//func Or(chs ...<-chan struct{}) <-chan struct{}
//```
//
//Которая закрывается, когда **любой** из входных каналов закрылся.
