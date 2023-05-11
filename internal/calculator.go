package internal

import "github.com/dgff07/testcontainers-go/external"

type PrintAPI interface {
	PrintSomething()
}

type Calculator struct {
	printAPI PrintAPI
}

func InitCalculator() *Calculator {
	return &Calculator{
		printAPI: &external.PrintAPIImpl{},
	}
}

func (c *Calculator) Sum(v1, v2 int) int {
	c.printAPI.PrintSomething()
	return v1 + v2
}

func (c *Calculator) Sub(v1, v2 int) int {
	return v1 - v2
}
