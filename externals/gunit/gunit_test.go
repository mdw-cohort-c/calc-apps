package gunit_test

import (
	"testing"

	"github.com/mdw-cohort-c/calc-apps/externals/gunit"
	"github.com/mdw-cohort-c/calc-apps/externals/should"
)

func TestMySuperCoolFixture(t *testing.T) {
	gunit.Run(t, new(MySuperCoolFixture))
}

type MySuperCoolFixture struct {
	*gunit.Fixture
}

func (this *MySuperCoolFixture) Test1() {
	this.So(1, should.Equal, 1)
}
