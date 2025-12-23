package persistence

type Host struct {
	BaseModel

	UID        string `gorm:"size:26;uniqueIndex"`
	UserUID    string `gorm:"size:26"`
	Identifier string `gorm:"size:16;uniqueIndex"`

	Title       string `gorm:"size:128"`
	Description string

	PlaceID    *uint
	HostTypeID *uint

	Hidden       bool
	Priority     int16
	Online       bool
	OutOfService bool

	Info   JSONB `gorm:"type:jsonb"`
	Status int16

	Rate      float64
	RateCount int16

	CurrencyID  *uint
	CancelLevel int16

	CreatedAt int64 `gorm:"type:int"`
	UpdatedAt int64 `gorm:"type:int"`

	Facilities []Facility `gorm:"many2many:host_facilities"`
	Tags       []HostTag  `gorm:"many2many:host_tags"`
}

func (Host) TableName() string { return "hosts" }
