package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	var nums []int
	var idxs []int
	r := bufio.NewScanner(f)
	for r.Scan() {
		nums = append(nums, optimistic.Atoi(r.Text()))
		idxs = append(idxs, len(idxs))
	}

	for i := 0; i < len(idxs); i++ {
		k := 0
		for kk, v := range idxs {
			if v == i {
				k = kk
				break
			}
		}
		num := nums[k]
		newK := (k + num) % (len(nums) - 1)
		if newK < 0 {
			newK = len(nums) + newK - 1
		}

		if k != newK {
			nums = append(nums[:k], nums[k+1:]...)
			nums = append(nums[:newK], append([]int{num}, nums[newK:]...)...)
			idxs = append(idxs[:k], idxs[k+1:]...)
			idxs = append(idxs[:newK], append([]int{i}, idxs[newK:]...)...)
		}
		//fmt.Println(nums)
	}

	zeroI := 0
	for i, v := range nums {
		if v == 0 {
			zeroI = i
			break
		}
	}

	sum := 0
	for i := 1; i <= 3000; i++ {
		if i == 1000 || i == 2000 || i == 3000 {
			num := nums[(zeroI+i)%len(nums)]
			fmt.Println(num)
			sum += num
		}
	}
	fmt.Println("sum", sum)
}
