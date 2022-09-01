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
