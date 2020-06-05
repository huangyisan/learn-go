## redis的基本使用
1. 查看key的数量 dbsize
2. 清空当前db数据 flushdb, 清空所有db数据 flushall

## redis五大数据类型
* string
    string是二进制安全的. 除了普通的字符串外, 也可以存放图片等数据.
    redis中字符串value最大是512M.
    set命令如果value存在,则是存储,如果已经存在,则是修改.
    del 命令用来删除key
    get 命令来获取值
    setex用来设置超时时间.  setex key second value
    mset 用来设置多个key value.  mset key value key value.
    gset 用来获取多个key value.  mget key key.


* hash
    hash是一个键值对的集合, 是一个string类型的field和value的映射表, hash特别适合用于存储对象.
    hset key field value, 如果设定一个相同的key的field的不同value, 则做update操作.
    hget 获取字段value
    hdel 删除
    hgetall key 获取key的所有字段和value. 该命令时间复杂度很高
    hmset key field1 value1 field2 value2 设置多个
    hexists key field 查看哈希表key中, 给定域field是否存在

* list
    列表是简单的字符串列表, 按照插入顺序排序. 你可以添加一个元素到列表的头部(左边)或者尾部(右边)
    list本质是个链表, list的元素是有序的, 元素的值可以重复.
    list的操作一般都分为两种lxx 和rxx ,l左边开始操作, r右边开始操作.
    lpush key value1 value2 value3. lpush表示从左边塞入,下一个value2 继续往左塞, 然后上一个value1被挤往右边. rpush相反
    lrange key start end. start从0开始, 倒数第一个可以表示为-1, 取出来的顺序.

4. set

5. zset(有序集合)



refer
> http://redisdoc.com/