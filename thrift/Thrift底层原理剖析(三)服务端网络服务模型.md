
# 网络服务模型
Thrift提供的网络服务模型：单线程，多线程，事件驱动，从另一角度划分：阻塞服务类型，非阻塞服务类型
阻塞服务类型：TSimpleServer，TThreadPoolServer
非阻塞服务类型：TNonblockingServer， THsHaServer和TThreadedSelectorServer
![avater](图片/TServer.png)
上图这些都是TServer的具体实现,但是在golang只有TSimpleServer的网络服务模型.
# TServerTransport(服务传输器)
```go
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

type TServerSocket struct {
    listener      net.Listener 
    addr          net.Addr
    clientTimeout time.Duration
    
    // Protects the interrupted value to make it thread safe.
    mu          sync.RWMutex
    interrupted bool
}
```
TServerTransport是一个interface，可以认为是服务端传输器，用来响应上游的请求。该interface分别定义了4个方法。
  * Listen：启动监听
  * Accept：接收一个请求
  * Close：关闭监听
  * Interrupt：中断/结束当前服务处理器的监听

TServerSocket是TServerTransport的一个具体实现，TServerSocket结构体中包含如下字段：
  * listener：面向协议的通用网络监听器，多个goroutine可以同时调用Listener上的方法
  * addr：监听地址， IP+port
  * clientTimeout：客户端超时时间
  * mu：读写锁，保护被中断的值，使线程安全
# TTransportFactory
```go
type TTransportFactory interface {
	GetTransport(trans TTransport) (TTransport, error)
}

type tTransportFactory struct{}

```
创建传输的封装实例工厂，从一个serverTransport获取到传输内容，然后可能想要改变他们(即基于base transport创建一个BufferedTransport)


# TSimpleServer
```go
type TSimpleServer struct {
	closed int32  // 是否处于监听状态
	wg     sync.WaitGroup // 协程管理组
	mu     sync.Mutex // 互斥锁

	processorFactory       TProcessorFactory // 对一次请求的inputProtocol和outputProtocol进行操作，处理流程中的关键函数。
	serverTransport        TServerTransport // 实现开启监听，接收请求，关闭监听等功能，TServerTransport的一个具体实现
	inputTransportFactory  TTransportFactory // 从serverTransport中获取对应的请求输入
	outputTransportFactory TTransportFactory // 处理服务端的resp
	inputProtocolFactory   TProtocolFactory // req的协议，比如二进制传输，conn超时时间，socket超时时间，最大容量等
	outputProtocolFactory  TProtocolFactory // resp的协议，比如二进制传输，conn超时时间，socket超时时间，最大容量等

	// THeaderProtocol中自动转发的报文头
	forwardHeaders []string
    // 请求或者返回日志等
	logger Logger
}
```
由于传统的TSimpleServer是阻塞式IO，实现方式简单明了，但是每次只能接收和处理一个socket链接效率比较低。但是golang的TSimpleServer在接受套接字后不会阻塞，他更像一个TThreadPoolServer，可以在不同的goroutine协程中处理不同连接。
接下来我们看下TSimpleServer是怎么实现非阻塞式IO的
```go
func (p *TSimpleServer) Serve() error {
	p.logger = fallbackLogger(p.logger)
    // 判断当前服务是否在监听中
	err := p.Listen()
	if err != nil {
		return err
	}
	// 循环接收请求
	p.AcceptLoop()
	return nil
}

func (p *TSimpleServer) AcceptLoop() error {
	// 当前实现是一个死循环，只要服务启动，这里就会一直接收请求，但是这里的依然是阻塞的
    for {
		// 接收&处理请求，底层实现在开启协程后，立刻进行下一个请求的处理
        closed, err := p.innerAccept()
        if err != nil {
            return err
        }
        if closed != 0 {
            return nil
        }
    }
}


func (p *TSimpleServer) innerAccept() (int32, error) {
	// serverTransport的类型为TServerSocket，这里接收一个请求
    client, err := p.serverTransport.Accept()
    p.mu.Lock()
    defer p.mu.Unlock()
    closed := atomic.LoadInt32(&p.closed)
    if closed != 0 {
        return closed, nil
    }
    if err != nil {
        return 0, err
    }
    if client != nil {
        p.wg.Add(1)
		// 开启协程
        go func() {
            defer p.wg.Done()
            if err := p.processRequests(client); err != nil {
                p.logger(fmt.Sprintf("error processing request: %v", err))
            }
        }()
    }
    return 0, nil
}

// p.serverTransport.Accept()实现
func (p *TServerSocket) Accept() (TTransport, error) {
	// 先加读锁，以防其他线程抢到该请求
    p.mu.RLock()
    interrupted := p.interrupted
    p.mu.RUnlock()
    
    if interrupted {
        return nil, errTransportInterrupted
    }
    // 添加写锁
    p.mu.Lock()
    listener := p.listener
    p.mu.Unlock()
    if listener == nil {
        return nil, NewTTransportException(NOT_OPEN, "No underlying server socket")
    }
	// 获取一个通用的conn
    conn, err := listener.Accept()
    if err != nil {
        return nil, NewTTransportExceptionFromError(err)
    }
	// 返回一个TTransport，具体实现是TSocket类型，使用阻塞式IO来进行数据传输
    return NewTSocketFromConnTimeout(conn, p.clientTimeout), nil
}

// listener.Accept()实现
// Accept实现Listener接口中的Accept方法;它等待下一个调用并返回一个通用的Conn
func (l *TCPListener) Accept() (Conn, error) {
    if !l.ok() {
        return nil, syscall.EINVAL
    }
    c, err := l.accept()
    if err != nil {
        return nil, &OpError{Op: "accept", Net: l.fd.net, Source: nil, Addr: l.fd.laddr, Err: err}
    }
    return c, nil
}

// p.Listen实现
// 这里的serverTransport我们以TServerSocket为例
func (p *TServerSocket) Listen() error {
    p.mu.Lock()
    defer p.mu.Unlock()
	// 判断TServerSocket的listener不为空，即确认网络监听器已经初始化过，serverTransport的实现不会初始化listener，所以这里返回为false
    if p.IsListening() {
        return nil
    }
	// 在listener为nil的情况，这里重新创建一个Listener
    l, err := net.Listen(p.addr.Network(), p.addr.String())
    if err != nil {
        return err
    }
    p.listener = l
    return nil
}
// Checks whether the socket is listening.
func (p *TServerSocket) IsListening() bool {
    return p.listener != nil
}


// p.processRequests(client)的实现
func (p *TSimpleServer) processRequests(client TTransport) (err error) {
    defer func() {
    err = treatEOFErrorsAsNil(err)
    }()
	// 获取processor，这里就是返回TSimpleServer的processorFactory变量值 = SimpleServiceProcessor对象
    processor := p.processorFactory.GetProcessor(client)
	// 获取输入
    inputTransport, err := p.inputTransportFactory.GetTransport(client)
    if err != nil {
        return err
    }
	// 获取输入的传输协议， TBinaryProtocol
    inputProtocol := p.inputProtocolFactory.GetProtocol(inputTransport)
    var outputTransport TTransport
    var outputProtocol TProtocol

    // 判断inputProtocol是否是THeaderProtocol类型，THeaderProtocol传输协议支持二进制或压缩协议作为封装协议，保持输入输出的传输协议一致
    headerProtocol, ok := inputProtocol.(*THeaderProtocol)
    if ok {
        outputProtocol = inputProtocol
    } else {
		// 查看输出协议是否已提前设置
		oTrans, err := p.outputTransportFactory.GetTransport(client)
        if err != nil {
            return err
        }
        outputTransport = oTrans
        outputProtocol = p.outputProtocolFactory.GetProtocol(outputTransport)
    }

    if inputTransport != nil {
        defer inputTransport.Close()
    }
    if outputTransport != nil {
        defer outputTransport.Close()
    }
    for {
        if atomic.LoadInt32(&p.closed) != 0 {
            return nil
        }
		// 设置resp的
        ctx := SetResponseHelper(
            defaultCtx,
            TResponseHelper{
            THeaderResponseHelper: NewTHeaderResponseHelper(outputProtocol),
            },
        )   
        if headerProtocol != nil {
            // We need to call ReadFrame here, otherwise we won't
            // get any headers on the AddReadTHeaderToContext call.
            //
            // ReadFrame is safe to be called multiple times so it
            // won't break when it's called again later when we
            // actually start to read the message.
            if err := headerProtocol.ReadFrame(ctx); err != nil {
                return err
            }
            ctx = AddReadTHeaderToContext(ctx, headerProtocol.GetReadHeaders())
            ctx = SetWriteHeaderList(ctx, p.forwardHeaders)
        }
        // 处理当前请求
        ok, err := processor.Process(ctx, inputProtocol, outputProtocol)
        if errors.Is(err, ErrAbandonRequest) {
            return client.Close()
        }
        if errors.As(err, new(TTransportException)) && err != nil {
            return err
        }
        var tae TApplicationException
        if errors.As(err, &tae) && tae.TypeId() == UNKNOWN_METHOD {
            continue
        }
        if !ok {
            break
        }
    }
    return nil
}

// processor.Process的具体实现
func (p *SimpleServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
    name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
    if err2 != nil { return false, thrift.WrapTException(err2) }
    if processor, ok := p.GetProcessorFunction(name); ok {
        return processor.Process(ctx, seqId, iprot, oprot)
    }
    iprot.Skip(ctx, thrift.STRUCT)
    iprot.ReadMessageEnd(ctx)
    x24 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
    oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
    x24.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, x24

}
```
