package main

type Add struct {
    Expression
    lhs Expression
    rhs Expression
}

func (add *Add) print() string {
    return add.lhs.print() + " + " + add.rhs.print()
}

func (add *Add) accept(v Visitor) {
    v.visitAdd(add)
}