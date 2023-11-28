package utilmm

import (
  "github.com/mumax/3/engine"
  "github.com/mumax/3/script"
  "github.com/mumax/3/data"
  "reflect"
)


type SclFun struct{
  fun func() float64
}

func NewSclFun(f func() float64) SclFun { return SclFun{f} }

func (f SclFun) Eval() interface{} { return f }
func (f SclFun) Type() reflect.Type { return script.ScalarFunction_t }
func (f SclFun) Child() []script.Expr { return []script.Expr{engine.World.Resolve("t")} }
func (f SclFun) Fix() script.Expr { return f }
func (f SclFun) Float() float64 { return f.fun() }


type VecFun struct{
  fun func() data.Vector
}

func NewVecFun(f func() data.Vector) VecFun { return VecFun{f} }

func (f VecFun) Eval() interface{} { return f }
func (f VecFun) Type() reflect.Type { return script.VectorFunction_t }
func (f VecFun) Child() []script.Expr { return []script.Expr{engine.World.Resolve("t")} }
func (f VecFun) Fix() script.Expr { return f }
func (f VecFun) Float3() data.Vector  { return f.fun() }
