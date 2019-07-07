package main

import (
    "syscall/js"
    "strconv"
)

func main() {
    c := make(chan struct{}, 0)

    println("WASM Go Initialized")
    // register functions
    registerCallbacks()
    <-c
}

func registerCallbacks() {
    js.Global().Set("add", js.NewCallback(add))
    js.Global().Set("subtract", js.NewCallback(subtract))
}

func add(i []js.Value) {
    val1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
    val2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
    i1, _ := strconv.Atoi(val1)
    i2, _ := strconv.Atoi(val2)
    a1 := i1 + i2
    println(a1)
    js.Global().Get("document").Call("getElementById", i[2].String()).Set("innerHTML", a1)
}

func subtract(i []js.Value) {
    val3 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
    val4 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
    i3, _ := strconv.Atoi(val3)
    i4, _ := strconv.Atoi(val4)
    a2 := i3 - i4
    println(a2)
    js.Global().Get("document").Call("getElementById", i[2].String()).Set("innerHTML", a2)
}