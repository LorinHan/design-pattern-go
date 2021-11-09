package builder

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

// Builder 建造者接口
type Builder interface {
	BuildFrame()
	BuildSeat()
	BuildTire()
	Get() *Bike
}

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
