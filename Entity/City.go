package Entity

type City struct {
	ID int	`json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
