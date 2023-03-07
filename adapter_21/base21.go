package main

import "fmt"

type transport interface {
	navigateToDestination()
}

type client struct {
}

func (t *client) startNavigation(trans transport) {
	fmt.Println("Клиент начинает процесс навигации")
	trans.navigateToDestination()
}

type boat struct {
}

func (b *boat) navigateToDestination() {
	fmt.Println("Корабль будет плыть к острову")
}

type car struct {
}

func (c *car) driveToDestination() {
	fmt.Println("Машина едет по маршруту")
}

type carAdapter struct {
	carTransport *car
}

func (c *carAdapter) navigateToDestination() {
	fmt.Println("Адаптер изменил машину чтобы была возможности двигаться по маршруту")
	c.carTransport.driveToDestination()
}

func main() {
	client := &client{}
	boat := &boat{}

	client.startNavigation(boat)

	car := &car{}
	carAdapter := &carAdapter{
		carTransport: car,
	}

	client.startNavigation(carAdapter)
}
