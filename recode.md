go语言中new和make的作用
---
- new函数
<br>
new(T)创建一个T类型的变量，初始化为类型零值，并返回变量地址。<br>
注意：new函数的返回值为T类型指针。<br>
````
p := new(int)   //p为int类型指针，指向未命名的变量。
fmt.Println(*p)  //输出“0”
````
- make
<br>
make 只可以处理三种：slice、map和channel，
会对数据初始化为空值，但是不返回地址。
<br>
内建函数make(T, args)与new(T)的用途不一样。
它只用来创建slice，map和channel，并且返回一个初始化的(而不是置零)，类型为T的值（而不是*T）。
之所以有所不同，是因为这三个类型的背后引用了使用前必须初始化的数据结构。
例如，slice是一个三元描述符，包含一个指向数据（在数组中）的指针，长度，以及容量，在这些项被初始化之前，slice都是nil的。对于slice，map和channel，make初始化这些内部数据结构，并准备好可用的值。