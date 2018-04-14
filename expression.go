package main

type Expression interface {
    print() string
    accept(v Visitor)
}