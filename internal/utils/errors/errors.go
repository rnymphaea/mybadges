package errors

import "errors"

var ErrUserNotFound = errors.New("user not found")

var ErrCheckingPassword = errors.New("error checking password")

var ErrInvalidCredentials = errors.New("invalid credentials")

var ErrNoSession = errors.New("cannot create session to connect s3")

var ErrUploadFile = errors.New("error uploading file")

var ErrGetUserByEmail = errors.New("error getting user by email")

var ErrNoAuthHeader = errors.New("authorization header is required")

var ErrNoBearerToken = errors.New("bearer token is required")
