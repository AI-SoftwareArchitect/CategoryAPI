package entities

type Category struct {
    ID          int `json:"id"`
    Name        string    `json:"categoryname"`
    Description string    `json:"categorydescription"`
}

func (Category) TableName() string {
    return "category"
}


