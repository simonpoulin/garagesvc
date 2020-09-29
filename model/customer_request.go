package model

// CustomerLoginPayload ...
type CustomerLoginPayload struct {
	Phone    string `json:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" valid:"required, type(string), stringlength(6|20)"`
}

// CustomerRegisterPayload ...
type CustomerRegisterPayload struct {
	Name     string `json:"name" valid:"required, stringlength(1|20)"`
	Phone    string `json:"phone" valid:"required, type(string), stringlength(10|10)"`
	Password string `json:"password" valid:"required, type(string), stringlength(6|20)"`
}

// CustomerUpdatePayload ...
type CustomerUpdatePayload struct {
	Name     string `json:"name" valid:"required, stringlength(1|20)"`
	Password string `json:"password" valid:"required, type(string), stringlength(6|20)"`
}

// CustomerQuery ...
type CustomerQuery struct {
	Name string `query:"name"`
	Page int    `query:"page"`
}
