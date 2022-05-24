package UserCache

import (
	"Practice/proto"
	"fmt"
	"sync"
)

type Cache struct {
	ClientCache map[int] *proto.ClientInfo
	CacheLock 	sync.RWMutex			//读写锁
}

var maxLen = 20000
var cache *Cache
var once sync.Once

//只执行一次
func GetCacheObj() *Cache {
	once.Do(func() {
		fmt.Println("Create Obj")
		cache = new(Cache)
		cache.ClientCache = make(map[int]*proto.ClientInfo, maxLen)
	})
	return cache
}


func GetCacheNum() int {
	//fmt.Printf("GetCacheNum %d \n", len(GetCacheObj().ClientCache) )
	return len(GetCacheObj().ClientCache)
}

func ExistUid(Uid int) bool {
	//defer GetCacheObj().CacheLock.RUnlock()
	//GetCacheObj().CacheLock.RLock()
	_, ok := GetCacheObj().ClientCache[Uid]
	return ok
}

func DelCache(Uid int)  {
	defer GetCacheObj().CacheLock.Unlock()
	GetCacheObj().CacheLock.Lock()
	delete(GetCacheObj().ClientCache,  Uid)

}

func AddCache(info *proto.ClientInfo) bool {

	defer GetCacheObj().CacheLock.Unlock()
	GetCacheObj().CacheLock.Lock()

	if len(GetCacheObj().ClientCache) >= maxLen{
		return false
	}

	if !ExistUid(info.Uid) {
		GetCacheObj().ClientCache[info.Uid] = info
		return true
	}
	return false
}
