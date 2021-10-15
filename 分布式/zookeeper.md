## 是什么？
是一个开源的分布式协调服务，他的目标是提供给高性能，高可用和顺序访问控制的能力，同时也是为了解决分布式环境下数据一致性问题。

## 集群 
zookeeper集群中有几个关键的概念，leader，folloer和observer。zookeeper中通常只有leader节点可以写入，follower和observer都只是负责读，但是follower会参与节点的选举和过半写成功， observer则不会，他只是单纯的提供读数据的功能。

## 数据节点
zookeeper中数据存储在内存之中，这个数据节点就叫做znode，他是一个树形结构，比如a/b/c类似 

znode又可以分为持久节点，临时节点，顺序节点三大类。

* 持久节点是指只要被创建除非主动移除，否则应该一直保存在zookeeper中
* 临时节点不同的是，他的生命周期和客户端session一样，会话失效那么临时节点就会被移除。
* 还有就是临时顺序节点和持久顺序节点，除了基本的特性之外，子节点的名称还具有有序性

## 会话session          
指zookeeper客户端和服务端之间的通信，他们使用长链接的方式保持通信，通常肯定是有心跳上报检测的机制，同时他可以接受来自服务器的watch事件的通知

## 事件监听器wather     
用户可以在指定的节点上注册wather，这样在事件触发的时候，客户端就会收到来自服务端的通知

## 权限控制ACL      
zookeeper使用ACL来进行权限控制，包含以下5种：   
1. create：创建子节点权限
2. delete：删除子节点权限
3. read：获取节点数据和子节点权限的
4. write：更新节点权限  
5. admin：设置节点ACL权限

所以，zookeeper通过集群的方式来做到高可用，通过内存数据节点znode来达到高性能，但是存储的数据量不能太大，通常适用于读多写少的场景。



## 应用场景     
1. 命名服务Name service，以来zookeeper可以生成全局唯一的节点ID，来对分布式中的资源进行管理  
2. 分布式协调，利用wather的监听机制，一个系统的某个节点状态发生改变，另外系统可以得到通知
3. 集群管理，分布式集群中状态的监控和管理，使用zookeeper来存储
4. master选举，利用zookeeper节点的全局唯一性，同时只有一个客户端能够创建并选举成功的特点，可以作为master选举使用


## wather监听机制和他的原理     
zookeeper可以提供分布式数据的发布/订阅，依赖的就是wather监听机制。客户端可以向服务端注册wather监听，服务端的指定事件触发之后，就会向客户端发送一个事件通知，他具有以下特性：    
1.  一次性：一旦一个watcher触发之后，zookeeper就会将他从存储中移除
2. 客户端串行：客户端的wather回调处理是串行同步的过程，不要因为一个wather的逻辑阻塞整个客户端
3. 轻量：wather通知的单位是watherevent，只包含通知状态，事件类型和节点路径，不包含具体事件的事件内容，具体的时间内需要客户端主动去获取数据
* 具体流程如下：    
    1. 客户端向服务端注册wather监听
    2. 保存wather对象到客户端本地的watherManager中
    3. 服务端wather事件触发之后，客户端收到服务端通知，从watherManager中取出对应的wather对象执行回调逻辑


## zookeeper如何保证数据一致性的        
zookeeper通过zab原子广播协议来实现数据的最终一致性，它是一个类似二阶段提交的过程。由于zookeeper只有leader节点可以写入数据，如果是其他节点收到写入数据的请求，则会将之转发给leader节点。     
* 主要流程如下：    
    1. leader再收到请求之后，将他转换为一个proposal提议，并且为每个提议分配一个全局唯一的递增事务ID：zxid，然后吧提议放入FIFO队列，按照FIFO的策略发送给所有的follower。
    2. Follower再收到提议之后，以事务日志的形式写入到本地磁盘中，写入成功后返回ack给到leader
    3. leader再收到超过半数的follower的ack之后，即可认为写入成功，就会发送commit命令给follower告诉他们可以提交proposal了

zab原子广播包含两种模式：崩溃恢复和消息广播         
整个集群服务在启动，网络中断或者重启等异常情况的时候，都会进入崩溃恢复状态，此时会通过选举产生leader节点，当集群过半的节点都和leader状态同步之后，zab就会推出恢复模式，之后进入广播模式。


## zookeeper是如何进行leader选举的      
leader的选举可以分为两个方面，同时选举主要包含事务zxid和myid，节点主要包含leading/following/looking3种状态。
1. 服务启动期间的选举
2. 服务运行期间的选举

* 服务启动期间的选举    
1. 首先每个节点都会对自己投票，然后吧投票信息广播给集群的其他节点
2. 节点接收到其他节点的投票信息，然后和自己的投票进行比较，首先zxid较大的优先，如果zxid相同则会选择myid更大者，此时大家都是looking的状态
3. 投票完成之后，开始统计投票信息，如果集群中过半的机器都选择了某个机器作为leader，那么投票结束
4. 最后更新各个节点的状态，leader改为leading状态，follower为following状态

* 服务运行期间的选举        
如果开始选出来的leader节点宕机了，那么运行期间就会重新进行leader的选举
1. leader宕机之后，非observer节点都会把自己的状态改成looking状态，然后重新进入选举流程
2. 生成投票信息(myid,zxid)，同样，第一轮的投票大家都会把投票投给自己，然后把投票信息广播出去
3. 接下来的流程和上面的选举一样，都会优先以zxid，然后选择myid，最后统计投票信息，修改节点状态，选举结束

## 选举之后怎么进行数据同步             
实际上zookeeper在选举之后，follower和observer就会向leader注册，然后就会开始数据同步的过程。     
数据同步包含3个主要值和4种形式。    
1. peerLastZxid：Learner服务器最后处理的zxid           
2. mincommittedLog：Leader提议缓存队列中最小zxid   
3. maxcommittedLog：leader提议缓存队列中最大的zxid   

* 直接差异化同步Diff同步    
如果peerLastZxid在minCommittedLog和maxCommittedLog之间，那么则说明learner服务器还没有完全同步最新的数据。
    1. 首先leader向learner发送diff指令，代表开始差异化同步，然后把差异数据(从peerlastzxid到maxcommittedlog之间的数据)提议proposal发送到l earner
    2. 发送完成之后发送一个newleader命令给leader，同时learner返回ack表示已经完成了同步
    3. 接着等待集群中过半的learner响应ack之后，就发送了一个uptodate命令，lerner返回ack，同步流程结束
* 先回滚在差异化同步 trunc + diff 同步      
这个设置针对的是一个异常的场景。    
如果leader刚生成一个proposal，还没有来的及发送出来，此时leader宕机，重新选举之后作为follower，但是新的leader没有这个proposal数据。

* 仅回滚同步 TRUNC同步          
针对peerLastZxid大于maxcommittedlog的场景，流程和上述的一样，事务会被回滚到maxcommitetedLog的记录

* 全量同步 SNAP同步

## 有可能会出现数据不一致的问题吗？