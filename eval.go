package main

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