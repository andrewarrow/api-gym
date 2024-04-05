package browser

import "time"

func flashThree(s string) {
	flash := Document.ById("flash")
	flash.Set("innerHTML", s)
	time.Sleep(time.Second * 3)
	flash.Set("innerHTML", "")
}
