package external

import "fmt"

type PrintAPIImpl struct {
}

func (p *PrintAPIImpl) PrintSomething() {
	fmt.Println("external service")
}
