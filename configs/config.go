package configs

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Specification is base structure of package
type Specification struct {
	Debug      bool
	Port       int
	User       string
	Users      []string
	Rate       float32
	Timeout    time.Duration
	ColorCodes map[string]int
}

// ReadConfig reads config from enviroments
func ReadConfig() error {
	fmt.Printf("\nЧтение переменных окружения:\n")
	var s Specification
	err := envconfig.Process("myapp", &s)
	if err != nil {
		return fmt.Errorf("error envconfig.Process:%w", err)
	}
	format := "Debug: %v\nPort: %d\nUser: %s\nRate: %f\nTimeout: %s\n"
	_, err = fmt.Printf(format, s.Debug, s.Port, s.User, s.Rate, s.Timeout)
	if err != nil {
		return fmt.Errorf("error fmt.Printf:%w", err)
	}

	fmt.Println("Users:")
	for _, u := range s.Users {
		fmt.Printf("  %s\n", u)
	}

	fmt.Println("Color codes:")
	for k, v := range s.ColorCodes {
		fmt.Printf("  %s: %d\n", k, v)
	}
	return nil
}
