package validatorsample

import (
	"regexp"

	null "gopkg.in/guregu/null.v3"

	"github.com/go-playground/validator"
)

const (
	userPhonenumberRegexString = "^0[789]0[0-9]{8}$"
)

var (
	userPhonenumberRegex = regexp.MustCompile(userPhonenumberRegexString)
)

// User user
type User struct {
	FirstName   string      `json:"firstname" validate:"required"`
	LastName    string      `json:"lastname" validate:"required"`
	Age         uint8       `json:"age" validate:"gte=0,lte=130"`
	Email       string      `json:"email" validate:"required,email"`
	PhoneNumber string      `json:"phonenumber" validate:"required,userPhonenumber"`
	Twitter     null.String `json:"twitter,omitempty" validate:"url"`
}

func userPhonenumber(fl validator.FieldLevel) bool {
	return userPhonenumberRegex.MatchString(fl.Field().String())
}

// Parent parent
type Parent struct {
	Name     string  `json:"name" validate:"required"`
	Children []Child `json:"children" validate:"required,dive"`
}

// Child child
type Child struct {
	Name string `json:"name" validate:"required"`
}
