（1）尽可能的使用:=去初始化声明一个变量（在函数内部）；

（2）尽可能的使用字符代替字符串；

（3）尽可能的使用切片代替数组；

（4）尽可能的使用数组和切片代替映射（详见参考文献15）；

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
结构体初始化变量	       通过	                不通过
结构体指针初始化变量	   通过	                 通过

未实现interface的情况下，值类型可以调用值方法，对于指针方法的调用，是golang的语法糖，
  存在转换（&p）,是否可以被寻址的。指针可以调用值方法同样存在隐式转换（*ptr）

传参都是值拷贝，当一个函数参数是interface类型时，调用会发生隐式的类型转换，
转换后的变量包含之前的变量和类型信息
Go 语言根据接口类型是否包含一组方法将接口类型分成了两类：

    使用 runtime.iface 结构体表示包含方法的接口
    使用 runtime.eface 结构体表示不包含任何方法的 interface{} 类型；

map中数据总个数/桶个数>6.5时，引发翻倍扩容
    使用了太多的溢出桶时，（溢出桶使用的太多会导致map处理速度降低）
    B <= 15，已使用的溢出桶个数 >=2的B次方时，引发等量扩容
    B > 15，已使用的溢出桶个数 >=2的15次方时，引发等量扩容