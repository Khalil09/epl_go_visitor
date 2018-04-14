package main

type Multi struct {
    Expression
    lhs Expression
    rhs Expression
}

func (multi *Multi) print() string {
    return multi.lhs.print() + " * " + multi.rhs.print()
}

func (multi *Multi) accept(v Visitor) {
    v.visitMulti(multi)
}