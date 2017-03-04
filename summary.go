// Package summary は簡単な集計を行うためのGoのパッケージです。
// ベンチマークテストの結果をAddしていって最後に集計結果を表示するような用途です。
package summary

import (
	"fmt"
	"io"
	"math"
	"sort"
)

// Summary は集計結果を逐次格納する構造体です。
type Summary struct {
	Min    float64
	Max    float64
	Sum    float64
	Cnt    int64
	values []float64
}

// NewSummary は Summary構造体のコンストラクタです。
func NewSummary() *Summary {
	return &Summary{
		Min: math.Inf(+1),
		Max: math.Inf(-1),
	}
}

// Add は集計値を追加するメソッドです。
func (s *Summary) Add(val float64) {
	if s.Min > val {
		s.Min = val
	}
	if s.Max < val {
		s.Max = val
	}
	s.Cnt++
	s.Sum += val
	s.values = append(s.values, val)
}

// Median は現時点の中央値を返すメソッドです。
func (s *Summary) Median() float64 {
	if len(s.values) == 0 {
		return 0.0
	}
	sort.Float64s(s.values)
	return s.values[len(s.values)/2]
}

// Average は現時点の平均値を返すメソッドです。
func (s *Summary) Average() float64 {
	return s.Sum / float64(s.Cnt)
}

// PrintSummary は現時点の最小・最大・中央値・平均値をwにテキスト出力します。
func (s *Summary) PrintSummary(w io.Writer) {
	fmt.Fprintf(w, `
min:    %v
mean:   %v
median: %v
max:    %v
`,
		s.Min, s.Average(), s.Median(), s.Max)
}
