package service

import "server-v1/model"

type SdsService interface {
	NewSds(str string) model.Sds
	NewEmptySds() model.Sds
	Equal(s1 model.Sds, s2 model.Sds) bool
}
