package goPuppy

import (
	"github.com/dogeyboy1932/goDog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return goDog.WhenGrownUp("Woof!")
}
