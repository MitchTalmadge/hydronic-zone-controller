package gpio

import (
	"github.com/stianeikeland/go-rpio/v4"
)

func Init() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	initRelays()
}
