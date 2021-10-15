## 什么是raft算法   
    raft算法是分布式系统中的保持数据一致性的算法        
## raft算法怎么保证数据的一致性 
    raft有三种角色：    
        leader：管理节点，负责告诉其他节点如何处理数据或者消息      
        candīdate：候选节点，用于选举   
        follower：跟随节点，负责处理leader和canditate的请求 
    所有一致性算法都会涉及到状态机，而状态机保证系统从一个一致的状态开始，以相同的顺序执行一些列指令最终会达到另一个一致的状态。

    其中集群的各节点的状态转化如下：
        1. 所有节点初始状态都是follower角色 
        2. 超时时间内没有收到leader的请求则转换为Candidate进行选举
        3. Candidate收到大多数节点的选票则转换成leader；发现leader或者收到更高任期的请求则转换为Follower
        4. Leader在收到更高任期的请求后转换为Follower
    raft算法的任期：
        每个任期都由一次选举开始，若选举失败则这个任期内没有leader；如果选举出了leader则这个任期内由leader负责集群状态的管理
