```go
package main

import (
  "fmt"
    "math"
      "sort"
      )

      func main() {
        // 创建一个包含随机数的切片
          numbers := []float64{3.5, 2.1, 5.6, 4.8, 1.2}
            // 使用sort包对切片进行排序
              sort.Float64s(numbers)
                // 打印排序后的切片
                  fmt.Println("Sorted numbers:", numbers)
                    // 计算平均值
                      average := sum(numbers) / float64(len(numbers))
                        fmt.Println("Average:", average)
                          // 计算标准差
                            variance := sumOfSquares(numbers, average) / float64(len(numbers))
                              standardDeviation := math.Sqrt(variance)
                                fmt.Println("Standard Deviation:", standardDeviation)
                                }

                                // sum计算切片中所有元素的和
                                func sum(numbers []float64) float64 {
                                  total := 0.0
                                    for _, number := range numbers {
                                        total += number
                                          }
                                            return total
                                            }

                                            // sumOfSquares计算每个元素与平均值的平方差的和
                                            func sumOfSquares(numbers []float64, average float64) float64 {
                                              var sumOfSquares float64
                                                for _, number := range numbers {
                                                    sumOfSquares += (number - average) * (number - average)
                                                      }
                                                        return sumOfSquares
                                                        }
                                                        ```
