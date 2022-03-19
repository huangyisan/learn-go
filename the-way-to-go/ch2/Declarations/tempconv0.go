package Declarations

type Celsius float64
type Fahrenhit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreeezingC    Celsius = 0
	Boiling       Celsius = 100
)

func CtoF(c Celsius) Fahrenhit {
	return Fahrenhit(c*9/5 + 32)
}
func FtoC(f Fahrenhit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
