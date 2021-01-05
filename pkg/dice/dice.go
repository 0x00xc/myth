/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/8/26 15:32
 */
package dice

import (
	"fmt"
	"myth/pkg/rand"
	"strings"
)

var polish *PolishNotation

func init() {
	polish = NewPolishNotation()
	var roll = func(m, n float64) float64 {
		v := 0
		for i := 0; i < int(m); i++ {
			v += rand.Intn(int(n)) + 1
		}
		return float64(v)
	}
	polish.SetOperator('d', roll, 3)
	polish.SetOperator('D', roll, 3)
}

func SafeRoll(dice string) (int, error) {
	v, err := polish.Calculate(strings.ToLower(dice))
	return int(v), err
}

func Roll(dice string) int {
	v, err := SafeRoll(dice)
	if err != nil {
		panic(err)
	}
	return v
}

func D100() int {
	return Roll("1D100")
}

func ND100(n int) int {
	return Roll(fmt.Sprintf("%dD100", n))
}

func MDN(m, n int) int {
	return Roll(fmt.Sprintf("%dD%d", m, n))
}
