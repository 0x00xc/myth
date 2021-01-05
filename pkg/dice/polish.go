/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/8/26 15:32
 */
package dice

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type PolishNotation struct {
	funcMap       map[string]func(m, n float64) float64
	operatorLevel map[byte]int
}

func NewPolishNotation() *PolishNotation {
	pn := new(PolishNotation)
	pn.funcMap = map[string]func(m, n float64) float64{
		"+": func(m, n float64) float64 { return m + n },
		"-": func(m, n float64) float64 { return m - n },
		"*": func(m, n float64) float64 { return m * n },
		"x": func(m, n float64) float64 { return m * n },
		"/": func(m, n float64) float64 { return m / n },
	}
	pn.operatorLevel = map[byte]int{'+': 1, '-': 1, '*': 2, '/': 2}
	return pn
}

//自定义运算符
// operator 运算符（不能为'e'、'E'、'('、')'、'.'等特殊字符（科学计数，小数点和括号））
// function 运算方法
// level    运算符优先级，（加减法优先级为1，乘除法优先级为2）
func (p *PolishNotation) SetOperator(operator byte, function func(m, n float64) float64, level int) {
	if operator == 'e' || operator == 'E' || operator == '.' || operator == '(' || operator == ')' {
		panic("invalid operator: " + string([]byte{operator}))
	}
	p.funcMap[string([]byte{operator})] = function
	p.operatorLevel[operator] = level
}

//计算波兰式
//不支持正负（+-）符号，如果需要使用正负符号，请使用0±n代替
//eg: -5*3 需要替换为 (0-5)*3
func (p *PolishNotation) Calculate(s string) (float64, error) {
	if s == "" {
		return 0, errors.New("invalid expression")
	}
	temp := p.convert(s)
	stack := make([]string, 0)
	for i := len(temp) - 1; i >= 0; i-- {
		v := temp[i]
		_, err := strconv.ParseFloat(v, 64)
		if err != nil {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			n1, _ := strconv.ParseFloat(stack[len(stack)-1], 64)
			n2, _ := strconv.ParseFloat(stack[len(stack)-2], 64)
			stack = stack[:len(stack)-2]
			if p.funcMap[v] == nil {
				return 0, errors.New("invalid expression")
			}
			var result = p.funcMap[v](n1, n2)
			stack = append(stack, fmt.Sprintf("%v", result))
		} else {
			stack = append(stack, v)
		}
	}
	res, _ := strconv.ParseFloat(stack[0], 64)
	return res, nil
}

//转为波兰式
func (p *PolishNotation) convert(s string) []string {
	s = strings.Replace(s, " ", "", -1)
	var s1 = make([]byte, 0)
	var s2 = make([]string, 0)
	n := ""
	for i := len(s) - 1; i >= 0; i-- {
		v := s[i]
		if (v >= '0' && v <= '9') || (v == 'e') || (v == 'E') || (v == '.') {
			n = string([]byte{v}) + n
		} else if v == '(' {
			if n != "" {
				s2 = append(s2, n)
				n = ""
			}
			for {
				last := s1[len(s1)-1]
				s1 = s1[:len(s1)-1]
				if last == ')' {
					break
				}
				s2 = append(s2, string([]byte{last}))
			}
		} else if v == ')' {
			if n != "" {
				s2 = append(s2, n)
				n = ""
			}
			s1 = append(s1, v)
		} else {
			if n != "" {
				s2 = append(s2, n)
				n = ""
			}

			for {
				if len(s1) == 0 || s1[len(s1)-1] == ')' {
					s1 = append(s1, v)
					break
				}
				last := s1[len(s1)-1]
				if p.operatorLevel[v] >= p.operatorLevel[last] {
					s1 = append(s1, v)
					break
				}
				s2 = append(s2, string([]byte{last}))
				s1 = s1[:len(s1)-1]
			}
		}
	}
	if n != "" {
		s2 = append(s2, n)
		n = ""
	}

	for i := len(s1) - 1; i >= 0; i-- {
		s2 = append(s2, string([]byte{s1[i]}))
	}

	res := make([]string, 0)
	for i := len(s2) - 1; i >= 0; i-- {
		res = append(res, s2[i])
	}
	return res
}
