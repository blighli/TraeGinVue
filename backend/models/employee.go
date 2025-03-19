package models

type Employee struct {
    ID         string `json:"id" gorm:"primaryKey"`
    Name       string `json:"name"`
    Department string `json:"department"`
    Position   string `json:"position"`
    JoinDate   string `json:"joinDate"`
    Status     string `json:"status"`
}