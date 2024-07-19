package time

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func Time() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err.Error())
		if err != nil {
			return
		}
		os.Exit(1)
	}
	fmt.Println(time.String())
}
