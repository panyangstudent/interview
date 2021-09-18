1. map          
    map的任务是设计一种数据结构用来维护一个集合的数据，并且可以同时对集合进行增删改查的操作。最主要的数据结构有两种：散列表，搜索树。   

    * 散列表(O(1))     
        散列表用一个哈希函数将key分配到不同的桶(bucket，也就是数组的不同index)，这样的开销主要在哈希函数的计算以及常数的访问时间。哈希查找表一般会遇到碰撞问题，就是说不同的key被哈希到了同一个bucket中。一般有两种应对方法：链表法和开放地址法 

        * 链表法：将一个bucket实现成一个链表，落在同一个bucket中的key都会插入这个链表。
        * 开发地址法：在碰撞发生后，通过一定的规律，在数组的后面挑选"空位"，用来放置新的key
    * 搜索树 （最差O(logN))）      
        搜索树一般采用平衡搜索树，包括AVL树，红黑树


2.  golang的map底层初探     
    * golang的map底层其实实现了一个散列表，因此实现map的过程其实是实现散列表的过程。在这个散列表中，主要的数据结构有两个，一个是hmap（a header for a go map），另一个是bmap（a bucket for a Go map，通常叫其bucket）

        ```golang
        // A header for a Go map.
        type hmap struct {
            // 元素个数，调用 len(map) 时，直接返回此值
            count     int
            // 状态标识，比如正在被写、buckets和oldbuckets在被遍历、等量扩容 
            flags     uint
            // buckets 的对数 log_2
            B         uint8
            // 溢出桶里bmap大致的数量
            noverflow uint16
            // 计算 key 的哈希的时候会传入哈希函数
            hash0     uint32
            // 指向一个数组(连续内存空间)，数组的类型为[]bmap
            // 大小为 2^B
            // 如果元素个数为0，就为 nil
            // 这个字段我们可以称之为正常桶
            buckets    unsafe.Pointer
            // 扩容的时候，存放之前的buckets，buckets 长度会是 oldbuckets 的两倍
            oldbuckets unsafe.Pointer
            // 指示扩容进度，小于此地址的 buckets 迁移完成
            nevacuate  uintptr
            extra *mapextra // 溢出桶结构，正常桶里面某个bmap存满了，会使用这里面的内存空间存放键值对
        }
        //溢出额外信息
        type mapextra struct {
            overflow    *[]*bmap
            oldoverflow *[]*bmap

            nextOverflow *bmap
        }
        ```   
        ![avater](详细.png)
        * hmap是map最外层的一个数组结构，包括了map的各种基础信息。包括元素个数(count)，数组指针(buckets)，原来的数组（oldbuckets），当桶的个数为0时，hmap的buckets指向了一个空数组，当buckets需要扩容时，它会开辟一倍的内存空间，并且渐进式的将数组进行拷贝，即在使用到的时候才会将旧数组拷贝到新数组中。
        * key经过hash后得到hash值，共64位bit(64位机器)，计算它落在哪个桶时，只会用到最后B个bit位。如果 B = 5，那么桶的数量，也就是 buckets 数组的长度是 2^5 = 32。例如：一个key经过hash，等到的hash值如下：
            >10010111 | 000011110110110010001111001010100010010110010101010 │ 01010
        * 用低5位，也就是01010，值也就是10，就算出这个位于10号桶，在用高8位找到key在bucket中的位置。
        * 当两个key哈希值一样，就会发生哈希冲突，解决哈希冲突的方式是链表法：在bucket中，从前往后找到第一个空位。这样在查找某个key时，先找到对应的桶，在去遍历桶中的key。
        * 如果在bucket中没找到，并且overflow不为空，还要继续去overflow bucket 中寻找，直到找到或者所有的key槽位都找遍了，包括所有的 overflow bucket
        * 因为bucket里的key的其实地址就是unsafe.Pointer(b)+dataOffset。第i个key的地址就要此地址上跨过i个key的大小；其次value的地址又是在所有key之后，因此第i个value的地址还需要加上所有key的偏移。
        ```golang
        // key 定位公式
        k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))

        // value 定位公式
        v := add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.valuesize))
        ```
        * 在说下minTopHash，当一个bucket中的 tophash数组的元素值小于 minTopHash 时，表示当前这个这个元素值对应的key/value正在迁移，为了和正常的哈希值区分开，会给key计算出来的哈希值新增一个minTopHash，
    * golang中用于存储的结构就是bucket数组，bmap的结构如下：
    
        ```golang
        //编译后的桶
        type bmap struct {
            topbits  [8]uint8 //hash高8位
            keys     [8]keytype //键
            values   [8]valuetype //值
            pad      uintptr
            overflow uintptr // 指向hmap.extre.overflow溢出桶,上面的topbits，keys，values存满8个后，就会往这个溢出桶里存
        }
        ```  
        ![avater](bmap的结构.png)
        ![avater](bmap的底层.png)
        ![avater](buckets.png)  

        正常桶存满了怎么办？        
        * 正常桶存满了之后，正常桶中的bmap结构中的overflow指针就会指向对应使用溢出桶hamp.extra.overflow里的bmap的地址

        bucket这三部分内容决定了他怎么工作：    
        * 他的tophash存储的是哈希函数算出来的哈希值的高8位。用来加快索引。因为把高八位存储起来，这样就不能完整比较key就能过滤掉不符合的key，加快查询速度当一个哈希值的高8位和存储的高8位相符合，在去比较完整的key值，进而取出value。
        * 第二部分存储的是key和value，底层排列方式是，key放在一起，value放在一起。当key大于128字节时，bucket的key字段存储的会是指针，指向key的实际内容，value也是一样

        * 这样做的好处是key和value的长度不同时，可以消除padding带来的空间浪费。并且每个bucket最多存放8个键值对
        * 第三部分，存储的是当bucket溢出时，指向下一个bucket的指针
   
    ![avater](总体架构.png)
   
    * 触发扩容的时机：    

        * 装载因子超过阈值，源码里定义的阈值是 6.5
        * overflow的bucket数量过多：当B小于15，也就是bucket总数 2^B 小于 2^15 时，如果overflow的bucket数量超过2^B; 当 B >= 15，也就是 bucket 总数 2^B 大于等于 2^15，如果 overflow 的 bucket 数量超过 2^15。
    * 如何扩容   
        map的装载因子定义：     
        >loadFactor := count / (2^B) //count 就是 map 的元素个数，2^B 表示 bucket 数量。

        * 由于map的扩容需要将原来的key/value重新搬迁到新的内存地址上，如果大量的key/value搬迁的话。会非常影响性能，因此gomap的扩容采取了一种渐进式的方式，原来的key不会一次性搬迁完毕，每次最多只会搬迁两个bucket。

        * 分配新的 buckets，并将老的 buckets 挂到了 oldbuckets 字段上。真正搬迁动作是发生在插入或修改，删除key的时候，都会尝试进行搬迁bucket的工作。先检查oldbuckets是否搬迁完毕，简单来说就是查看oldbuckets是否为nil

        * 搬迁的目的是将老的buckets搬迁到新的buckets。而通过前面说的我们知道，应对条件1，新的buckets数量是之前的一倍，应对条件2，新的bucket数量和之前一样

        * 对于条件1，从老的buckets搬迁到了新的buckets，由于buckets的数量不变，因此可以按照序号来搬，比如原来在0号buckets，到了新的buckets后，还是放在0号buckets。

        * 对于条件2，就没那么简单了。要重新计算key的哈希值，才能决定它落在哪个bucket。例如，原来 B = 5，计算出 key 的哈希后，只用看它的低 5 位，就能决定它落在哪个 bucket。扩容后，B 变成了 6，因此需要多看一位，它的低 6 位决定 key 落在哪个 bucket。这称为 rehash

    * 理解了上面的流程，就知道了为什么map是无序的   
        * map在扩容之后，会发生key的搬迁，原来落在同一个buckets中的key，搬迁后，有些key就要远走高飞了(bucket 序号加上了 2^B)，而遍历的过程是按照顺序便利的，搬迁扩容之后，经历了rehash，有些key的位置发生了变化。所以也就是无序了。
        * 但是如果我就一个hard code的map，我也不会向 map 进行插入删除的操作，按理说每次遍历这样的 map 都会返回一个固定顺序的 key/value 序列吧。的确是这样，但是 Go 杜绝了这种做法

        * 当然golang做的更绝，在我们遍历map时，并不是固定从0号bucket遍历的，每次都是从一个随机值序号的bucket开始遍历，并且是随机从这个bucket的一个元素开始    

3. 线程安全的map-sync.map           
    * 使用场景      
        * 一写多读  
        * 各个协程的操作的key集合没有交集
    * 整体思路      
        sync.map的整体思路是用两个数据结构(只读的read和可写的dirty)，尽量将读写分离，较少锁对性能的影响

    * 实现      
        ```golang
            type Map struct {
            mu Mutex
            // 基本上你可以把它看成一个安全的只读的map
            // 它包含的元素其实也是通过原子操作更新的，但是已删除的entry就需要加锁操作了
            read atomic.Value // readOnly

            // 包含需要加锁才能访问的元素
            // 包括所有在read字段中但未被expunged（删除）的元素以及新加的元素
            dirty map[interface{}]*entry

            // 记录从read中读取miss的次数，一旦miss数和dirty长度一样了，就会把dirty提升为read，并把dirty置空
            misses int
        }

        type readOnly struct {
            m       map[interface{}]*entry
            amended bool // 当dirty中包含read没有的数据时为true，比如新增一条数据
        }

        // expunged是用来标识此项已经删掉的指针
        // 当map中的一个项目被删除了，只是把它的值标记为expunged，以后才有机会真正删除此项
        var expunged = unsafe.Pointer(new(interface{}))

        // entry代表一个值
        type entry struct {
            p unsafe.Pointer // *interface{}
        }
        ``` 
        实现原理    

        * 通过read和dirty两个字段将读写分离，读的数据存在只读字段read上，将最新写入的数据存在dirty字段上，
        * 读取时，会先查询read，不存在查询dirty，写入时只会写入dirty
        * 读取read并不需要加锁，而读或者写dirty都需要加锁
        * 另外有misses字段统计read被穿透的次数，，超过一定次数则将dirty数据同步到read上
        * 对于删除数据则直接通过标记来延迟删除

