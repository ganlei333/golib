package concurrentMap

import (
	"strconv"
	"sync"
	"testing"
)

func BenchmarkConcurrentMap_Set(b *testing.B) {
	cm := CreateConcurrentMap(128)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Set(ConvertInt64(int64(i)), "test")
	}
}

func BenchmarkSyncMap_Store(b *testing.B) {
	var sm sync.Map
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Store(strconv.Itoa(i), "test")
	}
}

func BenchmarkConcurrentMap_Get(b *testing.B) {
	cm := CreateConcurrentMap(128)
	for i := 0; i < b.N; i++ {
		cm.Set(ConvertInt64(int64(i)), "test")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Get(ConvertInt64(int64(i)))
	}
}

func BenchmarkSyncMap_Load(b *testing.B) {
	var sm sync.Map
	for i := 0; i < b.N; i++ {
		sm.Store(strconv.Itoa(i), "test")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Load(strconv.Itoa(i))
	}
}

func BenchmarkConcurrentMap_Del(b *testing.B) {
	cm := CreateConcurrentMap(128)
	for i := 0; i < b.N; i++ {
		cm.Set(ConvertInt64(int64(i)), "test")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Del(ConvertInt64(int64(i)))
	}
}

func BenchmarkSyncMap_Delete(b *testing.B) {
	var sm sync.Map
	for i := 0; i < b.N; i++ {
		sm.Store(strconv.Itoa(i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Delete(strconv.Itoa(i))
	}
}
