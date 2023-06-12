package models

import (
    "errors"
)

var (
    ErrNoRecords = errors.New("models: no matching record found")
    ErrInvalidCredentials = errors.New("nodels: invalid credentials")
    ErrDuplicateEmail = errors.New("models: duplicate email")
)


