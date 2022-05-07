package main

import (
	"fmt"
	"github.com/rocketlaunchr/dataframe-go"
)

func main() {
	s1 := dataframe.NewSeriesInt64("day", nil, 1, 2, 3, 4, 5, 6, 7, 8)
	s2 := dataframe.NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2, nil, nil, 84.2, 72, 89)
	df := dataframe.NewDataFrame(s1, s2)


	fmt.Print(df.Table())
}