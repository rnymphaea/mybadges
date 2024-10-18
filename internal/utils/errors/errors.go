package errors

import "errors"

var ErrUserNotFound = errors.New("user not found")

var ErrCheckingPassword = errors.New("error checking password")

var ErrInvalidCredentials = errors.New("invalid credentials")

var ErrNoSession = errors.New("cannot create session to connect s3")

var ErrUploadFile = errors.New("error uploading file")
