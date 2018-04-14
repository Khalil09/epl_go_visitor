package main

import "testing"

var literal100 Literal = Literal{ value: 100 }
var literal300 Literal = Literal{ value: 300 }

var add Expression = &Add{lhs: &literal300, rhs: &literal100}
var sub Expression = &Sub{lhs: &literal300, rhs: &literal100}
var multi Expression = &Multi{lhs: &literal300, rhs: &literal100}
var eval = new(Eval)

func TestAdd(t *testing.T) {
    result := add.print()

    if (result != "300 + 100") {
        t.Errorf("Erro on add, got: %s, want: %s.", result, "300 + 100")
    }
}

func TestSub(t *testing.T) {
    result := sub.print()

    if (result != "300 - 100") {
        t.Errorf("Erro on sub, got: %s, want: %s.", result, "300 - 100")
    }
}

func TestMulti(t *testing.T) {
    result := multi.print()

    if (result != "300 * 100") {
        t.Errorf("Erro multi, got: %s, want: %s.", result, "300 * 100")
    }
}

func TestVisitLiteral(t *testing.T) {

    eval.visitLiteral(&literal100)
    result := eval.result()

    if(result != literal100.value) {
        t.Errorf("Erro visitor of literal, got: %d, want: %d.", result, literal100.value)
    }
}

func TestVisitAdd(t *testing.T) {

    add.accept(eval)
    result := eval.result()

    if(result != literal100.value + literal300.value) {
        t.Errorf("Erro visitor of add or accept, got: %d, want: %d.", result, literal100.value + literal300.value)
    }
}

func TestVisitSub(t *testing.T) {

    sub.accept(eval)
    result := eval.result()

    if(result != literal300.value - literal100.value) {
        t.Errorf("Erro visitor of sub or accept, got: %d, want: %d.", result, literal300.value - literal100.value)
    }
}

func TestVisitMulti(t *testing.T) {

    multi.accept(eval)
    result := eval.result()

    if(result != literal300.value * literal100.value) {
        t.Errorf("Erro visitor of multi or accept, got: %d, want: %d.", result, literal300.value * literal100.value)
    }
}








