package statements

import (
	"math/rand/v2"
	"testing"
)

func TestLimitStatement(t *testing.T) {
	t.Run("TestLimitWithOutOffset", func(t *testing.T) {
		limit := int64(rand.Float32() * 10000)

		statement := &Limit{
			Limit: limit,
		}

		expected := "LIMIT ?"
		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected limit statement to be (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestLimitWithOffset", func(t *testing.T) {
		limit := int64(rand.Float32() * 10000)
		offset := int64(rand.Float32() * 100)

		statement := &Limit{
			Limit:  limit,
			Offset: offset,
		}

		expected := "LIMIT ? OFFSET ?"
		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected limit statement to be (%s) but got (%s)", expected, actual)
		}
	})
}
