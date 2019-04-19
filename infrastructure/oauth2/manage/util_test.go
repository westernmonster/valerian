package manage_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"git.flywk.com/flywiki/api/infrastructure/oauth2/manage"
)

func TestUtil(t *testing.T) {
	Convey("Util Test", t, func() {
		Convey("ValidateURI Test", func() {
			err := manage.DefaultValidateURI("http://www.example.com", "http://www.example.com/cb?code=xxx")
			So(err, ShouldBeNil)
		})
	})
}
