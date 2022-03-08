package main

import (
	"machine"
	"time"
)

func main() {
	// 500Hz
	var period uint64 = 1e9 / 500

	pin := machine.LED
	pwm := machine.PWM4

	pin.Configure(machine.PinConfig{Mode: machine.PinPWM})

	err := pwm.Configure(machine.PWMConfig{Period: period})
	if err != nil {
		println(err.Error())
	}
	ch, err := pwm.Channel(pin)
	if err != nil {
		println(err.Error())
	}

	pwm.SetTop(uint32(255))

	for {
		// Fade in
		for i := 0; i <= 255; i++ {
			pwm.Set(ch, uint32(i))
			time.Sleep(time.Millisecond * 10)
		}

		// Fade out
		for i := 255; i >= 0; i-- {
			pwm.Set(ch, uint32(i))
			time.Sleep(time.Millisecond * 10)
		}

		// Wait one second
		time.Sleep(time.Millisecond * 1000)
	}
}
