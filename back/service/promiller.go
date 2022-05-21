package service

import "userBar/back/pkg/repository"

type PromillerService struct {
	rep repository.Promiller
}

func NewPromillerService(rep repository.Promiller) *PromillerService  {
	return &PromillerService{rep: rep}
}

func (p *PromillerService) PromilleDec() string {
	return p.rep.PromillerDec()
}