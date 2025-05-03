package main

import (
	"fmt"
	"time"
)

func main() {
	// Creamos un canal que permite 5 solicitudes por segundo
	limiter := time.Tick(200 * time.Millisecond)
	
	// Simulamos 10 solicitudes
	for i := 1; i <= 10; i++ {
		// Esperamos a que llegue un valor del ticker antes de procesar
		<-limiter
		
		fmt.Printf("Procesando solicitud %d en %s\n", i, time.Now().Format("15:04:05.000"))
	}
}