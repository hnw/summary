# summary [![Travis CI build status](https://travis-ci.org/hnw/summary.svg?branch=master)](https://travis-ci.org/hnw/summary)

``` go
    import "github.com/hnw/summary"
```

Package summary は簡単な集計を行うためのGoのパッケージです。
ベンチマークテストの結果をAddしていって最後に集計結果を表示するような用途です。

## Usage

#### type Summary

```go
type Summary struct {
	Min float64
	Max float64
	Sum float64
	Cnt int64
}
```

Summary は集計結果を逐次格納する構造体です。

#### func  NewSummary

```go
func NewSummary() *Summary
```
NewSummary は Summary構造体のコンストラクタです。

#### func (*Summary) Add

```go
func (s *Summary) Add(val float64)
```
Add は集計値を追加するメソッドです。

#### func (*Summary) Average

```go
func (s *Summary) Average() float64
```
Average は現時点の平均値を返すメソッドです。

#### func (*Summary) Median

```go
func (s *Summary) Median() float64
```
Median は現時点の中央値を返すメソッドです。

#### func (*Summary) PrintSummary

```go
func (s *Summary) PrintSummary(w io.Writer)
```
PrintSummary は現時点の最小・最大・中央値・平均値をwにテキスト出力します。
