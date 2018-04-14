package main

type Visitor interface {
    visitLiteral(lit *Literal)
    visitAdd(add *Add)
    visitSub(sub *Sub)
    visitMulti(multi *Multi)
}