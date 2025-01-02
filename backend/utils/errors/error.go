package errors

type AppError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func NewAppError(code int, message string) *AppError {
    return &AppError{
        Code:    code,
        Message: message,
    }
}