package bytes

import (
	"fmt"
	"math"
)

const (
	Byte  = 1
	KByte = Byte * 1024
	MByte = KByte * 1024
	GByte = MByte * 1024
	TByte = GByte * 1024
	PByte = TByte * 1024
	EByte = PByte * 1024
)

func log(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func Format(v uint64, base float64, sizes []string) string {
	if v < 10 {
		return fmt.Sprintf("%d B", v)
	}
	e := math.Floor(log(float64(v), base))
	suffix := sizes[int(e)]
	val := float64(v) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+" %s", val, suffix)
}
