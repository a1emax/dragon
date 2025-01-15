package vars

import (
	"dragon/pkg/domain"
)

var GamePlay struct {
	IgnoreDirectionPad bool
	Session            *domain.Session
}
