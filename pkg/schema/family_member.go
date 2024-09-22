package schema

type FamilyMember struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
}
