package anx

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	client     *Anx
	testClient = &Anx{}
)

// TestNew tests the functionality of New
func TestNew(t *testing.T) {
	Convey("Testing on New...", t, func() {
		client = New("demoKey", "demoSecret")
		So(client, ShouldNotResemble, testClient)
		So(client.APIKeyID, ShouldEqual, "demoKey")
		So(client.APIKeySecret, ShouldEqual, "demoSecret")
	})
}
