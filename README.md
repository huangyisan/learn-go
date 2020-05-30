# learn-go
go


### 函数调用底层机制
1. 栈区存放基本数据类型
2. 堆区存放引用数据类型
3. 当调用函数的时候，会为该函数开辟新的栈区，函数内部的变量等存放在该函数的栈区中。函数的数据空间是独立的。
4. main函数开始被压入栈，当调用其他函数的时候，其他函数被压入，执行完的函数，会销毁这个函数对应的栈空间。main函数永远都是最后出栈。

### 注意点
1. go里面的数组是值传递而不是引用，和python不一样。属于值传递的还有基本类型。
2. 如果希望函数内的变量能够修改函数外的变量(指的是默认以值传递的方式的数据类型),可以传入变量的地址&,函数内以指针的方式操作变量，从效果上看类似引用。

### 可排序、可比较和不可比较
1. 可排序的数据类型有三种，Integer，Floating-point，和String
2. 可比较的数据类型除了上述三种外，还有Boolean，Complex，Pointer，Channel，Interface和Array
3. 不可比较的数据类型包括，Slice, Map, 和Function

### 引用类型和值类型
1. 在golang中只有三种引用类型它们分别是切片slice、字典map、管道channel。其它的全部是值类型，