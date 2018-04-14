package main

type Sub struct {
    Expression
    lhs Expression
    rhs Expression
}

func (sub *Sub) print() string {
    return sub.lhs.print() + " - " + sub.rhs.print()
}

func (sub *Sub) accept(v Visitor) {
    v.visitSub(sub)
}