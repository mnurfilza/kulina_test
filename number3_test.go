package kulina

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestNumber3(t *testing.T) {
	convey.Convey("test", t, func() {
		number("1.345.679")
		// fmt.Println(res)
	})
}
