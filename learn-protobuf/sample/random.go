package sample

import (
	"learn-protobuf/pb"
	"math/rand"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_UNKNOWN
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUbrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUname(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xen E-2286M",
			"Core i9-2286M",
			"Core i7-2286M",
		)
	}
	return randomStringFromSet(
		"Ryzen 7 Pro 2780U",
		"Ryzen 5 Pro 2780U",
		"Ryzen 3 Pro 2780U",
	)
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomInt(min, max int) int {
	// return 0 - (max-min)

	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	// rand.Float64取值范围是0-1，所以这边乘以大小之差再加上小，则得到了最终范围
	return min + rand.Float64()*(max-min)
}
