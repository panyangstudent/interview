thrift版本：v0.14.2
# 基本概念
我们都知道rpc(Remote Procedure Call，远程过程调用)，是一个计算机通信协议，此协议允许进程间通信。简单来说就是，当A机器上的进程调用B机器上的进程时，A机器上的调用进程会被挂起，而B机器上的进程开始执行。
调用方将参数信息传送给被调用方，然后可以通过被调用方的结果得到返回。rpc可以不依赖应用层协议，直接给予rpc进行远程调用，在传输层就可以完成通信。由于rpc调用方式依赖客户端与服务端之间建立Socket连接来实现二
进制通信，底层会比较复杂，所以一些rpc框架就应运而生。 市面上目前主流的关于rpc协议实现有grpc，thrift,dubbo。而这些rpc框架一般来说都需要解决服务寻址，数据流的序列化和反序列化，网络传输这三个主要问题。

# rpc调用的基本流程
![avater](图片/img.png)
* 服务消费方（Client 客户端）通过本地调用的方式调用服务。
* 客户端存根（Client Stub）接收到调用请求后负责将方法、入参等信息序列化（组装）成能够进行网络传输的消息体。
* 客户端存根（Client Stub）找到远程的服务地址，并且将消息通过网络发送给服务端。
* 服务端存根（Server Stub）收到消息后进行解码（反序列化操作）。
* 服务端存根（Server Stub， 一段代码）根据解码结果调用本地的服务进行相关处理
* 服务端(Server)本地服务业务处理。
* 处理结果返回给服务端存根（Server Stub）。
* 服务端存根（Server Stub）序列化结果。
* 服务端存根（Server Stub）将结果通过网络发送至消费方。
* 客户端存根（Client Stub）接收到消息，并进行解码（反序列化）。
* 服务消费方得到最终结果。

# http和rpc
http和rpc其实不是对立面，我们知道rpc只是一个计算机通信协议框架，通信协议只是其中的一部分。而http协议作为网络七层模型中应用层的协议，他的主要职责是解决如何包装数据。http协议是建立在传输层tcp协议之上，而传输层的tcp协议主要来解决数据传输问题，
但是对于上层应用开发极其不友好，所以就存在了http协议。 除此之外还有常见的socket，socket是针对TCP或UDP的具体接口实现，提供了在传输层进行网络编程的方法。对于rpc的具体实现grpc和thrift来说，grpc的底层实现是http2
协议，而thrift的底层实现是tcp协议。因此我们可以认为http和rpc协议是活跃在应用层的网络协议，他们处于同一层级，互相独立且交织。

# thrift架构
![avater](图片/thrift.png)
从上图我们可以看到，thrift仍然是基于rpc的基本调用流程。上图的黄色部分就是用户实现的业务逻辑，接下来service client/write() read()
是thrift根据IDL(接口描述文件)生成的客户端和服务端代码，包括数据的读写部分，对应于rpc调用的基本流程中的client stub和server stub。
Tprotocol用来对数据进行序列化与反序列化，具体方法包括二进制，json或者apache thrift定义的格式。TTransport提供数据数据传输功能

# thrift特性
1. 开发速度快
   编写thrift IDL文件，利用编译器自动生成服务端骨架(skeletons)和客户端桩(stubs)，省去了开发者自定义和维护接口编解码，消息传输，服务器多线程模型等基础工作。
   服务端：只需要按照服务骨架即接口，编写好具体的业务处理逻辑
   客户端：拷贝IDL定义好的客户端桩和服务对象，然后像调用本地方法一样调用远端服务
2. 接口维护简单
3. 学习成本低
4. 多语言支持
5. 稳定且广泛使用


# thrift IDL文件的数据类型
基本类型
   * bool：布尔值
   * byte：8位有符号整数
   * i16：16位有符号整数
   * i32：32位有符号整数
   * i64：:64位有符号整数
   * double：64位浮点数
   * binary：二进制串
结构体类型
   * struct：定义的结构体对象
容器类型
   * list：有序元素列表
   * set：无序无重复元素集合
   * map：有序的key/value键值对
服务类型
   * service：具体对应的服务类
异常类型
   * exception：异常类型

# 传输协议(TProtocol)
Thrift可以让用户选择客户端和服务端之间传输通信协议的区别，在传输协议上总体分为二进制(binary)和文本传输协议，一般情况都会选择二进制，来节省带宽，提高传输效率。
   * TBinaryProtocol：二进制编码格式进行数据传输
   * TCompactProtocol：高效率的，密集的二进制编码格式进行数据传输
   * TJSONProtocol：使用JSON的数据编码协议进行数据传输
   * TSimpleJSONProtocol：只提供json只写的协议，适用于通过脚本语言解析


# thrift的数据传输方式(TTransport)
TTransport是与底层数据传输紧密相关的传输层。在这一层，数据是按照字节流处理的，把这些字节按照顺序进行发送和接收，并且不会关心数据是什么类型。数据类型的解析是TProtocol这一层
   * TSocket：使用阻塞式IO进行传输
   * THttpTransport：采用http协议进行数据传输
   * TNonblockingTransport：使用非阻塞方式，用于构建异步客户端
   * TFramedTransport：使用非阻塞方式，按块大小进行传输
   * TFileTransPort：以文件形式进行传输
   * TMemoryTransport：将内存用于IO传输
   * TZlibTransport：使用zlib进行压缩，与其他传输方式联合使用
   * TBufferedTransport：对某个transport对象操作的数进行buffer，即从buffer中读取数据进行传输，或将数据直接写入buffer。

# thrift的服务端网络模型(TServer)
TServer在thrift框架中的主要任务是接收client的请求
   * TSimpleServer：单线程服务器端，使用标准的阻塞式IO
   * TThreadPoolServer：多线程服务器端，使用标准的阻塞式IO
   * TNonblockingServer：单线程服务器端，使用非阻塞式IO
   * THsHaServer：半同步半异步服务器端，基于非阻塞式IO读写和多线程工作任务处理
   * TThreadedSelectorServer：多线程选择服务器端，对THsHaServer在异步IO模型上进行增强

对于golang来说，只有TSimpleServer服务模式，并且是非阻塞的

# TProcessor(服务端)
主要是对Tserver中一次请求的inputProtocol和outputProtocol进行操作，也就是从inputProtocol中读取的client请求数据，向outputProtocol写入用户逻辑的返回值。TProcessor是一个非常关键的处理函数，因为client所有的rpc调用都会经过该函数处理并转发。

# ThriftClient(客户端)
ThriftClient跟TProcessor一样主要操作inputProtocol和outputProtocol，不同的是thrift将rpc调用分为send和receive两个步骤：
* send步骤，将用户的调用参数作为一个整体的struct写入TProtocol，并发送到TServer。
* send结束后，thriftClient便立即进入receive状态等待TServer的响应。对于TServer的响应，使用返回值解析类惊醒返回值解析，完成rpc调用

# thrift各个元素之间的关系


# TSimpleServer的服务模式
这不是一个典型的TSimpleServer，因为它在接受套接字后不会阻塞。 它更像一个TThreadServer，可以处理不同的连接在不同的goroutines。 如果golang用户实现了一个conn-pool一样的东西在客户端，这将有效。
```go
type TSimpleServer struct {
   closed int32
   wg     sync.WaitGroup
   mu     sync.Mutex
   
   processorFactory       TProcessorFactory
   serverTransport        TServerTransport
   inputTransportFactory  TTransportFactory
   outputTransportFactory TTransportFactory
   inputProtocolFactory   TProtocolFactory
   outputProtocolFactory  TProtocolFactory
   // THeaderProtocol中自动转发的报头
   forwardHeaders []string
   logger Logger
}
```
# 服务端Server代码
server端IDL
```thrift
include "User.thrift"
namespace go Sample

typedef map<string, string> Data

struct Response {
    1:required i32 errCode; //错误码
    2:required string errMsg; //错误信息
    3:required Data data;
}

//定义服务
service Greeter {
    Response SayHello(
        1:required User.User user
    )

    Response GetUser(
        1:required i32 uid
    )
}

service SimpleService {
    i32 add(1:i32 num1, 2:string num2)
}
```
## 具体go服务端代码实现如下：
```go

func SimpleServer() {
	//
   conf := &thrift.TConfiguration{
      ConnectTimeout: time.Second, // 连接超时时间
      SocketTimeout:  time.Second, // socket超时时间
      MaxFrameSize: 1024 * 256,
      TBinaryStrictRead:  thrift.BoolPtr(true),
      TBinaryStrictWrite: thrift.BoolPtr(true),
   }
   // 定义传输协议-二进制
   protocolFactory := thrift.NewTBinaryProtocolFactoryConf(conf)
   // 定义数据传输方式
   transportFactory := thrift.NewTTransportFactory()
   // 定义secket监听地址端口
   transport, _ := thrift.NewTServerSocket(":8090")
   
   processor := Sample.NewSimpleServiceProcessor(&handler.SimpleServiceHandler{})
   //阻塞式单线程服务器，阻塞式IO
   server := thrift.NewTSimpleServer4(processor,transport,transportFactory,protocolFactory)
   server.Serve()
}

func (p *TSimpleServer) Serve() error {
	p.logger = fallbackLogger(p.logger)
    
	err := p.Listen()
	if err != nil {
		return err
	}
	p.AcceptLoop()
	return nil
}

func (p *TSimpleServer) AcceptLoop() error {
   for {
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
	// 此处的Accept()是阻塞的， 是调用listener.Accept()
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
       go func() {
          defer p.wg.Done()
          if err := p.processRequests(client); err != nil {
               p.logger(fmt.Sprintf("error processing request: %v", err))
          }
       }()
   }
   return 0, nil
}
```
这里我们看到服务端在处理请求时，采用了协程的方式，如果服务端重启，那么这里对于业务是有损的。但是go thrift的最新版本采用了golang waitgroup的方式实现了优雅重启。
```go
func (p *TSimpleServer) processRequests(client TTransport) (err error) {
	defer func() {
		err = treatEOFErrorsAsNil(err)
	}()
    // 获取client的输入数据
	processor := p.processorFactory.GetProcessor(client)
	// 获取client的数据传输方式
	inputTransport, err := p.inputTransportFactory.GetTransport(client)
	if err != nil {
		return err
	}
	// 获取client的传输协议
	inputProtocol := p.inputProtocolFactory.GetProtocol(inputTransport)
	var outputTransport TTransport
	var outputProtocol TProtocol
	
	// 对于THeaderProtocol，我们必须使用相同的协议实例进行输入和输出，以便响应使用服务器检测到请求所在的相同方言。
	headerProtocol, ok := inputProtocol.(*THeaderProtocol)
	if ok {
		outputProtocol = inputProtocol
	} else {
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

		ctx := SetResponseHelper(
			defaultCtx,
			TResponseHelper{
				THeaderResponseHelper: NewTHeaderResponseHelper(outputProtocol),
			},
		)
		if headerProtocol != nil {
		    // 读取报头
			if err := headerProtocol.ReadFrame(ctx); err != nil {
				return err
			}
			ctx = AddReadTHeaderToContext(ctx, headerProtocol.GetReadHeaders())
			ctx = SetWriteHeaderList(ctx, p.forwardHeaders)
		}

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
```

