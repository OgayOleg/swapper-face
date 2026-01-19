package dto

import "errors"

var ErrStatusFailed error = errors.New("vmodel: status failed")
var ErrStatusCanceled error = errors.New("vmodel: status canceled")

var ErrNotFoundForFace error = errors.New("no photo found for face image. Try again")
var ErrNotFoundForTarget error = errors.New("no photo found for target image. Try again")
