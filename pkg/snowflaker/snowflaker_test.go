package snowflaker

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNextID(t *testing.T) {
	err := Init("2020-12-12", 12)
	if err != nil {
		log.Fatal()
	}
	a := NextID()
	b := NextID()
	assert.NotEqual(t, a, b)
}

func BenchmarkSelect(b *testing.B) {
	err := Init("2020-12-12", 12)
	if err != nil {
		log.Fatal()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NextID()
	}
}
