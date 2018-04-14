package main

import "strconv"

type Literal struct {
    Expression
    value int
}

func (l *Literal) print() string {
    return strconv.Itoa(l.value)

}
func (l *Literal) accept(v Visitor) {
    v.visitLiteral(l)
}