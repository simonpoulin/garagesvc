package model

// CompanyCreatePayload ...
type CompanyCreatePayload struct {
	Name     string   `json:"name" valid:"required, stringlength(1|20)"`
	Location Location `json:"location"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone" valid:"required, type(string), stringlength(10|10)"`
	Address  string   `json:"address" valid:"required, stringlength(1|50)"`
}

// CompanyUpdatePayload ...
type CompanyUpdatePayload struct {
	Name     string   `json:"name" valid:"required, stringlength(1|20)"`
	Location Location `json:"location"`
	Email    string   `json:"email"`
	Address  string   `json:"address" valid:"required, stringlength(1|50)"`
	Phone    string   `json:"phone" valid:"required, type(string), stringlength(10|10)"`
	Active   bool     `json:"active" valid:"required, type(bool)"`
}

// CompanyQuery ...
type CompanyQuery struct {
	Name   string `query:"name"`
	Active string `query:"active"`
	Page   int    `query:"page"`
	Phone  string `query:"phone"`
}
