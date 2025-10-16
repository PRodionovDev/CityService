package Entity

type Region struct {
	ID int	`json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	//Number int `json:"number"`
}
