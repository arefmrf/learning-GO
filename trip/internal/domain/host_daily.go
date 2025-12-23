package domain

import (
	"errors"
	_ "errors"
)

type HostDaily struct {
	Price      float64
	ExtraPrice float64
	FinalPrice float64
	Discount   int

	Capacity      Capacity
	TotalCapacity int

	MaleBooked   int
	FemaleBooked int

	HostCategory int // 1 = shared, 2 = dedicated
}

func (d *HostDaily) CalculateFinalPrice() float64 {
	if d.Discount > 0 {
		return d.Price * float64(100-d.Discount) / 100
	}
	return d.Price
}

func (d *HostDaily) CalculateTotalCapacity() {
	basePerson := d.Capacity["base_person"]

	if basePerson == 0 {
		// shared
		d.TotalCapacity =
			d.Capacity["male"] +
				d.Capacity["female"]
	} else {
		// dedicated
		d.TotalCapacity =
			basePerson +
				d.Capacity["extra_person"]
	}
}

func (d *HostDaily) PrepareForSave() {
	d.FinalPrice = d.CalculateFinalPrice()
	d.CalculateTotalCapacity()
}

func (d *HostDaily) CanDelete() error {
	if d.MaleBooked != 0 || d.FemaleBooked != 0 {
		return errors.New("er1056")
	}
	return nil
}

func (d *HostDaily) IsBooked() bool {
	if d.HostCategory == 2 { // dedicated
		return d.MaleBooked != 0
	}
	return d.MaleBooked+d.FemaleBooked == d.TotalCapacity
}
