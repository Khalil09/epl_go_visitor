package main

import "strconv"

// Declaração interface Visitor //
type Visitor interface {
    visitLiteral(lit *Literal)
    visitAdd(add *Add)
    visitSub(sub *Sub)
    visitMulti(multi *Multi)
}

// Declaração struct Eval //
type Eval struct {
    Visitor
    res int
}

func (e *Eval) visitLiteral(lit *Literal){
    e.res = lit.value
}
func (e *Eval) visitAdd(add *Add) {
    add.lhs.accept(e)
    a := e.res
    add.rhs.accept(e)
    b := e.res

    e.res = (a + b)
}

func (e *Eval) visitSub(sub *Sub) {
    sub.lhs.accept(e)
    a := e.res
    sub.rhs.accept(e)
    b := e.res

    e.res = (a - b)
}

func (e *Eval) visitMulti(multi *Multi) {
    multi.lhs.accept(e)
    a := e.res
    multi.rhs.accept(e)
    b := e.res

    e.res = (a * b)
}

func (e *Eval) result() int{
    return e.res
}
// Declaração interface Expression //

type Expression interface {
    print() string
    accept(v Visitor)
}

// Declaração Struct Literal //
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

// Declaração Struct Add //
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

// Declaração Struct Sub //
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

// Declaração struct multi
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