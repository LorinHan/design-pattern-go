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

> 抽象工厂模式是工厂方法模式的升级版本，在有多个业务品种、业务分类时，通过抽象工厂模式产生需要的对象是一种非常好的解决方式。
> 
> 如果抽象工厂里面只定义一个方法，直接创建产品，那么就退化为工厂方法。
> 
> 一个对象族（或是一组没有任何关系的对象）都有相同的约束，则可以使用抽象工厂模式。什么意思呢？例如一个文本编辑器和一个图片处理器，都是软件实体，但是 Linux下的文本编辑器和Windows下的文本编辑器虽然功能和界 面都相同，但是代码实现是不同的，图片处理器也有类似情况。也就是具有了共同的约束条件：操作系统类型。
> 
> 讲白了就是多种类型之间有着共同的约束，例如CPU和主板，都区分Inter和AMD，再如两个不同的软件，都区分操作系统的不同，两个人种、黄种人和白种人，区分男女；手机和充电器、也是得分android和 ios。
> 
> 抽象工厂模式的使用场景定义非常简单：一个对象族（或是一组没有任何关系的对象） 都有相同的约束，则可以使用抽象工厂模式。

- 按照抽象工厂模式的定义，模拟一个场景：手机和充电器作为产品，操作系统作为约束条件，比如，IPhone13充电用Lightning接口的充电器，华为P30用Type-C接口的充电器。往下想就是：
  
  - 工厂得生产两种产品：一个手机、一个充电器；两个工厂：一个IPhone13工厂，一个华为P30工厂。
  
  - 两种产品，一个手机，一个是充电器，那么来一个手机接口和一个充电器接口：
    
    ```go
    type phone interface {
        Call()  // 打电话
        ConnectCharger(charger)  // 连接充电器
    }
    
    type charger interface {
        Charge()  // 充电
    }
    ```
  
  - 手机假设有IPhone13和华为P30，分别实现接口
    
    ```go
    // IPhone13 is phone的实现类
    type IPhone13 struct {
    }
    
    func (i IPhone13) Call() {
        fmt.Println("IPhone13 is calling...")
    }
    
    func (i IPhone13) ConnectCharger(c charger) {
        fmt.Println("IPhone13 is charging...")
        c.Charge()
    }
    
    // HuaWeiP30 is phone的实现类 华为P30
    type HuaWeiP30 struct {
    }
    
    func (h HuaWeiP30) Call() {
        fmt.Println("HuaWeiP30 is calling...")
    }
    
    func (h HuaWeiP30) ConnectCharger(c charger) {
        fmt.Println("HuaWeiP30 is charging...")
        c.Charge()
    }
    ```
  
  - 充电器也分Lightning接口和Type-C，分别实现接口：
    
    ```go
    type LightningCharger struct {
    }
    
    func (l LightningCharger) Charge() {
        fmt.Println("Charge by Lightning Charger")
    }
    
    type TypeCCharger struct {
    }
    
    func (t TypeCCharger) Charge() {
        fmt.Println("Charge by Type-C Charger")
    }
    ```
  
  - 接着说工厂：两个工厂，都是生产手机和充电器，那就来一个工厂接口：
    
    ```go
    type phoneFactory interface {
        CreatePhone() phone
        CreateCharger() charger
    }
    ```
  
  - IPhone13工厂生产IPhone13和Lightning，HuaweiP30工厂生产华为P30手机和Type-C充电器
    
    ```go
    type IPhone13Factory struct {
    }
    
    func (f IPhone13Factory) CreatePhone() phone {
        return &IPhone13{}
    }
    
    func (f IPhone13Factory) CreateCharger() charger {
        return &LightningCharger{}
    }
    
    type HuaWeiP30Factory struct {
    }
    
    func (h HuaWeiP30Factory) CreatePhone() phone {
        return &HuaWeiP30{}
    }
    
    func (h HuaWeiP30Factory) CreateCharger() charger {
        return &TypeCCharger{}
    }
    ```
  
  - 以上，全部实现完成，接下来看看调用：
    
    ```go
    i13Factory := &factory.IPhone13Factory{}
    iPhone13 := i13Factory.CreatePhone()
    iPhone13.Call()
    lightning := i13Factory.CreateCharger()
    iPhone13.ConnectCharger(lightning)
    
    hwFactory := &factory.HuaWeiP30Factory{}
    hwP30 := hwFactory.CreatePhone()
    hwP30.Call()
    typeC := hwFactory.CreateCharger()
    hwP30.ConnectCharger(typeC)
    ```
  
  - 每个工厂生产出来的手机和充电器接口函数都一样，那完全可以面向接口编程，反正调用的函数都一样，那么，也可以这样调用
    
    ```go
    // GetPhone 根据类型获取不同的工厂，然后面向接口编程即可
    func GetPhone(t int) (phone, charger) {
        var pf phoneFactory
    
        switch t {
        case 1:
            pf = &IPhone13Factory{}
        case 2:
            pf = &HuaWeiP30Factory{}
        default:
            return nil, nil
        }
    
        return pf.CreatePhone(), pf.CreateCharger()
    }
    
    func main() {
        p1, c1 := factory.GetPhone(1)
        p1.Call()
        p1.ConnectCharger(c1)
    
        p2, c2 := factory.GetPhone(2)
        p2.Call()
        p2.ConnectCharger(c2)
    }
    ```

- [参考这篇知乎的文章](https://www.zhihu.com/question/20367734)

- 优点：封装性，每个产品的实现类不是高层模块要关心的，它要关心的是什么？是接口，是抽象，它不关心对象是如何创建出来，这由谁负责呢？工厂类，只要知道工厂类是谁，就能创建出一个需要的对象并且只需面向它的接口就能使用它的所有能力，无需关心这个对象具体的实现，省时省力。

- 缺陷：抽象工厂模式虽然横向扩展容易，但最大缺点就是垂直扩展非常困难。横向与纵向是对于工厂而言的，如上面例子中的工厂，如果新增一个新型手机xxx，那么只需实现一个新的xxx工厂即可，这是横向：
  
  ```go
  type IPhone13Factory struct {                                        type xxxFactory struct {
  }                                                                    }
  
  func (f IPhone13Factory) CreatePhone() phone {                        func (f xxxFactory) CreatePhone() phone {    
      return &IPhone13{}                                                    return &xxx{}
  }                                                                    }
  
  func (f IPhone13Factory) CreateCharger() charger {                    func (f xxxFactory) CreateCharger() charger {
      return &LightningCharger{}                                            return &xxxCharger{}    
  }                                                                    }
  ```
  
  但是如果想要多出一个产品，比如这个工厂现在不仅生产手机和充电器了，还要生产耳机了，那么手机工厂的接口 也就是 phoneFactory就得多出一个方法 CreateEarphone()，然后Iphone13工厂实现这个函数生产Air pods，华为p30工厂也得实现，生产Free buds耳机，这就是纵向的扩展。

---

### 三、模板方法模式

> Define the skeleton of an algorithm in an operation,deferring some steps to subclasses.Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.（定义一个操作中的算法的框架，而将一些步骤延迟到子类中。使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。

- 直接来尝试一个例子，假设一个给用户发送提醒消息的场景，提醒分为短信提醒和邮件提醒，这个场景在项目中十分常见

- 首先，定义一个提醒者接口和一个具体的提醒者帮助类接口

```go
type reminder interface {
    SendTo(id int)  // 模板方法
}
type reminderHelper interface {
    GetUser(id int) string
    Msg(user string) bool
}
```

- 其次，实现提醒者的模板类，在模板方法SendTo整合调用流程：先获取发送对象，然后执行具体发送逻辑

```go
// RemindTemp 提醒发送的模板类
type RemindTemp struct {
    reminderHelper
}

func (rt *RemindTemp) SendTo(id int) {
    user := rt.GetUser(id)
    res := rt.Msg(user)
    fmt.Printf("记录日志，提醒用户：%s，提醒结果：%t\n", user, res)
    fmt.Printf("插入数据库，提醒用户：%s，插入结果：%t\n", user, res)
}
```

- 最后，实现具体的提醒者，只需实现reminderHelper接口即可
  
  - 邮件提醒
  
  ```go
  type EmailReminder struct {
  }
  
  func (e *EmailReminder) GetUser(id int) string {
      if id == 1 {
          return "2450978570@qq.com"
      }
      return "123456@163.com"
  }
  
  func (e *EmailReminder) Msg(user string) bool {
      fmt.Println("发送邮件咯，目标：", user)
      fmt.Println("发送成功")
      return true
  }
  ```
  
  - 短信提醒
  
  ```go
  type PhoneReminder struct {
  }
  
  func (p *PhoneReminder) GetUser(id int) string {
      if id == 1 {
          return "15345922954"
      }
      return "10086"
  }
  
  func (p *PhoneReminder) Msg(user string) bool {
      fmt.Println("发送短信咯，目标：", user)
      fmt.Println("发送成功")
      return true
  }
  ```

- 增加一个获取提醒者的方法

```go
func GetReminder(mode string) reminder {
    switch mode {
    case "email":
        return &RemindTemp{&EmailReminder{}}
    case "shortMsg":
        return &RemindTemp{&PhoneReminder{}}
    }
    return nil
}
```

- 调用

```go
func TestTemplate() {
    reminder := template.GetReminder("email")
    reminder.SendTo(1)

    reminder = template.GetReminder("shortMsg")
    reminder.SendTo(1)
}
```

- 模板方法模式非常简单，是一个应用非常广泛的模式。其中，RemindTemp叫做抽象模板，它的方法分为两类：
  - 基本方法，也叫做基本操作，是由子类实现的方法，并且在模板方法被调用。 
  - 模板方法，可以有一个或几个，一般是一个具体方法，也就是一个框架，实现对基本方法的调度，完成固定的逻辑。
- 优点：
  - 封装不变部分，扩展可变部分。把认为是不变部分的算法封装到父类实现，而可变部分的则可以通过继承来继续扩展。
  - 提取公共部分代码，便于维护。
  - 行为由父类控制，子类实现。基本方法是由子类实现的，因此子类可以通过扩展的方式增加相应的功能，符合开闭原则。
- 场景：
  - 多个类有公有的方法，并且逻辑基本相同时。
  - 重要、复杂的算法，可以把核心算法设计为模板方法，周边的相关细节功能则由各个子类实现。
  - 重构时，模板方法模式是一个经常使用的模式，把相同的代码抽取到父类中，然后通过钩子函数约束其行为。

---

### 四、建造者模式

> 建造者模式（Builder Pattern）也叫做生成器模式.
> 
> Separate the construction of a complex object from its representation so that the same 
> 
> construction process can create different representations.（将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。）

- 以共享单车的制造为例
- 1. 一辆单车具有车架、座椅、轮胎

```go
// Bike 自行车，包含框架、座椅、轮胎
type Bike struct {
    Frame
    Seat
    Tire
}

// Frame 框架
type Frame struct {
    Name string
}

// Seat 座椅
type Seat struct {
    Name string
}

// Tire 轮胎
type Tire struct {
    Name string
}
```

- 2. 单车建造者

```go
// Builder 建造者接口
type Builder interface {
    BuildFrame()
    BuildSeat()
    BuildTire()
    Get() *Bike
}
```

- 3. 建造者的实现类，ofo单车建造者和摩拜单车建造者

```go
// OfoBuilder Ofo单车建造者实现类
type OfoBuilder struct {
    bike Bike
}

func (o *OfoBuilder) BuildFrame() {
    o.bike.Frame = Frame{Name: "ofo合金框架"}
}

func (o *OfoBuilder) BuildSeat() {
    o.bike.Seat = Seat{Name: "ofo软座"}
}

func (o *OfoBuilder) BuildTire() {
    o.bike.Tire = Tire{Name: "ofo塑料轮胎"}
}

func (o *OfoBuilder) Get() *Bike {
    return &o.bike
}

// MobikeBuilder 摩拜单车建造者实现类
type MobikeBuilder struct {
    bike Bike
}

func (m *MobikeBuilder) BuildFrame() {
    m.bike.Frame = Frame{Name: "摩拜碳框架"}
}

func (m *MobikeBuilder) BuildSeat() {
    m.bike.Seat = Seat{Name: "摩拜橡胶座椅"}
}

func (m *MobikeBuilder) BuildTire() {
    m.bike.Tire = Tire{Name: "摩拜充气轮胎"}
}

func (m *MobikeBuilder) Get() *Bike {
    return &m.bike
}
```

- 4. 导演类，用来控制单车的生产过程

```go
// Director 导演类
type Director struct {
    B Builder
}

func (d *Director) Bike() *Bike {
    if d.B == nil {
        return nil
    }

    d.B.BuildFrame()
    d.B.BuildSeat()
    d.B.BuildTire()
    return d.B.Get()
}
```

- 调用

```go
func TestBike() {
    // 建造ofo单车
    d1 := &builder.Director{B: &builder.OfoBuilder{}}
    ofoBike := d1.Bike()
    fmt.Println(ofoBike)

    // 建造摩拜单车
    d2 := &builder.Director{B: &builder.MobikeBuilder{}}
    mobike := d2.Bike()
    fmt.Println(mobike)

    emptyDirector := &builder.Director{}
    fmt.Println(emptyDirector.Bike())
}
```

- 优点
  
  - 封装性：使用建造者模式可以使客户端不必知道产品内部组成的细节，如例子中我们就不需要关心每一个具体的单车内部是如何实现的，产生的对象类型就是Bike接口。
  - 建造者独立，容易扩展：OfoBuilder和MobikeBuilder是相互独立的，对系统的扩展非常有利。
  - 便于控制细节风险：由于具体的建造者是独立的，因此可以对建造过程逐步细化，而不对其他的模块产生任何影响。 

- 使用场景：
  
  - 相同的方法，不同的执行顺序，产生不同的事件结果时，可以采用建造者模式。
  - 多个部件或零件，都可以装配到一个对象中，但是产生的运行结果又不相同时，则可以使用该模式。 
  - 产品类非常复杂，或者产品类中的调用顺序不同产生了不同的效能，这个时候使用建造者模式非常合适。
  - 在对象创建过程中会使用到系统中的一些其他对象，这些对象在产品对象的创建过程中不易得到时，也可以采用建造者模式封装该对象的创建过程。该种场景只能是一个补偿方法，因为一个对象不容易获得，而在设计阶段竟然没有发觉，而要通过创建者模式柔化创建过程，本身已经违反设计的最初目标。 

- 注意事项
  
  建造者模式关注的是零件类型和装配工艺（顺序），这是它与工厂方法模式最大不同的地方，虽然同为创建类模式，但是注重点不同。

---

### 五、代理模式

> 代理模式（Proxy Pattern）也叫委托模式，是一个使用率非常高的模式，其定义如下： 
> 
> Provide a surrogate or placeholder for another object to control access to it.（为其他对象提供一种代理以控制对这个对象的访问。
> 
> 使用场景：
> 
> 1.日志的采集  
> 2.权限控制  
> 3.实现aop  
> 4.Mybatis mapper  
> 5.Spring的事务  
> 6.全局捕获异常  
> 7.Rpc远程调用接口  
> 8.分布式事务原理代理数据源

##### 1. 基本的代理模式

以游戏为例，有玩家和代练两种角色，模拟代练帮玩家玩游戏的场景：

- 玩家和代练都属于player，定义一个接口

```go
type IGamePlayer interface {
    Login(name, password string)
    KillBoss()
    UpGrade()
}
```

- 玩家

```go
type GamePlayer struct {
    Name  string
    level int
}

...GamePlayer实现IGamePlayer接口
```

- 代练，组合一个玩家，调用其功能，以此实现代理

```go
type GamePlayerProxy struct {
    Player IGamePlayer
}
func (gp *GamePlayerProxy) Login(name, password string) {
    gp.Player.Login(name, password)
}

func (gp *GamePlayerProxy) KillBoss() {
    gp.Player.KillBoss()
}

func (gp *GamePlayerProxy) UpGrade() {
    gp.Player.UpGrade()
}
```

- 调用

```go
func TestBaseProxy() {
    player := &proxy.GamePlayer{Name: "Lorin"}
    p := &proxy.GamePlayerProxy{Player: player}

    p.Login("lorin", "123456")
    p.KillBoss()
    p.UpGrade()

    p.KillBoss()
    p.UpGrade()
}
```

- 代理模式的优点：
  
  - 职责清晰 
    
    真实的角色就是实现实际的业务逻辑，不用关心其他非本职责的事务，通过后期的代理 
    
    完成一件事务，附带的结果就是编程简洁清晰。
  
  - 高扩展性 
    
    具体主题角色是随时都会发生变化的，只要它实现了接口，甭管它如何变化，都逃不脱 
    
    如来佛的手掌（接口），那我们的代理类完全就可以在不做任何修改的情况下使用。
  
  - 智能化 
    
    这在以上还没有体现出来，不过在我们以下的动态代理中就会看到 

- 为什么要用代理？想想现实世界吧，打官司为什么要找个律师？因为你不想参与中间过程的是是非非，只要完成自己的答辩，其他的比如事前调查、事后追查都由律师来搞定，这就是为了减轻你的负担。代理模式的使用场景非常多，Spring AOP就是一个非常典型的动态代理。

> 网络上代理服务器设置分为透明代理和普通代理，是什么意思呢？
> 
> 透明代理就是用户 不用设置代理服务器地址，就可以直接访问，也就是说代理服务器对用户来说是透明的，不用知道它存在的；
> 
> 普通代理则是需要用户自己设置代理服务器的IP地址，用户必须知道代理的存在。

##### 2. 普通代理模式

普通代理，它的要求就是客户端只能访问代理角色，而不能访问真实角色。

以上面的例子作为扩展，我自己作为一个游戏玩家，我肯定自己不练级了，也就是场景类不能再直接new一个GamePlayer对象了，它必须由GamePlayerProxy来进行。

- 修改玩家，限制外部无法创建

```go
// 限制外部无法创建
type normalGamePlayer struct {
    name  string
    level int
}
```

- 修改代练，提供初始化函数来为其植入玩家角色

```go
type NormalGamePlayerProxy struct {
    player IGamePlayer
}

func (gp *NormalGamePlayerProxy) Init(name string) {
    gp.player = &normalGamePlayer{name: name}
}
```

- 调用，调用者只知道代理存在就可以，不用知道代理了谁

```go
func TestNormalProxy() {
    p := &proxy.NormalGamePlayerProxy{}
    p.Init("Lorin")

    p.Login("Lorin", "123456")
    p.KillBoss()
    p.UpGrade()

    p.KillBoss()
    p.UpGrade()
}
```

在该模式下，调用者只知代理而不用知道真实的角色是谁，屏蔽了真实角色的变更对高层模块的影响，真实的主题角色想怎么修改就怎么修改，对高层次的模块没有任何的影响，只要你实现了接口所对应的方法，该模式非常适合对扩展性要求较高的场合。

##### 3. 强制代理模式

> 强制代理在设计模式中比较另类，为什么这么说呢？一般的思维都是通过代理找到真实的角色，但是强制代理却是要“强制”，你必须通过真实角色查找到代理角色，否则你不能访问。
> 
> 甭管你是通过代理类还是通过直接new一个主题角色类，都不能访问，只有通过真实角色指定的代理类才可以访问，也就是说由真实角色管理代理角色。
> 
> 这么说吧，高层模块new了一个真实角色的对象，返回的却是代理角色，就好比是你和一个明星比较熟，相互认识，有件事情你需要向她确认一下，于是你就直接拨通了明星的电话，明星却告诉你他很忙，这种事情找他助理去，或者他让助理帮你办。 

- 修改玩家接口，增加一个GetProxy函数

```go
type IForceGamePlayer interface {
    Login(name, password string)
    KillBoss()
    UpGrade()
    GetProxy() IForceGamePlayer
}
```

- 修改玩家，增加一个代练者、实现GetProxy函数，并且修改其它函数，使其它函数只能通过代理调用

```go
type ForceGamePlayer struct {
    Name  string
    level int
    Proxy IForceGamePlayer
}

func (g *ForceGamePlayer) GetProxy() IForceGamePlayer {
    if g.Proxy == nil {
        g.Proxy = &ForceGamePlayerProxy{player: g} // 传入自己
    }
    return g.Proxy
}

func (g *ForceGamePlayer) Login(name, password string) {
    if g.IsProxy() {
        fmt.Printf("用户 %s 登陆成功！\n", name)
    } else {
        fmt.Println("请使用指定的代理访问")
    }
}

func (g *ForceGamePlayer) IsProxy() bool {
    if g.Proxy != nil {
        return true
    }
    return false
}
```

- 代练没有太大变化，也是实现一下GetProxy函数，这里可以思考下，如果代理还有代理呢？也就是代练的代练？？

```go
type ForceGamePlayerProxy struct {
    player IGamePlayer
}

func (gp *ForceGamePlayerProxy) Login(name, password string) {
    gp.player.Login(name, password)
}

func (gp *ForceGamePlayerProxy) KillBoss() {
    gp.player.KillBoss()
}

func (gp *ForceGamePlayerProxy) UpGrade() {
    gp.player.UpGrade()
}

// GetProxy 如果没有代理的代理，就返回自己
func (gp *ForceGamePlayerProxy) GetProxy() IForceGamePlayer {
    return gp
}
```

- 调用

```go
func TestForceProxy() {
   player := &proxy.ForceGamePlayer{Name: "Lorin"}
   // 无法访问
   player.Login("lorin", "123456")
   player.KillBoss()
   player.UpGrade()

   fmt.Println("==================================")

   // 获取player的代理，通过代理访问
   p := player.GetProxy()
   p.Login("lorin", "123456")
   p.KillBoss()
   p.UpGrade()
}
```
