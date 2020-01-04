package base

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	vtzh "github.com/go-playground/validator/v10/translations/zh"
	"github.com/idcpj/red_packet/infra"
	"github.com/sirupsen/logrus"
)

var validate *validator.Validate
var translator *ut.Translator

func Validate() *validator.Validate {
	return validate
}

func Transtate() *ut.Translator {
	return translator
}

type ValidatorStater struct {
	infra.BaseStarter
}

func (V *ValidatorStater) Init(ctx infra.StarterContext) {

	validate = validator.New()
	cn := zh.New()
	uni := ut.New(cn, cn)

	// 设置默认中文
	transZH, found := uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, transZH)
		if err != nil {
			logrus.Error(err)
		}
	} else {
		logrus.Error("not found translationsc")
	}

}
