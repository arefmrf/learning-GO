package services

import (
	"trip/internal/domain"
	"trip/internal/persistence"

	"gorm.io/gorm"
)

func SaveHostDaily(
	db *gorm.DB,
	daily *persistence.HostDaily,
	hostCategory int,
) error {

	d := domain.HostDaily{
		Price:        daily.Price,
		ExtraPrice:   daily.ExtraPrice,
		Discount:     int(daily.Discount),
		Capacity:     daily.Capacity,
		MaleBooked:   int(daily.MaleBooked),
		FemaleBooked: int(daily.FemaleBooked),
		HostCategory: hostCategory,
	}

	d.PrepareForSave()

	daily.FinalPrice = d.FinalPrice
	daily.TotalCapacity = int16(d.TotalCapacity)

	now := int64(domain.Now())
	if daily.CreatedAt == 0 {
		daily.CreatedAt = now
	}
	daily.UpdatedAt = now

	return db.Save(daily).Error
}

func DeleteHostDaily(db *gorm.DB, daily *persistence.HostDaily) error {
	d := domain.HostDaily{
		MaleBooked:   int(daily.MaleBooked),
		FemaleBooked: int(daily.FemaleBooked),
	}

	if err := d.CanDelete(); err != nil {
		return err
	}

	return db.Delete(daily).Error
}
