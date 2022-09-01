package factory

import "sync"

func init() {
	LF = &lazyInitFactory{make(map[string]*product)}
}

var (
	LF *lazyInitFactory
	lock = &sync.Mutex{}
)

type product struct {
}

type lazyInitFactory struct {
	m map[string]*product
}

func (lf *lazyInitFactory) Get(key string) *product {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := lf.m[key]; !ok {
		lf.m[key] = &product{}
	}
	return lf.m[key]
}
