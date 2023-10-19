package main

import "fmt"

const ebulicaoKelvin = 373.15

func main() {
	celsius := ebulicaoKelvin - 273.15
	fmt.Printf("A temperatura de ebulição da água em ºC = %.2f\n", celsius)
}
