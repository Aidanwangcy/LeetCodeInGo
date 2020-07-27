//https://juejin.im/post/5bdbd074e51d450549408fa8
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//懒汉模式
//example 示例结构体
type example struct {
	name string
}

var instance *example	//设置私有变量作为每次返回的单例

//GetExample 获取单例
func GetExample() *example {
	if instance == nil {
		instance = new(example)
	}
	return instance
}

func main1()  {
	s := GetExample()
	s.name = "第一次赋值单例模式"
	fmt.Println(s.name)

	s2 := GetExample()
	fmt.Println(s2.name)
}

//饿汉模式
type example2 struct {
	name string
}

var instance2 *example2

func init()  {
	instance2 = new(example2)
	instance2.name = "初始化单例模式2"
}

func GetExample2() *example2 {
	return instance2
}

func main2() {
	s := GetExample()
	fmt.Println(s.name)
}

//双重检查机制
var mux sync.Mutex
func GetInstance() *example {
	mux.Lock()
	defer mux.Unlock()
	if instance == nil {
		instance = &example{}
	}
	return instance
}

func GetInstance2() *example {
	if instance == nil {
		mux.Lock()
		defer mux.Unlock()
		if instance == nil {
			instance = &example{}
		}
	}
	return instance
}

var initialized uint32

func GetInstance3() *example {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mux.Lock()
	defer mux.Unlock()
	if initialized == 0 {
		instance = &example{}
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

//使用sync.Once
type example4 struct {
	name string
}

var instance4 	*example4
var once 		sync.Once

func GetInstance4() *example4 {
	once.Do(func () {
		instance4 = new(example4)
		instance4.name = "第一次赋值单例"
	})
	return instance4
}

func main()  {
	e1 := GetInstance4()
	fmt.Println(e1.name)

	e2 := GetInstance4()
	fmt.Println(e2.name)
}