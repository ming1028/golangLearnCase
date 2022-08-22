（1）尽可能的使用:=去初始化声明一个变量（在函数内部）；

（2）尽可能的使用字符代替字符串；

（3）尽可能的使用切片代替数组；

（4）尽可能的使用数组和切片代替映射；

（5）如果只想获取切片中某项值，不需要值的索引，尽可能的使用for range去遍历切片，这比必须查询切片中的每个元素要快一些；

（6）当数组元素是稀疏的（例如有很多0值或者空值nil），使用映射会降低内存消耗；(使用二维数组保存数组中特殊的值，索引是其位置)

（7）初始化映射时指定其容量；

（8）当定义一个方法时，使用指针类型作为方法的接受者；（内存）

（9）在代码中使用常量或者标志提取常量的值；

（10）尽可能在需要分配大量内存时使用缓存；

数组在栈上还是静态存储区都是一连串的内存空间，通过指向数组开头的指针、元素的数量以及元素类型占的空间大小表示数组
如果数组中元素的个数小于或者等于 4 个，那么所有的变量会直接在栈上初始化；
如果数组元素大于 4 个，变量就会在静态存储区初始化然后拷贝到栈上

runtime.mallocgc 是用于申请内存的函数
如果遇到了比较小的对象会直接初始化在 Go 语言调度器里面的 P 结构中，
而大于 32KB 的对象会在堆上初始化

哈希表实现方法：
底层的数据结构是数组，索引获得方法：index := hash("") % len(array) ;
开放寻址法： 写入数据时，如果发生冲突，就会将键值写入到下一个索引位置为空的位置; 读取的时候先获取键（key）的哈希
并取模，然后比较key是否相等或者空内存（未查找到）。开放寻址法中对性能影响最大的是装载因子，数组中
元素的数量与数组大小的比值，随着装载因子的增加，线性探测的用时会逐渐增加。
拉链法：相当于二维数组，等过取模选择到一个桶，然后遍历当前桶的链表，①找到键相同的键值对就更新键对应的值；
②没有找到对应的键值对，在链表末尾增加新的键值对；
装载因子 := 元素数量 ÷ 桶数量。在一般情况下使用拉链法的哈希表装载因子都不会超过 1，当哈希表的装载因子较大时会触发哈希的扩容，
创建更多的桶来存储哈希中的元素，保证性能不会出现严重的下降。

Interface
一种是带有一组方法的接口，另一种是不带任何方法的 interface{}

	                结构体实现接口	        结构体指针实现接口

结构体初始化变量 通过 不通过
结构体指针初始化变量 通过 通过

未实现interface的情况下，值类型可以调用值方法，对于指针方法的调用，是golang的语法糖，
存在转换（&p）,是否可以被寻址的。指针可以调用值方法同样存在隐式转换（*ptr）

传参都是值拷贝，当一个函数参数是interface类型时，调用会发生隐式的类型转换，
转换后的变量包含之前的变量和类型信息
Go 语言根据接口类型是否包含一组方法将接口类型分成了两类：

    使用 runtime.iface 结构体表示包含方法的接口
    使用 runtime.eface 结构体表示不包含任何方法的 interface{} 类型；

```
type eface struct {
  _type *_type // 类型信息
  data unsafe.Pointer // 指向数据的指针
}

type iface struct {
  tab *itab
  data unsafe.Pointer
}

type itab struct {
  inter *interfacetype // 接口自身的元信息
  _type *_type  // 具体类型的元信息
  link  *itab
  bad   int32
  hash  int32
  fun   [1]uintptr
}
```
interface是所有类型的父类，interface{}可以传入任何类型

map中数据总个数/桶个数>6.5时，引发翻倍扩容
使用了太多的溢出桶时，（溢出桶使用的太多会导致map处理速度降低）
B <= 15，已使用的溢出桶个数 >=2的B次方时，引发等量扩容
B > 15，已使用的溢出桶个数 >=2的15次方时，引发等量扩容

make 初始化内置的数据结构，切片、map、channel
new 根据传入的类型分配一片空间并返回这片内存空间的指针

### 变量

* 函数外的每个语句都必须以关键字开始（var, const, func等）
* :=不能使用在函数外
* const同时声明多个常量时，如果省略了值则表示和上面一行值相同。
* 常量不同于变量在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用。所以常量不可以被寻址。

***

### slice底层实现

```
type slice struct {
    array unsafe.Pointer // 指向数组的指针
    len int
    cap int
}
```

* 切片底层数组容量超过一定限制会在堆上分配内存
* 空切片、nil切片区别

```
slice := make([]int, 0) // 空切片 指向一个内存地址，没有分配任何内存空间
slice := []int{} // nil切片
```

* 扩容策略，切片容量小于1024，扩容就翻倍增加容量；超过1024，增长因子就变成1.25

### 指针

* 对于**引用类型**的变量，使用的时候不仅要声明还要分配内存空间
* **值类型的声明不需要分配内存空间**，因为在声明的时候已经默认分配了内存空间
* new(Type) 接收一个类型，返回一个该类型内存地址的指针，类型指针，值为类型的零值
* make只用于slice、map以及chan的内存创建返回类型本身，而不是指针类型，这三种类型就是引用类型

***

### Map原理

    通过key进行hash运算，转换为一个整形数字然后对数组长度取余，得到key存储在数组的下标位置，将key和value组装成结构体放在下标处

#### hash冲突

* 开放寻址法：写入数据时，如果发生冲突，就会将键值写入到下一个索引位置为空的位置;
  读取的时候先获取键（key）的哈希并取模，然后比较key是否相等或者空内存（未查找到）。
  开放寻址法中对性能影响最大的是装载因子，数组中元素的数量与数组大小的比值，
  随着装载因子的增加，线性探测的用时会逐渐增加。
    + 线性探测：从冲突的下标处往后探测，到达数组末尾，再从开始从探测，直到找到一个空位置存储，
      都找不到情况下会扩容。查找的时候，先哈希然后取模，比较key是否相等，如果不等按照顺寻探测，
      直到探测到一个空位置，说明key不存在。
* 拉链法：相当于二维数组，等过取模选择到一个桶，然后遍历当前桶的链表，①找到键相同的键值对就更新键对应的值；
  ②没有找到对应的键值对，在链表末尾增加新的键值对；
* 开放寻址法和拉链法优缺点：
    + 拉链法比线性探测处理简单
    + 线性探测查找会被拉链法更消耗时间
    + 线性探测会更容易导致扩容，而拉链不会
    + 拉链存储了指针，所以空间上会比线性探测占用多
    + 拉链是动态申请存储空间，更适合链长不确定

#### map存储

* bucket数据结构，tophash用来快速查找key值是否在该bucket中，kv存放方式 k1k2 v1v2,内存对齐节约内存，
  通过k获取hash值，hash值的低八位和bucket数组长度取余，定位数组的下标，hash值的高八位存储在bucket中的
  tophash中，用来快速判断key是否存在，key和value具体值通过指针运算存储，当一个bucket满时，通过overflow
  指针链接到下一个bucket

```
type bmap struct {
  tophash [bucketCnt]uint8 // bucketCnt等于8
}
编译期间gc.bmap结构
type bmap struct {
  topbits  [8]uint8
  keys     [8]keyType
  values   [8]valueType
  pad      uintptr
  overflow uintptr 
}
```

* 装载因子 := 元素数量 除 桶数量，装载因子越大，哈希的读写性能越差。
* 当哈希表元素数量少于或者等于25个时，会将键值一次性加入哈希表，超过25个时，会分别存储键值，会通过for循环加入。

#### map无序

* map扩容后会将部分key移至新内存，并未记录原数据位置，扩容后就已经是无序的
* map读取根据随机数来读取数据

### 闭包

闭包复制的是原对象的指针

***

### panic/recover问题

* 当前goroutine是否panic
* 当前函数是否已经退出
* 当前函数是否已被recover
* argp(调用当前recover的函数指针)和当前最上层函数调用defer的函数指针，中间必须隔着仅仅一层函数
  调用

***

### defer

* return之后的语句先执行，defer后的语句后执行
* panic仅有最后一个可以被revover捕获。

### 类型比较

* map只能和nil比较
* slice不能比较
* 结构体比较
    - 相同类型的结构体才可以比较，与属性类型个数、和属性顺序（匿名结构体）
    - 结构体中有不可以比较的类型：map、slice
* reflect.DeepEqual比较两个类型
* nil可以用作interface、function、pointer、map、slice和channel的“空值”

### 内存四区

* 栈区：空间较小，读写性能高，数据存放时间较短暂。由编译器自动分配和释放。
* 堆区：空间充裕，数据存放时间较久
* 全局区-静态全局变量区：全局变量的开辟是程序在main之前就已经放在内存中，对外完全可见，由操作系统回收。
* 全局区-常量区：常量为存放数值字面值单位，不可修改，常量名是其值的对等符号，无法取出地址，因为字面量符号并没有地址而言。

### 单元测试

* go test -cover(测试覆盖率) -coverprofile=c.out(输出到文件) -v(查看测试函数名称时间)
  -run="名称"(正则匹配)
* go tool cover -html=c.out 
* go test -bench=Split(基准测试函数名) -benchmem(获得内存分配的统计数据：123B/op每次操作分配123字节， 3 allocs/op每次操作3次内存分配)
* -benchtime=20(最小基准时间)