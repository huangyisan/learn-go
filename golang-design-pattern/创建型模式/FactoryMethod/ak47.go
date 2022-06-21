package main

// ak47具体产品
type ak47 struct {
	// 继承gun属性
	gun
}

// 返回iGun
func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "ak47",
			power: 4,
		},
	}
}
