package controller

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sysatom/framework/pkg/types"
	"github.com/sysatom/framework/pkg/types/protocol"
	"net/http"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type ValidateController struct {
}

func NewValidateController() ValidateController {
	return ValidateController{}
}

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	Gender         string     `validate:"oneof=male female prefer_not_to"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func (controller ValidateController) Validate(c echo.Context) error {
	zhT := zh.New()
	uni := ut.New(zhT, zhT)

	trans, _ := uni.GetTranslator("zh")

	validate = validator.New(validator.WithRequiredStructEnabled())
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)

	validateStruct(trans)
	validateVariable(trans)

	return c.JSON(http.StatusOK, protocol.NewSuccessResponse(types.KV{"type": "validate"}))
}

func validateStruct(trans ut.Translator) {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Gender:         "male",
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
			return
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				fmt.Println(e.Namespace())
				fmt.Println(e.Field())
				fmt.Println(e.StructNamespace())
				fmt.Println(e.StructField())
				fmt.Println(e.Tag())
				fmt.Println(e.ActualTag())
				fmt.Println(e.Kind())
				fmt.Println(e.Type())
				fmt.Println(e.Value())
				fmt.Println(e.Param())
				fmt.Println(e.Translate(trans))
				fmt.Println()
			}
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

	// save user to database
}

func validateVariable(trans ut.Translator) {
	myEmail := "joeybloggs.gmail.com"

	err := validate.Var(myEmail, "required,email")

	if err != nil {
		fmt.Println(err) // output: Key: "" Error:Field validation for "" failed on the "email" tag

		errs := err.(validator.ValidationErrors)
		fmt.Println(errs.Translate(trans))

		return
	}

	// email ok, move on
}
