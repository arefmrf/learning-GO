package persistence

type Place struct {
	BaseModel
	UID     string `gorm:"size:26;uniqueIndex"`
	UserUID string `gorm:"size:26"`

	Country  *string `gorm:"size:26"`
	Province *string `gorm:"size:26"`
	City     *string `gorm:"size:26"`

	Info    JSONB `gorm:"type:jsonb"`
	Address string
	Lat     float64
	Lng     float64
}

func (Place) TableName() string { return "places" }
