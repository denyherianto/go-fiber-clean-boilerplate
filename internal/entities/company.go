package entities

// Company struct to describe Company object.
type Company struct {
	BaseEntity
	ID   uint   `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	Name string `db:"name" json:"name" validate:"required,lte=255"`
}
