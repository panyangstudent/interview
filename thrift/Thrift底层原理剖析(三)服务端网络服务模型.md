
# 网络服务模型
Thrift提供的网络服务模型：单线程，多线程，事件驱动，从另一角度划分：阻塞服务类型，非阻塞服务类型
阻塞服务类型：TSimpleServer，TThreadPoolServer
非阻塞服务类型：TNonblockingServer， THsHaServer和TThreadedSelectorServer
![avater](图片/TServer.png)
上图这些都是TServer的具体实现,但是在golang只有TSimpleServer的网络服务模型.


# TSimpleServer
```go
type TSimpleServer struct {
	closed int32
	wg     sync.WaitGroup
	mu     sync.Mutex

	processorFactory       TProcessorFactory // 对一次请求的inputProtocol和outputProtocol进行操作，处理流程中的关键函数。
	serverTransport        TServerTransport // 实现开启监听，接收请求，关闭监听等功能
	inputTransportFactory  TTransportFactory
	outputTransportFactory TTransportFactory
	inputProtocolFactory   TProtocolFactory
	outputProtocolFactory  TProtocolFactory

	// Headers to auto forward in THeaderProtocol
	forwardHeaders []string

	logger Logger
}
```
由于传统的TSimpleServer是阻塞式IO，实现方式简单明了，但是每次只能接收和处理一个socket链接效率比较低。但是golang的TSimpleServer在接受套接字后不会阻塞，他更像一个TThreadPoolServer，可以在不同的goroutine协程中处理不同连接。


# TServerTransport(服务传输器)
```go
// Server transport. Object which provides client transports.
type TServerTransport interface {
	Listen() error
	Accept() (TTransport, error)
	Close() error

	// Optional method implementation. This signals to the server transport
	// that it should break out of any accept() or listen() that it is currently
	// blocked on. This method, if implemented, MUST be thread safe, as it may
	// be called from a different thread context than the other TServerTransport
	// methods.
	Interrupt() error
}
```
TServerTransport是一个interface，可以认为是服务端传输器，用来响应上游的请求。该interface分别定义了4个方法。
* Listen：启动监听
* Accept：接收一个请求
* Close：关闭监听
* Interrupt：中断/结束当前服务处理器的监听
