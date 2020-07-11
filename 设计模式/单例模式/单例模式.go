package 单例模式

import "sync"

// 懒汉方式：指全局的单例实例在第一次被使用时构建。
// 缺点：非线程安全。当正在创建时，有线程来访问此时ins = nil就会再创建，单例类就会有多个实例了。
//type singleton struct {}
//var sin *singleton
//func GetIns() *singleton {
//	if sin == nil {
//		return &singleton{}
//	}
//	return sin
//}

//饿汉方式：指全局的单例实例在类装载时构建。
//缺点：如果singleton创建初始化比较复杂耗时时，加载时间会延长。
//type singleton struct {}
//var sin *singleton = &singleton{}
//func GetIns() *singleton {
//	return sin
//}

//懒汉加锁
//缺点：虽然解决并发的问题，但每次加锁是要付出代价的。不管实例是否创建都都会加锁
//type singleton struct {}
//var sin *singleton
//var mu sync.Mutex
//func GetIns() *singleton {
//	mu.Lock()
//	defer mu.Unlock()
//	if sin == nil {
//		return &singleton{}
//	}
//	return sin
//}

//双重锁
//避免了每次加锁，提高代码效率。如果实例已被创建则不需要加锁
type singleton struct{}

var sin *singleton
var mu sync.Mutex

func GetIns() *singleton {
	if sin == nil {
		mu.Lock()
		defer mu.Unlock()
		if sin == nil {
			return &singleton{}
		}
	}
	return sin
}
