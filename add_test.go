package soa_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/andygeiss/assert"
	"github.com/andygeiss/soa"
)

func TestAdd(t *testing.T) {
	a := soa.DefaultManager.Allocate([]float64{}).([]float64)
	b := soa.DefaultManager.Allocate([]float64{}).([]float64)
	c := soa.DefaultManager.Allocate([]float64{}).([]float64)

	for i := 0; i < len(a); i++ {
		a[i] = 1.0
		b[i] = 8.0
	}

	soa.AddFloat64s(a, b, c)
	fmt.Printf("%v", c)

	assert.That("len c should be 512", t, len(c), 512)
	assert.That("c[0] should be 9.0", t, c[0], 9.0)
	assert.That("c[0] should be 9.0", t, c[511], 9.0)
}

func BenchmarkAdd(b *testing.B) {
	b.Run("Native", func(b *testing.B) {

		u := soa.DefaultManager.Allocate([]float64{}).([]float64)
		v := soa.DefaultManager.Allocate([]float64{}).([]float64)
		w := soa.DefaultManager.Allocate([]float64{}).([]float64)

		for i := 0; i < len(u); i++ {
			u[i] = rand.Float64()
			v[i] = rand.Float64()
		}

		for i := 0; i < b.N; i++ {
			for j := 0; j < len(u); j++ {
				w[j] = u[j] + v[j]
			}
		}
	})
	b.Run("Assembly SSE", func(b *testing.B) {

		u := soa.DefaultManager.Allocate([]float64{}).([]float64)
		v := soa.DefaultManager.Allocate([]float64{}).([]float64)
		w := soa.DefaultManager.Allocate([]float64{}).([]float64)

		for i := 0; i < len(u); i++ {
			u[i] = rand.Float64()
			v[i] = rand.Float64()
		}

		for i := 0; i < b.N; i++ {
			soa.AddFloat64s(u, v, w)
		}
	})
}
