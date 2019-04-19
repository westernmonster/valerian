package generates_test

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"git.flywk.com/flywiki/api/infrastructure/oauth2"
	"git.flywk.com/flywiki/api/infrastructure/oauth2/generates"
	"git.flywk.com/flywiki/api/infrastructure/oauth2/models"
)

func TestAuthorize(t *testing.T) {
	Convey("Test Authorize Generate", t, func() {
		data := &oauth2.GenerateBasic{
			Client: &models.Client{
				ID:     "123456",
				Secret: "123456",
			},
			UserID:   "000000",
			CreateAt: time.Now(),
		}
		gen := generates.NewAuthorizeGenerate()
		code, err := gen.Token(data)
		So(err, ShouldBeNil)
		So(code, ShouldNotBeEmpty)
		Println("\nAuthorize Code:" + code)
	})
}
