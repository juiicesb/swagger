package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juiicesb/swagger/testdata/composition/common"
)

type Foo struct {
	Field1 string
}
type Bar struct {
	Field2 string
}

type FooBar struct {
	Foo
	Bar
}

type FooBarPointer struct {
	*common.ResponseFormat
	*Foo
	*Bar
}

type BarMap map[string]Bar

type FooBarMap struct {
	Field3 map[string]MapValue
}

type MapValue struct {
	Field4 string
}

// @Description get Foo
// @ID get-foo
// @Accept json
// @Produce json
// @Success 200 {object} api.Foo
// @Router /testapi/get-foo [get]
func GetFoo(c *gin.Context) {
	//write your code
	var _ = Foo{}
}

// @Description get Bar
// @ID get-bar
// @Accept json
// @Produce json
// @Success 200 {object} api.Bar
// @Router /testapi/get-bar [get]
func GetBar(c *gin.Context) {
	//write your code
	var _ = Bar{}
}

// @Description get FooBar
// @ID get-foobar
// @Accept json
// @Produce json
// @Success 200 {object} api.FooBar
// @Router /testapi/get-foobar [get]
func GetFooBar(c *gin.Context) {
	//write your code
	var _ = FooBar{}
}

// @Description get FooBarPointer
// @ID get-foobar-pointer
// @Accept json
// @Produce json
// @Success 200 {object} api.FooBarPointer
// @Router /testapi/get-foobar-pointer [get]
func GetFooBarPointer(c *gin.Context) {
	//write your code
	var _ = FooBarPointer{}
}

// @Description get BarMap
// @ID get-bar-map
// @Accept json
// @Produce json
// @Success 200 {object} api.BarMap
// @Router /testapi/get-barmap [get]
func GetBarMap(c *gin.Context) {
	//write your code
	var _ = BarMap{}
}

// @Description get FoorBarMap
// @ID get-foo-bar-map
// @Accept json
// @Produce json
// @Success 200 {object} api.FooBarMap
// @Router /testapi/get-foobarmap [get]
func GetFooBarMap(c *gin.Context) {
	//write your code
	var _ = FooBarMap{}
}
