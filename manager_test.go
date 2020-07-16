package soa_test

import (
	"os"
	"testing"
	"unsafe"

	"github.com/andygeiss/assert"
	"github.com/andygeiss/soa"
)

func Test(t *testing.T) {
	a := soa.DefaultManager.Allocate([]int32{}).([]int32)
	a[0] = 1
	a[1] = 2
	a[111] = 3
	a[999] = 4
	b := soa.DefaultManager.Grow(a).([]int32)
	c := soa.DefaultManager.Pack(b).([]int32)
	d := soa.DefaultManager.Unpack(c).([]int32)
	assert.That("allocated size of [a] is equal to page size", t, len(a)*int(unsafe.Sizeof(int32(0))), os.Getpagesize())
	assert.That("grow size  of [b] is equal to 2* page size", t, len(b)*int(unsafe.Sizeof(int32(0))), 2*os.Getpagesize())
	assert.That("pack should shrink [c] to 4", t, len(c), 4)
	assert.That("unpack should grow [d] to page size", t, len(d), int(1024))
	assert.That("d[0] should be 1", t, d[0], int32(1))
	assert.That("d[1] should be 2", t, d[1], int32(2))
	assert.That("d[2] should be 3", t, d[2], int32(3))
	assert.That("d[3] should be 4", t, d[3], int32(4))
}

func Benchmark(b *testing.B) {
	b.Run("Alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			soa.DefaultManager.Allocate([]int32{})
		}
	})
	b.Run("Grow", func(b *testing.B) {
		s := soa.DefaultManager.Allocate([]int32{}).([]int32)
		for i := 0; i < b.N; i++ {
			soa.DefaultManager.Grow(s)
		}
	})
	b.Run("Pack", func(b *testing.B) {
		s := soa.DefaultManager.Allocate([]int32{}).([]int32)
		for i := 0; i < b.N; i++ {
			soa.DefaultManager.Pack(s)
		}
	})
	b.Run("Unpack", func(b *testing.B) {
		s := soa.DefaultManager.Allocate([]int32{}).([]int32)
		s = soa.DefaultManager.Pack(s).([]int32)
		for i := 0; i < b.N; i++ {
			soa.DefaultManager.Unpack(s)
		}
	})
}
