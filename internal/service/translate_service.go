package service

import "void-project/pkg/translation"

type TranslateService struct{}

func NewTranslateService() *TranslateService {
	return &TranslateService{}
}

func (t *TranslateService) Translate(text, source, target string) ([]translation.TranslationList, error) {
	res, err := translation.Translate(text, source, target)
	return res.TranslationList, err
}
