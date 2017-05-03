package validatorsample

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/kylelemons/godebug/pretty"
)

func TestUserValidateWithCustomFunc(t *testing.T) {
	validate := validator.New()
	if err := validate.RegisterValidation("userPhonenumber", userPhonenumber); err != nil {
		t.Fatal(err)
	}

	// perfect request
	v1 := `
	{
		"firstname": "Akira",
		"lastname": "Chiku",
		"age": 30,
		"email": "akira.chiku@gmail.com",
		"phonenumber": "09012341234",
		"twitter": "https://twitter/_achiku"
	}
	`
	// wrong email fromat
	// wrong phonenumber fromat
	// wrong age range
	// wrong http url fromat
	v2 := `
	{
		"firstname": "Akira",
		"lastname": "Chiku",
		"age": 140,
		"email": "akira.chikugmail.com",
		"phonenumber": "0901234123411",
		"twitter": "://twitter/_achiku"
	}
	`
	data := []struct {
		Request string
	}{
		{Request: v1},
		{Request: v2},
	}
	for _, d := range data {
		decoder := json.NewDecoder(strings.NewReader(d.Request))
		var u User
		if err := decoder.Decode(&u); err != nil {
			t.Fatal(err)
		}

		if err := validate.Struct(u); err != nil {
			pretty.Print(u)
			for _, e := range err.(validator.ValidationErrors) {
				t.Errorf(
					"%s(%s) %s validation failed: %s", e.Namespace(), e.Kind(), e.ActualTag(), e.Value())
			}
		}
	}
}

func TestUserValidateWithNest(t *testing.T) {
	validate := validator.New()

	// good request
	v1 := `
	{
		"name": "parent A",
		"children": [
			{"name": "child A"},
			{"name": "child B"}
		]
	}
	`
	// bad request
	v2 := `
	{
		"name": "parent A",
		"children": [
			{"name": "child A"},
			{}
		]
	}
	`
	// good request
	v3 := `
	{
		"name": "parent A"
	}
	`
	data := []struct {
		Request string
	}{
		{Request: v1},
		{Request: v2},
		{Request: v3},
	}
	for _, d := range data {
		decoder := json.NewDecoder(strings.NewReader(d.Request))
		var p Parent
		if err := decoder.Decode(&p); err != nil {
			t.Fatal(err)
		}

		if err := validate.Struct(p); err != nil {
			pretty.Print(p)
			for _, e := range err.(validator.ValidationErrors) {
				t.Errorf(
					"%s(%s) %s validation failed: %s", e.Namespace(), e.Kind(), e.ActualTag(), e.Value())
			}
		}
	}
}
