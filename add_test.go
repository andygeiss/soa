package soa_test

import (
	"math/rand"
	"testing"

	"github.com/andygeiss/assert"
	"github.com/andygeiss/soa"
)

func TestAddFloat32s(t *testing.T) {
	a := soa.DefaultManager.Allocate([]float32{}).([]float32)
	b := soa.DefaultManager.Allocate([]float32{}).([]float32)
	c := soa.DefaultManager.Allocate([]float32{}).([]float32)

	for i := 0; i < len(a); i++ {
		a[i] = 1.0
		b[i] = 8.0
	}

	soa.AddFloat32s(a, b, c)

	assert.That("len c should be 512", t, len(c), 1024)
	assert.That("c[0] should be 9.0", t, c[0], 9.0)
	assert.That("c[511] should be 9.0", t, c[1023], 9.0)
}

func TestAddFloat64s(t *testing.T) {
	a := soa.DefaultManager.Allocate([]float64{}).([]float64)
	b := soa.DefaultManager.Allocate([]float64{}).([]float64)
	c := soa.DefaultManager.Allocate([]float64{}).([]float64)

	for i := 0; i < len(a); i++ {
		a[i] = 1.0
		b[i] = 8.0
	}

	soa.AddFloat64s(a, b, c)

	assert.That("len c should be 512", t, len(c), 512)
	assert.That("c[0] should be 9.0", t, c[0], 9.0)
	assert.That("c[511] should be 9.0", t, c[511], 9.0)
}

func TestAddInt32s(t *testing.T) {
	a := soa.DefaultManager.Allocate([]int32{}).([]int32)
	b := soa.DefaultManager.Allocate([]int32{}).([]int32)
	c := soa.DefaultManager.Allocate([]int32{}).([]int32)

	for i := 0; i < len(a); i++ {
		a[i] = 0x1234567
		b[i] = 0x8765432
	}

	soa.AddInt32s(a, b, c)

	assert.That("len c should be 512", t, len(c), 1024)
	assert.That("c[0] should be 141972530", t, c[0], 141972530)
	assert.That("c[1023] should be 141972530", t, c[1023], 141972530)
}

func TestAddInt64s(t *testing.T) {
	a := soa.DefaultManager.Allocate([]int64{}).([]int64)
	b := soa.DefaultManager.Allocate([]int64{}).([]int64)
	c := soa.DefaultManager.Allocate([]int64{}).([]int64)

	for i := 0; i < len(a); i++ {
		a[i] = 0x12345678
		b[i] = 0x87654321
	}

	soa.AddInt64s(a, b, c)

	assert.That("len c should be 512", t, len(c), 512)
	assert.That("c[0] should be 2576980377", t, c[0], 2576980377)
	assert.That("c[511] should be 2576980377", t, c[511], 2576980377)
}

func BenchmarkAddFloat32s(b *testing.B) {
	b.Run("Native", func(b *testing.B) {
		b.ResetTimer()

		u := soa.DefaultManager.Allocate([]float32{}).([]float32)
		v := soa.DefaultManager.Allocate([]float32{}).([]float32)
		w := soa.DefaultManager.Allocate([]float32{}).([]float32)

		for i := 0; i < len(u); i++ {
			u[i] = rand.Float32()
			v[i] = rand.Float32()
		}

		for j := 0; j < len(u); j++ {
			w[j] = u[j] + v[j]
		}
		for i := 0; i < b.N; i++ {
			for j := 0; j < len(u); j++ {
				w[j] = u[j] + v[j]
			}
		}
	})
	b.Run("Assembly SSE", func(b *testing.B) {
		b.ResetTimer()

		u := soa.DefaultManager.Allocate([]float32{}).([]float32)
		v := soa.DefaultManager.Allocate([]float32{}).([]float32)
		w := soa.DefaultManager.Allocate([]float32{}).([]float32)

		for i := 0; i < len(u); i++ {
			u[i] = rand.Float32()
			v[i] = rand.Float32()
		}

		soa.AddFloat32s(u, v, w)
		for i := 0; i < b.N; i++ {
			soa.AddFloat32s(u, v, w)
		}
	})
}

func BenchmarkAddFloat64s(b *testing.B) {
	b.Run("Native", func(b *testing.B) {
		b.ResetTimer()

		u := soa.DefaultManager.Allocate([]float64{}).([]float64)
		v := soa.DefaultManager.Allocate([]float64{}).([]float64)
		w := soa.DefaultManager.Allocate([]float64{}).([]float64)

		for i := 0; i < len(u); i++ {
			u[i] = rand.Float64()
			v[i] = rand.Float64()
		}

		for j := 0; j < len(u); j++ {
			w[j] = u[j] + v[j]
		}
		for i := 0; i < b.N; i++ {
			for j := 0; j < len(u); j++ {
				w[j] = u[j] + v[j]
			}
		}
	})
	b.Run("Assembly SSE", func(b *testing.B) {
		b.ResetTimer()

		u := soa.DefaultManager.Allocate([]float64{}).([]float64)
		v := soa.DefaultManager.Allocate([]float64{}).([]float64)
		w := soa.DefaultManager.Allocate([]float64{}).([]float64)

		for i := 0; i < len(u); i++ {
			u[i] = rand.Float64()
			v[i] = rand.Float64()
		}

		soa.AddFloat64s(u, v, w)
		for i := 0; i < b.N; i++ {
			soa.AddFloat64s(u, v, w)
		}
	})
}
func BenchmarkAddInt32s(b *testing.B) {
	b.Run("Native", func(b *testing.B) {
		b.ResetTimer()

		u := soa.DefaultManager.Allocate([]int32{}).([]int32)
		v := soa.DefaultManager.Allocate([]int32{}).([]int32)
		w := soa.DefaultManager.Allocate([]int32{}).([]int32)

		for i := 0; i < len(u); i++ {
			u[i] = rand.Int31()
			v[i] = rand.Int31()
		}

		for j := 0; j < len(u); j++ {
			w[j] = u[j] + v[j]
		}
		for i := 0; i < b.N; i++ {
			for j := 0; j < len(u); j++ {
				w[j] = u[j] + v[j]
			}
		}
	})
	b.Run("Assembly SSE", func(b *testing.B) {
		b.ResetTimer()

		u := soa.DefaultManager.Allocate([]int32{}).([]int32)
		v := soa.DefaultManager.Allocate([]int32{}).([]int32)
		w := soa.DefaultManager.Allocate([]int32{}).([]int32)

		for i := 0; i < len(u); i++ {
			u[i] = rand.Int31()
			v[i] = rand.Int31()
		}

		soa.AddInt32s(u, v, w)
		for i := 0; i < b.N; i++ {
			soa.AddInt32s(u, v, w)
		}
	})
}

func BenchmarkAddInt64s(b *testing.B) {
	b.Run("Native", func(b *testing.B) {
		b.ResetTimer()

		u := soa.DefaultManager.Allocate([]int64{}).([]int64)
		v := soa.DefaultManager.Allocate([]int64{}).([]int64)
		w := soa.DefaultManager.Allocate([]int64{}).([]int64)

		for i := 0; i < len(u); i++ {
			u[i] = rand.Int63()
			v[i] = rand.Int63()
		}

		for j := 0; j < len(u); j++ {
			w[j] = u[j] + v[j]
		}
		for i := 0; i < b.N; i++ {
			for j := 0; j < len(u); j++ {
				w[j] = u[j] + v[j]
			}
		}
	})
	b.Run("Assembly SSE", func(b *testing.B) {
		b.ResetTimer()

		u := soa.DefaultManager.Allocate([]int64{}).([]int64)
		v := soa.DefaultManager.Allocate([]int64{}).([]int64)
		w := soa.DefaultManager.Allocate([]int64{}).([]int64)

		for i := 0; i < len(u); i++ {
			u[i] = rand.Int63()
			v[i] = rand.Int63()
		}

		soa.AddInt64s(u, v, w)
		for i := 0; i < b.N; i++ {
			soa.AddInt64s(u, v, w)
		}
	})
}
