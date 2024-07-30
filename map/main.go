package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02 15:04"
	loc, err := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(layout, "2024-07-30 16:37", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Unix())
	return
	fmt.Println(fmt.Sprintf("  %x", 455))
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	// mj, _ := json.Marshal(m1)
	/*for i := 0; i <= 20; i++ {
		j := i
		// m22 := make(map[string]int)
		// json.Unmarshal(mj, m22)
		go func() {
			time.Sleep(time.Second)
			m1["a"] = j
			fmt.Println(m1["a"])
		}()
	}
	time.Sleep(time.Second * 15)
	fmt.Println(m1)*/
	m2 := make(map[string]int)
	m2 = m1
	m2["a"] = 333
	fmt.Println(m1, m2)
	for idx, v := range m1 {
		m1[idx+"12"] = v + 1
	}
	for v := range m1 {
		fmt.Println(v)
	}
	fmt.Println(m1)
}
