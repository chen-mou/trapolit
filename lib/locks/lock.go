package locks

import (
	"sync"
	"trapolit/lib/conf"
	"trapolit/lib/i18n"
	"trapolit/lib/utils"
)

//TODO 编写锁相关的方法

var locks sync.Map

var storeLock = sync.Mutex{}

func Lock(name string) {
	temp, ok := locks.Load(name)
	if !ok {
		storeLock.Lock() //加锁后再检查一遍
		if temp, ok = locks.Load(name); !ok {
			temp = sync.Mutex{}
			locks.Store(name, temp)
		}
	}
	lock := temp.(sync.Locker)
	lock.Lock()
}

func UnLock(name string) {
	temp, ok := locks.Load(name)
	if !ok {
		panic(utils.NewError(i18n.Lang(conf.Cfg.Language), "ERROR.COMMON.NOT_FOUND"))
	}
	lock := temp.(sync.Locker)
	lock.Unlock()
}
