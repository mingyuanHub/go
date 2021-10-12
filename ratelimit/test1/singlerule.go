package main

import (
	"sync"
	"time"
)

type singleRule struct {
	defaultExpiration            time.Duration       //表示计时周期,同时也是每条访问记录需要保存的时长，超过这个时长的数据记录将会被清除
	cleanupInterval              time.Duration       //默认多长时间需要执行一次清除过期数据操作
	numberOfAllowedAccesses      int                 //在计时周期内最多允许访问的次数
	estimatedNumberOfOnlineUsers int                 //在计时周期内预计有多少个用户会访问网站，建议选用一个稍大于实际值的值，以减少内存分配次数
	visitorRecords               []*circleQueueInt64 //用于存储用户的每一条访问记录
	usedRecordsIndex             sync.Map            //visitorRecords中已使用的数据索引,key代表用户名或IP,value代表visitorRecords中的下标位置
	notUsedVisitorRecordsIndex   map[int]struct{}    //对应visitorRecords中未使用的数据的下标位置，其自身非并发安全，其并发安全由locker实现,因sync.Map计算长度不优
	locker                       *sync.Mutex         //并发安全锁
}

func newSingleRule(defaultExpiration time.Duration, numberOfAllowedAccesses int) *singleRule {

	var this = &singleRule{
		defaultExpiration:          defaultExpiration,
		cleanupInterval:            60,
		numberOfAllowedAccesses:    numberOfAllowedAccesses,
		notUsedVisitorRecordsIndex: map[int]struct{}{},
		locker:                     &sync.Mutex{},
	}

	this.visitorRecords = make([]*circleQueueInt64 , 5000)

	for index := range this.visitorRecords {
		this.visitorRecords[index] = newCircleQueueInt64(this.numberOfAllowedAccesses)
		this.notUsedVisitorRecordsIndex[index] = struct{}{}
	}

	go this.deleteExpired()

	return this
}

func (this *singleRule) deleteExpired() {
	ok := true
	for range time.Tick(this.cleanupInterval) {
		if ok {
			ok = false
			this.deleteExpiredOnce()
		}
	}
}

func (this *singleRule) deleteExpiredOnce() {
	this.usedRecordsIndex.Range(func(key, index interface{}) bool {
		this.locker.Lock()
		usedIndex := index.(int)
		if usedIndex >= 0 {
			this.visitorRecords[usedIndex].DeleteExpired(key)
			if this.visitorRecords[usedIndex].UsedSize() == 0 {
				this.usedRecordsIndex.Delete(usedIndex)
				this.notUsedVisitorRecordsIndex[usedIndex] = struct{}{}
			}
		} else {
			this.usedRecordsIndex.Delete(key)
		}
		this.locker.Unlock()
		return true
	})
}

func (this *singleRule) allowVisit(key interface{}) bool{
	return this.add(key) == nil
}

func (this *singleRule) add(key interface{}) error{
	this.locker.Lock()
	defer  this.locker.Unlock()

	if  v, ok := this.usedRecordsIndex.Load(key); ok {
		usedIndex := v.(int)
		if this.visitorRecords[usedIndex].key == key {
			this.visitorRecords[usedIndex].DeleteExpired(key)
			return this.visitorRecords[usedIndex].Push(time.Now().Add(this.defaultExpiration).UnixNano())
		}
	}

	if len(this.notUsedVisitorRecordsIndex) > 0 {
		for usedIndex := range this.notUsedVisitorRecordsIndex {
			delete(this.notUsedVisitorRecordsIndex, usedIndex)
			this.visitorRecords[usedIndex].key = key
			this.usedRecordsIndex.Store(key, usedIndex)
			return this.visitorRecords[usedIndex].Push(time.Now().Add(this.defaultExpiration).UnixNano())
		}
	}

	circleQueue := newCircleQueueInt64(this.numberOfAllowedAccesses)
	circleQueue.key = key
	this.visitorRecords = append(this.visitorRecords, circleQueue)
	usedIndex := len(this.visitorRecords) -1
	this.usedRecordsIndex.Store(key, usedIndex)
	return this.visitorRecords[usedIndex].Push(time.Now().Add(this.defaultExpiration).UnixNano())
}