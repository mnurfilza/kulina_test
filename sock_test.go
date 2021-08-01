package kulina

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSock(t *testing.T) {
	convey.Convey("test", t, func() {
		res := findPairColor(7, []int{1, 2, 1, 2, 1, 3, 2})
		fmt.Println(res)
	})
}
