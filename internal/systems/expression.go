package systems

import (
	"github.com/faiface/pixel"
	"numberjungle/internal/data"
	"numberjungle/internal/vars"
)

func ExprPosSystem() {
	currWidth := 0
	currLine := 0
	var currRow []*data.Coin
	var lastCoin *data.Coin
	for _, coin := range data.Expression.Coins {
		coin.Sprs[0].Key = "coin"
		if currWidth == vars.MaxExprWidth || (lastCoin != nil && lastCoin.Value.IsInt() && !coin.Value.IsInt()) {
			y := vars.ExprYStart - 28. * float64(currLine)
			for i, c := range currRow {
				x := vars.ExprXStart - 24. * float64(len(currRow) - i - 1)
				c.Object.Pos = pixel.V(x, y)
			}
			currRow = []*data.Coin{}
			currWidth = 0
			currLine++
		}
		currRow = append(currRow, coin)
		lastCoin = coin
		currWidth++
	}
	y := vars.ExprYStart - 28. * float64(currLine)
	for i, c := range currRow {
		x := vars.ExprXStart - 24. * float64(len(currRow) - i - 1)
		c.Object.Pos = pixel.V(x, y)
	}
}

//goland:noinspection GoNilness
func ExprCheckSystem() {
	fail := false
	sides, eqs := splitOnEq(data.Expression.Coins)
	if len(eqs) < 1 {
		fail = true
	}
	var totals []int
	var errs []bool
	for _, expr := range sides {
		if len(expr) > 0 {
			newExpr := combine(expr)
			total, err := solve(newExpr)
			totals = append(totals, total)
			errs = append(errs, err)
		} else {
			totals = append(totals, 0)
			errs = append(errs, true)
		}
	}
	for i, expr := range sides {
		if len(expr) == 0 {
			if i-1 > -1 && i-1 < len(eqs) {
				eqs[i-1].Sprs[0].Key = "bad_coin"
			}
			if i > -1 && i < len(eqs) {
				eqs[i].Sprs[0].Key = "bad_coin"
			}
			fail = true
		} else {
			if i-1 > -1 && i-1 < len(sides) && i-1 < len(eqs) {
				if totals[i-1] != totals[i] {
					eqs[i-1].Sprs[0].Key = "bad_coin"
					fail = true
				}
			}
			if i+1 > -1 && i+1 < len(sides) && i < len(eqs) {
				if totals[i+1] != totals[i] {
					eqs[i].Sprs[0].Key = "bad_coin"
					fail = true
				}
			}
		}
		if errs[i] {
			fail = true
		}
	}
	data.CurrPuzzle.Done = !fail
	if !data.CurrPuzzle.Done {
		for _, coin := range data.Expression.Coins {
			if coin.Value == data.Equals {
				coin.Sprs[0].Key = "bad_coin"
			}
		}
	}
}

func combine(expr []*data.Coin) []*node {
	var newExpr []*node
	intVal := 0
	anInt := false
	for _, coin := range expr {
		if coin.Value.IsInt() {
			anInt = true
			intVal *= 10
			intVal += coin.Value.Int()
		} else {
			if anInt {
				newExpr = append(newExpr, &node{
					coin: nil,
					num:  intVal,
				})
			}
			intVal = 0
			anInt = false
			newExpr = append(newExpr, &node{
				coin: coin,
				num:  0,
			})
		}
	}
	if anInt {
		newExpr = append(newExpr, &node{
			coin: nil,
			num:  intVal,
		})
	}
	return newExpr
}

func splitOnEq(a []*data.Coin) ([][]*data.Coin, []*data.Coin) {
	var result [][]*data.Coin
	var eqs []*data.Coin
	var side []*data.Coin
	for _, r := range a {
		if r.Value == data.Equals {
			eqs = append(eqs, r)
			result = append(result, side)
			side = []*data.Coin{}
		} else {
			side = append(side, r)
		}
	}
	result = append(result, side)
	return result, eqs
}

func solve(nodes []*node) (int, bool) {
	fail := false
	var operators []*data.Coin
	var operands []int
	needOperand := true
	negate := false
	var unary *data.Coin
	for _, n := range nodes {
		// todo: add parenthesis
		if n.coin == nil {
			// if an operand appears
			if !needOperand {
				panic("double operands found!")
			}
			if negate {
				operands = append(operands, -n.num)
			} else {
				operands = append(operands, n.num)
			}
			needOperand = false
			unary = nil
		} else if needOperand {
			if n.coin.Value == data.Minus && !negate {
				negate = true
				unary = n.coin
			} else {
				if unary != nil {
					unary.Sprs[0].Key = "bad_coin"
					unary = nil
				}
				n.coin.Sprs[0].Key = "bad_coin"
				fail = true
			}
			continue
		} else {
			if unary != nil {
				unary.Sprs[0].Key = "bad_coin"
				unary = nil
			}
			for len(operators) > 0 && operators[len(operators)-1].Value.Priority() >= n.coin.Value.Priority() {
				top := operators[len(operators)-1]
				operators = operators[:len(operators)-1]
				if len(operands) > 1 {
					//goland:noinspection GoNilness
					o2 := operands[len(operands)-1]
					operands = operands[:len(operands)-1]
					o1 := operands[len(operands)-1]
					operands = operands[:len(operands)-1]
					res := 0
					switch top.Value {
					case data.Plus:
						res = o1 + o2
					case data.Minus:
						res = o1 - o2
					case data.Times:
						res = o1 * o2
					case data.Divide:
						if o2 == 0 {
							top.Sprs[0].Key = "bad_coin"
							fail = true
						} else {
							if o1%o2 != 0 {
								top.Sprs[0].Key = "bad_coin"
								fail = true
							}
							res = o1 / o2
						}
					}
					operands = append(operands, res)
				} else {
					top.Sprs[0].Key = "bad_coin"
					fail = true
					continue
				}
			}
			operators = append(operators, n.coin)
			needOperand = true
		}
	}
	for len(operators) > 0 {
		top := operators[len(operators)-1]
		operators = operators[:len(operators)-1]
		if len(operands) > 1 {
			//goland:noinspection GoNilness
			o2 := operands[len(operands)-1]
			operands = operands[:len(operands)-1]
			o1 := operands[len(operands)-1]
			operands = operands[:len(operands)-1]
			res := 0
			switch top.Value {
			case data.Plus:
				res = o1 + o2
			case data.Minus:
				res = o1 - o2
			case data.Times:
				res = o1 * o2
			case data.Divide:
				if o1 % o2 != 0 {
					top.Sprs[0].Key = "bad_coin"
					fail = true
				}
				res = o1 / o2
			}
			operands = append(operands, res)
		} else {
			top.Sprs[0].Key = "bad_coin"
			fail = true
			continue
		}
	}
	if unary != nil {
		unary.Sprs[0].Key = "bad_coin"
		unary = nil
	}
	if len(operands) == 1 {
		return operands[0], fail
	}
	return 0, true
}

type node struct {
	coin *data.Coin
	num  int
}

type bTree struct {
	left *bTree
	node *node
	right *bTree
}