package common

import (
	"testing"

	apiclient "github.com/openservicebrokerapi/osb-checker/client"
	"github.com/openservicebrokerapi/osb-checker/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCatalog(t *testing.T) {
	Convey("Query service catalog", t, func() {

		So(testAPIVersionHeader(config.GenerateCatalogURL(), "GET"), ShouldEqual, nil)
		So(testAuthentication(config.GenerateCatalogURL(), "GET"), ShouldEqual, nil)

		Convey("should return list of registered service classes as JSON payload", func() {
			code, body, err := apiclient.Default.GetCatalog()

			So(err, ShouldEqual, nil)
			So(code, ShouldEqual, 200)
			So(testJSONSchema(body), ShouldEqual, nil)
		})
	})
}
