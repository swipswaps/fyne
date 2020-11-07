package binding

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
)

type floatToString struct {
	base

	from Float
}

func FloatToString(f Float) String {
	str := &floatToString{from: f}
	f.AddListener(str)
	return str
}

func (f *floatToString) Get() string {
	val := f.from.Get()

	return fmt.Sprintf("%0.2f", val) // TODO format string
}

func (f *floatToString) Set(val string) {
	fVal, err := strconv.ParseFloat(val, 0)
	if err != nil {
		fyne.LogError("Float parse error", err)
		return
	}
	if fVal == f.from.Get() {
		return
	}
	f.from.Set(fVal)

	f.trigger(f)
}

func (f *floatToString) DataChanged(_ DataItem) {
	f.trigger(f)
}
