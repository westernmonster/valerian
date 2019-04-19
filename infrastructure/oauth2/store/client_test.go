package store_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"git.flywk.com/flywiki/api/infrastructure/oauth2/models"
	"git.flywk.com/flywiki/api/infrastructure/oauth2/store"
)

func TestClientStore(t *testing.T) {
	Convey("Test client store", t, func() {
		clientStore := store.NewClientStore()

		err := clientStore.Set("1", &models.Client{ID: "1", Secret: "2"})
		So(err, ShouldBeNil)

		cli, err := clientStore.GetByID("1")
		So(err, ShouldBeNil)
		So(cli.GetID(), ShouldEqual, "1")
	})
}
