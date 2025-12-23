package persistence

type HostType struct {
	BaseModel
	UID         string `gorm:"size:26;uniqueIndex"`
	Category    int16
	Title       string `gorm:"size:128"`
	Translation JSONB  `gorm:"type:jsonb"`
	Icon        *string
	Active      bool
	Priority    int16
}

func (HostType) TableName() string { return "host_types" }
