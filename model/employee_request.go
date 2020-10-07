package model

// EmployeeLoginPayload ...
type EmployeeLoginPayload struct {
	Phone    string `json:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" valid:"required, type(string), stringlength(6|20)"`
}

// EmployeeRegisterPayload ...
type EmployeeRegisterPayload struct {
	Name     string `json:"name" valid:"required, stringlength(1|20)"`
	Phone    string `json:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" valid:"required, type(string), stringlength(6|20)"`
}

// EmployeeUpdatePayload ...
type EmployeeUpdatePayload struct {
	Name     string `json:"name" valid:"required, stringlength(1|20)"`
	Phone    string `json:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" valid:"required, type(string), stringlength(6|50)"`
	Active   bool   `json:"active" valid:"required, type(bool)"`
}

// EmployeeQuery ...
type EmployeeQuery struct {
	Name   string `query:"name"`
	Active string `query:"active"`
	Page   int    `query:"page"`
}
