package hystrix

import (
	hgo "github.com/afex/hystrix-go/hystrix"
)

func ConfigureHystrix() {

	// Configure test command.
	// Default rolling window is 10s.
	hgo.ConfigureCommand(TestCommand, hgo.CommandConfig{
		Timeout:                350, // TimeoutMs
		MaxConcurrentRequests:  1,
		RequestVolumeThreshold: 10,
		SleepWindow:            15000, // SleepWindowMs
		ErrorPercentThreshold:  20,
	})

}
