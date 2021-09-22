package go_module_experiments

import "github.com/aaronc/go-module-experiments/foo/v1beta1"
import foo_v2 "github.com/aaronc/go-module-experiments/foo/v2"
import "github.com/aaronc/go-module-experiments/bar"
import bar_v2 "github.com/aaronc/go-module-experiments/bar/v2"
//import "github.com/aaronc/go-module-experiments/foo/v1beta1"
//import "github.com/aaronc/go-module-experiments/foo/v2"

type Test struct {
	Bar bar.Bar
	Bar2 bar_v2.Bar
	FooV1Beta1 v1beta1.Foo
	Foo2 foo_v2.Foo
}