1. map          
    map的任务是设计一种数据结构用来维护一个集合的数据，并且可以同时对集合进行增删改查的操作。最主要的数据结构有两种：哈希查找表，搜索树。   

    * 哈希查找表(O(1))     
        哈希查找表用一个哈希函数将key分配到不同的桶(bucket，也就是数组的不同index)，这样的开销主要在哈希函数的计算以及常数的访问时间。哈希查找表一般会遇到碰撞问题，就是说不同的key被哈希到了同一个bucket中。一般有两种应对方法：链表法和开放地址法 

        * 链表法：将一个bucket实现成一个链表，落在同一个bucket中的key都会插入这个链表。
        * 开发地址法：在碰撞发生后，通过一定的规律，在数组的后面挑选"空位"，用来放置新的key
    * 搜索树 （最差O(logN))）      
    搜索树一般采用平衡搜索树，包括AVL树，红黑树
2. map 内存模型
    ```golang
    // A header for a Go map.
    type hmap struct {
        // 元素个数，调用 len(map) 时，直接返回此值
        count     int
        flags     uint8
        // buckets 的对数 log_2
        B         uint8
        // overflow 的 bucket 近似数
        noverflow uint16
        // 计算 key 的哈希的时候会传入哈希函数
        hash0     uint32
        // 指向 buckets 数组，大小为 2^B
        // 如果元素个数为0，就为 nil
        buckets    unsafe.Pointer
        // 扩容的时候，buckets 长度会是 oldbuckets 的两倍
        oldbuckets unsafe.Pointer
        // 指示扩容进度，小于此地址的 buckets 迁移完成
        nevacuate  uintptr
        extra *mapextra // optional fields
    }
    ```     
* 