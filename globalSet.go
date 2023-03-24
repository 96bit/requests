package requests

import (
	"sync"
)

var WeChatUserAuth = NewSet(100)

// 集合结构体

type Set struct {
	m            map[string]struct{} // 用字典来实现，因为字段键不能重复
	len          int                 // 集合的大小
	sync.RWMutex                     // 锁，实现并发安全
}

// 新建一个空集合

func NewSet(cap int64) *Set {
	temp := make(map[string]struct{}, cap)
	return &Set{
		m: temp,
	}
}

// 增加一个元素
// 空结构体的内存地址都一样，并且不占用内存空间
// 元素作为字典的键，会自动去重。同时，集合大小重新生成
// 这里省去了自己判断数据是否在集合中的过程，是非常有用的技巧

func (s *Set) Add(item string) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{} // 实际往字典添加这个键
	s.len = len(s.m)       // 重新计算元素数量
}

// 移除一个元素
// 时间复杂度等于字典删除键值对的复杂度，哈希不冲突的时间复杂度为：O(1)，否则为 O(n)

func (s *Set) Remove(item string) {
	s.Lock()
	defer s.Unlock()

	// 集合没元素直接返回
	if s.len == 0 {
		return
	}

	delete(s.m, item) // 实际从字典删除这个键
	s.len = len(s.m)  // 重新计算元素数量
}

// 查看是否存在元素

func (s *Set) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// 查看集合大小

func (s *Set) Len() int {
	return s.len
}

// 清除集合所有元素

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]struct{}{} // 字典重新赋值
	s.len = 0                   // 大小归零
}

// 集合是够为空

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// 将集合转化为列表

func (s *Set) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := make([]string, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
