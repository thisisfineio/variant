package variant

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	v = NewVersion("test version", Beta)
)

func init(){
	v.Major = 1
	v.Minor = 4
}

func TestVersion_VersionString(t *testing.T) {
	Convey("We can test a version string", t, func(){
		So(v.VersionString(), ShouldEqual, "1.4_beta")
	})
}

func TestVersions_Append(t *testing.T) {
	Convey("We can Append a version to versions", t, func(){
		vers := &Versions{}
		vers.Append(v)
		So(vers.Len(), ShouldEqual, 1)
		So(vers.History[0], ShouldEqual, v)
	})
}

func TestVersions_NewMajor(t *testing.T) {
	Convey("We can create a new major version", t, func(){
		vers := &Versions{}
		vers.NewMajor("New Major version", Beta)
		So(v.VersionString(), ShouldEqual, "1.0_beta")
	})
}