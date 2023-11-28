package utilmm

import (
  "reflect"
  "fmt"
)

func sort(arr []struct{v interface{}; e float64}) {
  for i,_ := range arr {
    for j:=i; j<len(arr); j++ {
      if arr[i].e > arr[j].e { arr[i],arr[j] = arr[j],arr[i] }
    }
  }
}

func assertSimplexAndFuncs(simplex []struct{v interface{}; e float64}, Mul func(float64,interface{}) interface{}, Add func(interface{},interface{}) interface{}, Sub func(interface{},interface{}) interface{}) {
  ok:=true
  t:=reflect.TypeOf(simplex[0].v)
  for _,s := range simplex {
    ok = ok && t==reflect.TypeOf(s.v)
  }
  if !ok {panic("invalid simplex")}
  if reflect.TypeOf(Mul(1.,simplex[0].v))!=t {panic("invalid Mul function")}
  if reflect.TypeOf(Add(simplex[0].v,simplex[1].v))!=t {panic("invalid Add function")}
  if reflect.TypeOf(Sub(simplex[0].v,simplex[1].v))!=t {panic("invalid Sub function")}
}

func NelderMead(simplex []struct{v interface{}; e float64}, eval func(interface{}) float64, Mul func(float64,interface{}) interface{}, Add func(interface{},interface{}) interface{}, Sub func(interface{},interface{}) interface{}, Len func(interface{}) float64, log func(int), step *int, lenAcc float64, params ...float64) {
  assertSimplexAndFuncs(simplex,Mul,Add,Sub)

  //coefficients for reflection, expansion, contraction, shrinking
  refC := 1.
  expC := 2.
  conC := 0.5
  shrC := 0.5
  if len(params)>0 {refC = params[0];fmt.Println("refC=",refC)}
  if len(params)>1 {expC = params[1];fmt.Println("expC=",expC)}
  if len(params)>2 {conC = params[2];fmt.Println("conC=",conC)}
  if len(params)>3 {shrC = params[3];fmt.Println("shrC=",shrC)}
  if len(params)>4 {panic("too many parameters to NelderMead")}

  done := false
  for i,a := range simplex {
    simplex[i].e = eval(a.v)
  }
  sort(simplex)
  log(0)

  for !done {
    *step++
    sort(simplex)
    //calculate center
    ac:=simplex[0].v
    for i:=1; i<len(simplex)-1; i++ {
      ac=Add(ac,simplex[i].v)
    }
    ac=Mul(1./float64(len(simplex)-1),ac)
    //calculate reflection
    ar := Add(ac,Mul(refC,Sub(ac,simplex[len(simplex)-1].v)))
    er := eval(ar)

    if simplex[0].e <= er && er < simplex[len(simplex)-2].e {
      //replace worst with reflected -- reiterate
      simplex[len(simplex)-1].v = ar
      simplex[len(simplex)-1].e = er
      log(1)
    } else if er < simplex[0].e {
      //calculate expansion
      ae := Add(ac,Mul(expC,Sub(ar,ac)))
      ee := eval(ae)
      if ee < er {
        //replace worst with expanded -- reiterate
        simplex[len(simplex)-1].v = ae
        simplex[len(simplex)-1].e = ee
        log(2)
      } else {
        //replace worst with reflected -- reiterate
        simplex[len(simplex)-1].v = ar
        simplex[len(simplex)-1].e = er
        log(1)
      }
    } else {
      //now er >= simplex[len(simplex)-2].e
      //calculate contraction
      aco := Add(ac,Mul(conC,Sub(simplex[len(simplex)-1].v,ac)))
      eco := eval(aco)
      if eco < simplex[len(simplex)-1].e {
        //replace worst with contracted -- reiterate
        simplex[len(simplex)-1].v = aco
        simplex[len(simplex)-1].e = eco
        log(3)
      } else {
        //shrink -- reiterate
        for i:=1; i<len(simplex); i++ {
          simplex[i].v = Add(simplex[0].v,Mul(shrC,Sub(simplex[i].v,simplex[0].v)))
          simplex[i].e = eval(simplex[i].v)
        }
        log(4)
      }
    }
    //check convergence
    done = true
    for i,s := range simplex {
      for _,t := range simplex[i:] {
        done = done && Len(Sub(s.v,t.v)) < lenAcc
      }
    }
  }
  sort(simplex)
  simplex[0].e = eval(simplex[0].v)
  log(0)
}
