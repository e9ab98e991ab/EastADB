// +build !windows

package vcl

import (
	"fmt"
)

func showError(err interface{}) {
	fmt.Println(err)
}

func tryLoadAppIcon() {
	// no code
}
