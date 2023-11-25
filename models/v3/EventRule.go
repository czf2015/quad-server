package models_v3

type EventRule struct {
	Base
	Name                  string    `gorm:"TYPE:varchar(255)" json:"name"`
	PrimaryClassification string    `gorm:"TYPE:varchar(255)" json:"primary_classification"`
	ClassType             string    `gorm:"TYPE:varchar(255)" json:"class_type"`
	ThreatLevel           string    `gorm:"TYPE:varchar(255)" json:"threat_level"`
	GenerateTimeWindow    FlatMap   `gorm:"TYPE:json" json:"generate_time_window" form:"-"`
	AggregateTimeWindow   FlatMap   `gorm:"TYPE:json" json:"aggregate_time_window" form:"-"`
	Enable                bool      `gorm:"default:true" json:"enable"`
	Description           string    `gorm:"TYPE:varchar(255)" json:"description"`
	Prerequisites         FlatArray `gorm:"TYPE:json" json:"prerequisites" form:"-"`
}
