package summary

import (
	"os"
	"testing"
)

func TestSum(t *testing.T) {
	s := NewSummary()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	if s.Sum != 6 {
		t.Fatal("s.Sum should be 6, however doesn't match")
	}
}

func TestMedian(t *testing.T) {
	s := NewSummary()
	s.Add(1)
	s.Add(10)
	s.Add(100)
	if s.Median() != 10 {
		t.Fatal("s.Median() should be 10, however doesn't match")
	}
	s.Add(100)
	s.Add(1000)
	s.Add(1000)
	if s.Median() != 100 {
		t.Fatal("s.Median() should be 100, however doesn't match")
	}
}

func TestAverage(t *testing.T) {
	s := NewSummary()
	s.Add(1)
	s.Add(10)
	s.Add(100)
	if s.Average() != 37 {
		t.Fatal("s.Average() should be 37, however doesn't match")
	}
}

func TestMin(t *testing.T) {
	s := NewSummary()
	s.Add(100)
	s.Add(100000)
	s.Add(1)
	s.Add(10)
	s.Add(10000)
	if s.Min != 1 {
		t.Fatal("s.Min should be 1, however doesn't match")
	}
}

func TestMax(t *testing.T) {
	s := NewSummary()
	s.Add(100)
	s.Add(100000)
	s.Add(1)
	s.Add(10)
	s.Add(10000)
	if s.Max != 100000 {
		t.Fatal("s.Max should be 100000, however doesn't match")
	}
}

func ExampleSummary_PrintSummary() {
	s := NewSummary()
	s.Add(50)
	s.Add(0.5)
	s.Add(500)
	s.PrintSummary(os.Stdout)

	// Output:
	// min:    0.5
	// mean:   183.5
	// median: 50
	// max:    500
}
