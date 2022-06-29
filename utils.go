package main

import (
	"encoding/json"
	"math/rand"
)

func If[T any](cond bool, vtrue T, vfalse T) T {
    if cond {
        return vtrue
    }
    return vfalse
}

func RandomIntBetween(min int, max int) int {
	return rand.Intn(max - min) + min
}

func Map[T AnyT](arr []T, fn func(T) T) []T {
	result := []T{}
    // result := make([]T, len(arr))
    for i, e := range arr {
        result[i] = fn(e)
	}

    return result
}

func Filter[T AnyT](arr []T, cond func(fn T) bool) []T {
	// result := []func(value T){}
	result := []T{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}

func Stringify(obj any) string {
	jsonValue, _ := json.Marshal(obj)
	return string(jsonValue)
}

func constrain[T AnyNumber](number T, min T, max T) T {
	if number < min {
		return min
	} 
	if number > max {
		return max
	}
	return number
}

func IX(x int, y int, width int) int {
	// TODO remplacer constrain par un seuil/boucle pour fair un espace donut ???
	x = constrain(x, 0, width-1)
	y = constrain(y, 0, width-1)
	return x + (y * width)
}

func SWAP(px, py *[]float64) {
	tempx := *px
	tempy := *py
	*px = tempy
	*py = tempx
}
