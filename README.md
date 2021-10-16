### 工厂模式

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

  通过工厂方法模式创建一个单例对象，该框架可以继续扩展，在一个项目中可以
  产生一个单例构造器，所有需要产生单例的类都遵循一定的规则（小写不对外导出），然
  后通过扩展该框架，只要输入一个类型就可以获得唯一的一个实例。

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

  延迟加载框架是可以扩展的，例如限制某一个产品类的最大实例化数量，可以通过判断
  Map中已有的对象数量来实现，这样的处理是非常有意义的，例如数据库连接池，都会
  要求设置一个MaxConnections最大连接数量，该数量就是内存中最大实例化的数量。
  延迟加载还可以用在对象初始化比较复杂的情况下，例如硬件访问，涉及多方面的交
  互，则可以通过延迟加载降低对象的产生和销毁带来的复杂性。