package locale

import (
	"github.com/gobuffalo/packr"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var (
	bundle = i18n.NewBundle(language.SimplifiedChinese)
)

func LoadTranslateFile() {
	bundle.RegisterUnmarshalFunc("yml", yaml.Unmarshal)

	box := packr.NewBox("../../translate")
	for _, v := range box.List() {
		bundle.MustParseMessageFileBytes(box.Bytes(v), v)
	}

}

func Tr(locale, messageID string) string {

	localizer := i18n.NewLocalizer(bundle, locale)
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
}
