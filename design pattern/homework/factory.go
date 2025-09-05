package main

import "fmt"

type Processor interface {
	Process(amount float64) error
}

type PayPal struct{}

func (p *PayPal) Process(amount float64) error {
	fmt.Printf("Processing %.2f with PayPal\n", amount)
	return nil
}

type Stripe struct{}

func (s *Stripe) Process(amount float64) error {
	fmt.Printf("Processing %.2f with Stripe\n", amount)
	return nil
}

type Crypto struct{}

func (c *Crypto) Process(amount float64) error {
	fmt.Printf("Processing %.2f with Crypto\n", amount)
	return nil
}

// NewProcessor Factory
func NewProcessor(name string) (Processor, error) {
	switch name {
	case "paypal":
		return &PayPal{}, nil
	case "stripe":
		return &Stripe{}, nil
	case "crypto":
		return &Crypto{}, nil
	default:
		return nil, fmt.Errorf("unknown processor: %s", name)
	}
}

func main() {
	names := []string{"paypal", "stripe", "crypto"}
	for _, name := range names {
		p, err := NewProcessor(name)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if err := p.Process(10); err != nil {
			fmt.Println("Process failed:", err)
		}
	}
}
