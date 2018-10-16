// h 20181016
//
// DrvLot.SHA

package sha

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println(sha([]byte("hello world")))
	// 47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad
}
