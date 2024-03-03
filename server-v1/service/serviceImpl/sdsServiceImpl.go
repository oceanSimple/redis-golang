package serviceImpl

import "server-v1/model"

type SdsServiceImpl struct {
}

func (s SdsServiceImpl) NewSds(str string) model.Sds {
	return model.Sds{Data: []byte(str)}
}

func (s SdsServiceImpl) NewEmptySds() model.Sds {
	return model.Sds{Data: []byte{}}
}

func (s SdsServiceImpl) Equal(s1 model.Sds, s2 model.Sds) bool {
	return string(s1.Data) == string(s2.Data)
}
