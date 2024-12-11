package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK                      = "OK"
	StatusError                   = "Error"
	InternalErrorMessage          = "внутренняя ошибка сервиса. Попробуйте позже."
	DecodeErrorMessage            = "ошибка декодирования тела запроса"
	InvalidRequestErrorMessage    = "невалидный запрос"
	WrongCredentialsErrorMessage  = "неправильные имя пользователя или пароль"
	UserAlreadyExistsErrorMessage = "пользователь уже существует"
	UnauthorizedErrorMessage      = "unauthorized"
	ForbiddenErroMessage          = "forbidden"
	BadRequestMessage             = "некорректный запрос"
	NotFoundMessage               = "not found"
	RateLImitErrorMessage         = "превышен лимит запросов"
	UserNOtActivatedErrorMessage  = "пользователь не активирован. Обратитесь к администратору"
	WrongApiTokenMessage          = "невалидный api token. Измените его в управлении магазина"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func InternalError() Response {
	return Error(InternalErrorMessage)
}

func RateLimitExceeded() Response {
	return Error(RateLImitErrorMessage)
}

func DecodeError() Response {
	return Error(DecodeErrorMessage)
}

func InvalidRequestError() Response {
	return Error(InvalidRequestErrorMessage)
}

func WrongCredentialsError() Response {
	return Error(WrongCredentialsErrorMessage)
}

func UserNotActivatedError() Response {
	return Error(UserNOtActivatedErrorMessage)
}

func UnauthorizedError() Response {
	return Error(UnauthorizedErrorMessage)
}

func ForbiddenError() Response {
	return Error(ForbiddenErroMessage)
}

func WrongApiTokenError() Response {
	return Error(WrongApiTokenMessage)
}

func BadRequestError(msg string) Response {
	return Error(fmt.Sprintf("%s: %s", BadRequestMessage, msg))
}

func UserAlreadyExistsError() Response {
	return Error(UserAlreadyExistsErrorMessage)
}

func NotFoundError() Response {
	return Error(NotFoundMessage)
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("поле %s явялется обязательным", err.Field()))
		case "min":
			errMsgs = append(
				errMsgs,
				fmt.Sprintf(
					"поле %s должно содержать больше, чем %s символов",
					err.Field(),
					err.Param(),
				),
			)
		case "max":
			errMsgs = append(
				errMsgs,
				fmt.Sprintf(
					"поле %s должно содержать меньше, чем %s символов",
					err.Field(),
					err.Param(),
				),
			)
		case "containsany":
			errMsgs = append(
				errMsgs,
				fmt.Sprintf("поле %s must contains any of special character", err.Field()),
			)
		case "email":
			errMsgs = append(
				errMsgs,
				fmt.Sprintf("поле %s не является корректной электронной почтой", err.Field()),
			)
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("поле %s не прошло валидацию", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
