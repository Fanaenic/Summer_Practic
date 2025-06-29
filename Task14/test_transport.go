package main

import "fmt"

type Transport interface {
	Move()
	Stop()
	Name() string
}

type Car struct {
	model string
	speed int
}

func (c Car) Move() {
	fmt.Printf("%s едет со скоростью %d км/ч\n", c.model, c.speed)
}

func (c Car) Stop() {
	fmt.Printf("%s остановился\n", c.model)
}

func (c Car) Name() string {
	return c.model
}

type Bicycle struct {
	brand string
}

func (b Bicycle) Move() {
	fmt.Printf("%s крутит педали\n", b.brand)
}

func (b Bicycle) Stop() {
	fmt.Printf("%s тормозит\n", b.brand)
}

func (b Bicycle) Name() string {
	return b.brand
}

type Train struct {
	number int
}

func (t Train) Move() {
	fmt.Printf("Поезд №%d отправляется в путь\n", t.number)
}

func (t Train) Stop() {
	fmt.Printf("Поезд №%d прибыл на станцию\n", t.number)
}

func (t Train) Name() string {
	return fmt.Sprintf("Поезд №%d", t.number)
}

func testTransport(t Transport) {
	fmt.Println("\nЗапускаем:", t.Name())
	t.Move()
	t.Stop()
}

func main() {
	vehicles := []Transport{
		Car{"Жигули", 160},
		Bicycle{"Велосипед"},
		Train{3666},
	}

	for _, vehicle := range vehicles {
		testTransport(vehicle)
	}
}
