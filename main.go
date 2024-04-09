package main

import (
	"fmt"
	"strconv"
	"strings"
)

func evaluateExpression(expr string) int {
	stack := make([]int, 0)
	var num int
	var op byte = '+'

	for i := 0; i < len(expr); i++ {
		if expr[i] >= '0' && expr[i] <= '9' {
			num = num*10 + int(expr[i]-'0')
		}
		if (expr[i] < '0' && expr[i] != ' ') || i == len(expr)-1 {
			if op == '+' {
				stack = append(stack, num)
			} else if op == '-' {
				stack = append(stack, -num)
			}
			op = expr[i]
			num = 0
		}
	}

	result := 0
	for _, val := range stack {
		result += val
	}
	return result
}

func cleanExpression(expr string) string {
	expr = strings.TrimLeft(expr, "+")
	if len(expr) > 0 && expr[len(expr)-1] == '0' && (len(expr) == 1 || (expr[len(expr)-2] == '+' || expr[len(expr)-2] == '-')) {
		return expr[:len(expr)-2]
	}
	return expr
}

func generateExpressions(curr string, num int, target int, results map[string]struct{}) {
	if num == -1 {
		if evaluateExpression(curr) == target {
			curr = cleanExpression(curr)
			results[curr] = struct{}{}
		}
		return
	}

	generateExpressions(curr+"+"+strconv.Itoa(num), num-1, target, results)
	generateExpressions(curr+"-"+strconv.Itoa(num), num-1, target, results)
	generateExpressions(curr+strconv.Itoa(num), num-1, target, results)
}

func main() {
	target := 200
	results := make(map[string]struct{})
	generateExpressions("", 9, target, results)
	for expr := range results {
		if expr != "" {
			fmt.Println(expr)
		}
	}
}
