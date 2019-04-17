package generates_test

import (
	"testing"
	"time"

	"git.flywk.com/flywiki/api/infrastructure/oauth2"
	. "github.com/smartystreets/goconvey/convey"

	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/models"
)

func TestAccess(t *testing.T) {
	Convey("Test Access Generate", t, func() {
		data := &oauth2.GenerateBasic{
			Client: &models.Client{
				ID:     "123456",
				Secret: "123456",
			},
			UserID:   "000000",
			CreateAt: time.Now(),
		}
		gen := generates.NewAccessGenerate()
		access, refresh, err := gen.Token(data, true)
		So(err, ShouldBeNil)
		So(access, ShouldNotBeEmpty)
		So(refresh, ShouldNotBeEmpty)
		Println("\nAccess Token:" + access)
		Println("Refresh Token:" + refresh)
	})
}
