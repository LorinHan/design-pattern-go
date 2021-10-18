package singleton

import "sync"

// 单例模式
type singleton struct {
}

// S 方式1，直接通过变量向外暴露
var S = &singleton{}

// 方式2，通过函数向外暴露，并发不安全
var s *singleton

func GetSingleton() *singleton {
	if s == nil {
		s = &singleton{}
	}
	return s
}

// 方式2.2，通过函数向外暴露， 通过锁来确保并发安全
var lock = &sync.Mutex{}

func GetSingletonByLock() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if s == nil {
		s = &singleton{}
	}
	return s
}

// 方式2.3，通过函数向外暴露， 通过once来确保并发安全
var once sync.Once

func GetSingletonByOnce() *singleton {
	once.Do(func() {
		s = &singleton{}
	})
	return s
}
