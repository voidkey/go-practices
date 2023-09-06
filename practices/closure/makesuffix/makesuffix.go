package suffix

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func SuffixTest() {

	f := makeSuffix(".jpg")
	fmt.Println(f("spring"))
	fmt.Println(f("summer.jpg"))
	fmt.Println(f("autumn.png"))
	fmt.Println(f("winter."))

}
