package example

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expect be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + (-20) expect be -30, but %d got", ans)
	}
}

func TestMul(t *testing.T) {
	if ans := Mul(1, 2); ans != 2 {
		t.Errorf("1 * 2 expect 2, but %d got", ans)
	}

	if ans := Mul(-1, -2); ans != 2 {
		t.Errorf("-1 * -2 expect 2, but %d got", ans)
	}
}

// 子测试
func TestAdd1(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if ans := Add(3, 4); ans != 7 {
			t.Fatal("fail")
		}
	})
	t.Run("neg", func(t *testing.T) {
		if ans := Add(-1, -5); ans != -6 {
			t.Fatal("fail")
		}
	})
}

func TestMul2(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 4, 8},
		{"neg", -2, -5, 10},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
			}
		})
	}
}

type calcCase struct {
	Name           string
	A, B, Expected int
}

func createMulTestCase(t *testing.T, c *calcCase) {
	// 可以打印出具体出问题的具体函数
	t.Helper()
	t.Run(c.Name, func(t *testing.T) {
		if ans := Mul(c.A, c.B); ans != c.Expected {
			t.Fatalf("%d * %d expected %d, but %d got",
				c.A, c.B, c.Expected, ans)
		}
	})
}

func TestMul3(t *testing.T) {
	createMulTestCase(t, &calcCase{"test1", 3, 6, 18})
	createMulTestCase(t, &calcCase{"test2", -3, -6, 18})
	createMulTestCase(t, &calcCase{"test3", 0, 1, 9}) // wrong case
}

func BenchmarkHello(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello %d", i)
	}
}
