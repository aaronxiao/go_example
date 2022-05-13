package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成图标
func GenerateIcons(iconsCfg [][]uint16, row int, injectIcons ...byte) [][]byte {
	if len(iconsCfg) == 0 { // 没有设置就直接返回
		return nil
	}

	fmt.Println(injectIcons)

	var start = time.Now()
	var col = len(iconsCfg)
	var res = make([][]byte, row)
	for k, v := range iconsCfg { // col number
		tape := GetRandTape(v, row, injectIcons...)
		l := len(tape)
		x := rand.Intn(l)

		for i := 0; i < row; i++ { // row
			if len(res[i]) == 0 {
				res[i] = make([]byte, col)
			}

			res[i][k] = tape[(x+i)%l]
		}
	}

	fmt.Printf(fmt.Sprintf("generate icons duration %s \n", time.Since(start)))
	return res
}


// 获取随机带
func GetRandTape(colSlice []uint16, row int, injectIcons ...byte) []byte {
	fmt.Printf(fmt.Sprintf("col slice %v row %d inject icons %v \n", colSlice, row, injectIcons))
	var length int
	for _, v := range colSlice {
		length += int(v)
	}

	var res = make([]byte, length)
	var idx int
	for i := range colSlice {
	LABEL:
		for j := 0; uint16(j) < colSlice[i]; j++ {
			for _, v := range injectIcons { // 过滤特殊图标
				if byte(i) == v {
					continue LABEL
				}
			}

			res[idx] = byte(i)
			idx++
		}
	}

	// 打乱
	rand.Shuffle(idx, func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})

	// 插入特殊图标
	var totalLen = idx
	for _, v := range injectIcons {
		totalLen = injectSpecialIcon(colSlice, res, totalLen, row, v)
	}

	return res
}

// 插入特殊图标
func injectSpecialIcon(colSlice []uint16, tape []byte, length, row int, icon byte) int {
	if int(icon) >= len(colSlice) { // 超出
		return length
	}

	count := int(colSlice[icon])
	if count == 0 { // 图标不存在
		return length
	}

	length += count
	if length/row < count { // 不够平均
		fmt.Printf(fmt.Sprintf("icon %d length %d row %d count %d \n", icon, length, row, count))
		return length
	}

	positions := rand.Perm(length / row)[0:count]
	fmt.Printf("sssssssssssssssssssss:%v \n", positions)
	for i := range positions {
		positions[i] *= row
	}
	fmt.Printf("wwwwwwwwwwwwwwwwwwwww:%v \n", positions)

	// 交换值
	for i := 0; i < count; i++ {
		tape[length-i-1] = icon
		tape[positions[i]], tape[length-i-1] = tape[length-i-1], tape[positions[i]]
	}

	return length
}


func InjectIcons() []byte {
	return []byte{11}
}


func main() {
	rand.Seed(time.Now().Unix())
	var cfg = [][]uint16{
		{30, 40, 40, 50, 50, 60, 60, 60, 60, 60, 5, 3},
		{30, 40, 40, 50, 50, 60, 60, 60, 60, 60, 3, 0},
		{30, 30, 40, 50, 50, 60, 60, 60, 60, 60, 5, 5},
		{30, 40, 50, 50, 50, 60, 60, 60, 60, 60, 10, 0},
		{30, 40, 50, 50, 50, 50, 60, 60, 60, 60, 3, 5},
	}

	GenerateIcons(cfg, 3, InjectIcons()... )

}