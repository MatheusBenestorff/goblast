package main

import (
	"github.com/MatheusBenestorff/goblast/internal/loadtester"
)

func main() {
	loadtester.RunLoadTest("http://localhost:8080", 100)
}