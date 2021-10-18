# 设计模式-Golang

> 本工程中关于设计模式的一些理解和描述，均参考于《设计模式之禅》，程序代码也是参考书中的Java示例后以Golang改写，书中的程序示例有些写法适用于Java但不适合于Golang，也有些用Golang的特性可以更好的实现。



### 一、单例模式

> Ensure a class has only one instance, and provide a global point of access to it.
>
> 确保某一个类只有一个实例，而且自行实例化并向整个系统提供这个实例。

- 一个需要单例的类

```go
type singleton struct {
}
```

- 方式1，直接通过变量(首字母大写)向外暴露

```go
var S = &singleton{}
```

- 方式2，通过函数向外暴露

```go
var s *singleton

func GetSingleton() *singleton {
   if s == nil {
      s = &singleton{}
   }
   return s
}
```

 	但是这样实现会存在并发安全问题。

- 方式2.1，通过锁来确保并发安全

```go
var lock = &sync.Mutex{}

func GetSingletonByLock() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if s == nil {
		s = &singleton{}
	}
	return s
}
```

- 方式2.3，通过once来确保并发安全

```go
var once sync.Once

func GetSingletonByOnce() *singleton {
   once.Do(func() {
      s = &singleton{}
   })
   return s
}
```

- ***目的：***

  - 单例模式在内存中只有一个实例，减少了内存开支，特别是一个对象需要频繁地创建、销毁时，而且创建或销毁时性能又无法优化，单例模式的优势就非常明显。
  - 当一个对象的产生需要比较多的资源时，如读取配置、产生其他依赖对象时，则可以通过在应用启动时直接产生一 

  个单例对象，然后用永久驻留内存的方式来解决。

  



### 二、工厂模式

> Define an interface for creating an object,but let subclasses decide which class to instantiate.Factory Method lets a class defer instantiation to subclasses.
>
> 定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法使一个类的实例化延迟到其子类。



##### 1.简单工厂

##### 2.多工厂模式

> 当我们在做一个比较复杂的项目时，经常会遇到初始化一个对象很耗费精力的情况，所有的产品类都放到一个工厂方法中进行初始化会使代码结构不清晰。
>
> 例如，一个产品类有5个具体实现，每个实现类的初始化（不仅仅是new，初始化包括new一个对象，并对对象设置一定的初始值）方法都不相同，如果写在一个工厂方法中，势必会导致该方法巨大无比。

- ***目的***：

  考虑到需要结构清晰，我们就为每个产品定义一个创造者，然后由调用者自己去选择与
  哪个工厂方法关联。

- 将每个人种的工厂分离

```go
package factory

type ChineseFactory struct {
}

func (cf *ChineseFactory) Create() *chinese {
	return &chinese{
		humanBase{"黄", "中文"},
	}
}

type AmericanFactory struct {
}

func (cf *AmericanFactory) Create() *chinese {
	return &chinese{
		humanBase{"white", "English"},
	}
}
```

- 调用

```go
cf := &factory.ChineseFactory{}
c := cf.Create()
c.Talk()

af := &factory.AmericanFactory{}
a := af.Create()
a.Talk()
```

##### 3.单例工厂

> 单例模式的核心要求就是在内存中
> 只有一个对象，通过工厂方法模式也可以只在内存中生产一个对象

- singleTon单例和singleTonFactory单例工厂都小写，保证不能通过正常渠道建立对象

```go
type singleTon struct {
}

type singleTonFactory struct {
	single *singleTon
}
```

- 通过 init 来初始化单例以及静态单例工厂

```go
package factory

func init() {
	SF = &singleTonFactory{
		single: &singleTon{},
	}
}

var SF *singleTonFactory

type singleTon struct {
}

type singleTonFactory struct {
	single *singleTon
}

func (sf *singleTonFactory) Single() *singleTon {
	return sf.single
}
```

- 对外只暴露了SF这一个访问点，同时也无法通过其他方式来创建 singleTon 和 singleTonFactory，保证内存
  中的对象唯一。

- ***目的***：

  通过工厂方法模式创建一个单例对象，该框架可以继续扩展，在一个项目中可以产生一个单例构造器，所有需要产生单例的类都遵循一定的规则（小写不对外导出），然后通过扩展该框架，只要输入一个类型就可以获得唯一的一个实例。

##### 4.延迟初始化工厂

> 何为延迟初始化（Lazy initialization）？一个对象被消费完毕后，并不立刻释放，工厂类
> 保持其初始状态，等待再次被使用。延迟初始化是工厂方法模式的一个扩展应用

- lazyInitFactory 负责产品类对象的创建工作，并且通过m map变量产生一个缓存，对需要
  再次被重用的对象保留

```go
type product struct {
}

type lazyInitFactory struct {
	m map[string]*product
}
```

- 延迟加载的工厂类

```go
func init() {
	LF = &lazyInitFactory{make(map[string]*product)}
}

var (
	LF *lazyInitFactory
	lock = &sync.Mutex{}
)

func (lf *lazyInitFactory) Get(key string) *product {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := lf.m[key]; !ok {
		lf.m[key] = &product{}
	}
	return lf.m[key]
}
```

- 代码比较简单，通过定义一个Map容器，容纳所有产生的对象，如果在Map容器中已
  经有的对象，则直接取出返回；如果没有，则根据需要的类型产生一个对象并放入到Map容
  器中，以方便下次调用，需要注意的是map并发不安全，通过lock来确保。

- ***目的***：

  - 延迟加载框架是可以扩展的，例如限制某一个产品类的最大实例化数量，可以通过判断
    Map中已有的对象数量来实现，这样的处理是非常有意义的，例如数据库连接池，都会
    要求设置一个MaxConnections最大连接数量，该数量就是内存中最大实例化的数量。
  - 延迟加载还可以用在对象初始化比较复杂的情况下，例如硬件访问，涉及多方面的交
    互，则可以通过延迟加载降低对象的产生和销毁带来的复杂性。

##### 5. 抽象工厂模式

如果抽象工厂里面只定义一个方法，直接创建产品，那么就退化为工厂方法。

一个对象族（或是一组没有任何关系的对象）都有相同的约束，则可以使用抽象工厂模式。什么意思呢？例如一个文本编辑器和一个图片处理器，都是软件实体，但是 Linux下的文本编辑器和Windows下的文本编辑器虽然功能和界 面都相同，但是代码实现是不同的，图片处理器也有类似情况。也就是具有了共同的约束条 件：操作系统类型。

讲白了就是多种类型之间有着共同的约束，例如CPU和主板，都区分Inter和AMD，再如两个不同的软件，都区分操作系统的不同，两个人种、黄种人和白种人，区分男女；手机和充电器、也是得分android和 ios。

抽象工厂模式的使用场景定义非常简单：一个对象族（或是一组没有任何关系的对象） 都有相同的约束，则可以使用抽象工厂模式。