package validator

import (
	"github.com/go-playground/validator"
)

func Validation(s interface{}) string {
	validate := validator.New()
	err := validate.Struct(s)
	errMessage := ""

	json := ""

	if err != nil {
		json = `{"errors": {`
		errs := err.(validator.ValidationErrors)
		for i, e := range errs {

			switch e.ActualTag() {
			case "required":
				errMessage = `"` + e.Field() + `" : {"message": "` + FieldString(e.Field()) + `は必須入力です。"}`
			case "email":
				errMessage = `"Email" : {"message": "正しいメールアドレスの形式で入力してください。"}`
			case "max":
				errMessage = `"` + e.Field() + `" : {"message": "` + FieldString(e.Field()) + `は` + e.Param() + `文字以内で入力してください"}`
			default:
			}
			if i < len(errs)-1 {
				json = json + errMessage + ","
			} else {
				json = json + errMessage
			}
		}
		json = json + `}}`
	}
	return json
}

func FieldString(field string) string {
	switch field {
	case "Firstname":
		return "名字"
	case "Lastname":
		return "名前"
	case "Email":
		return "メールアドレス"
	case "Password":
		return "パスワード"
	case "PasswordConfirmation":
		return "パスワード(確認用)"
	default:
		return ""
	}
	return ""
}
