package winenv

import "strconv"

func (w *WindowsEnvInfo) compareValue() (v int) {
	v, _ = strconv.Atoi(w.Build)
	return
}
