package main

func main() {
	var m map[int]bool
	// m[123] = false // 声明没有分配内存空间
	_ = m[123]

	var p *[5]string
	for range p {
		_ = len(p)
	}

	var s []int
	// s[0] = 1 声明没有分配内存空间
	_ = s[:]
	s, s[0] = []int{1, 2}, 9 // 先左边 在右边
}
