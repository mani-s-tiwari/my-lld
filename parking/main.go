package main

import "parking/services"

func main() {
	sm := services.SpotManager{}
	sm.Init()

	sm.Park("Indica", 122)
	sm.Park("Renault", 2231)

	sm.Checkmycar(122)
	sm.Statusall()

	sm.Unpark(122)
	sm.Statusall()

	sm.Statusall()
	sm.Unpark(122)
}
