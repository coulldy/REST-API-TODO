package todo

import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlereadyExists = errors.New("task already exists")
