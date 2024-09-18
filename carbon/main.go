package main

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
)

func main() {
	carbon.SetDefault(carbon.Default{
		Layout:       carbon.DateTimeLayout,
		Timezone:     carbon.Local,
		WeekStartsAt: carbon.Sunday,
		Locale:       "zh_cn",
	})
	fmt.Printf("%s\n", carbon.Now())
	fmt.Println(carbon.Now().String())
}
