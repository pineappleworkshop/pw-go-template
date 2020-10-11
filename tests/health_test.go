package test

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// This tests the /health endpoint
func TestProject(t *testing.T) {
	Convey("If API exists", t, func() {
		baseUrl := getBaseUrl(ENV)
		So(baseUrl, ShouldNotEqual, "")

		Convey("When consuming the /health endpoint", func() {
			res, body, err := Health(baseUrl)

			Convey("Then it should return a 200 with the correct body", func() {
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
				So(body, ShouldNotBeNil)
				So(res.StatusCode, ShouldEqual, http.StatusOK)
				for k, v := range body {
					if k == "service" {
						So(v.(string), ShouldEqual, "test-services")
					}
					if k == "status" {
						So(int(v.(float64)), ShouldEqual, http.StatusOK)
					}
				}
			})
		})
	})
}
