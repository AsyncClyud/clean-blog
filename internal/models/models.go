package models

type Article struct {
	Id         int    `json:"Id"`
	Title      string `json:"Title"`
	Content    string `json:"Content"`
	Created_at string `json:"Created_at"`
	Author     int    `json:"Author"`
}

type User struct {
	Id         int     `json:"Id"`
	Username   string  `json:"Username"`
	Password   string  `json:"Password"`
	Bio        *string `json:"Bio"`
	Created_at *string `json:"Created_at"`
}

type Message struct {
	Message string `json:"Message"`
}

type NewPassword struct {
	Old_Password string `json:"Old_Password"`
	New_Password string `json:"New_Password"`
}
