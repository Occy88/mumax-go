package utilmm

import (
  "github.com/mumax/3/engine"
  "math"
)

func ConeOpeningAngleX() (float64,float64) {
  size := engine.Mesh().Size()
  th:=0.
  vari:=0.
  for ix:=0; ix<size[0]; ix++ {
    for iy:=0; iy<size[1]; iy++ {
      for iz:=0; iz<size[2]; iz++ {
        s:=engine.M.GetCell(ix,iy,iz)
        dummy:=math.Atan2(math.Sqrt(s.Y()*s.Y()+s.Z()*s.Z()),s.X())
        th+=dummy
        vari+=dummy*dummy
      }
    }
  }
  th/=float64(size[0]*size[1]*size[2])
  vari/=float64(size[0]*size[1]*size[2])
  vari-=th*th
  return th,math.Sqrt(vari)
}

func ConeWindingsX() [][]float64 {
  size:= engine.Mesh().Size()
  res := make([][]float64, size[1])
  for iy:=0; iy<size[1]; iy++ {
    res[iy]=make([]float64, size[2])
    for iz:=0; iz<size[2]; iz++ {
      s:=engine.M.GetCell(0,iy,iz)
      phi0:=math.Atan2(s.Z(),s.Y())/(2*math.Pi)
      pp:=phi0
      for ix:=1; ix<size[0]; ix++ {
        t:=engine.M.GetCell(ix,iy,iz)
        np:=math.Atan2(t.Z(),t.Y())/(2*math.Pi)
        res[iy][iz]+=math.Mod(np-pp+1.5,1)-0.5
        pp=np
      }
      res[iy][iz]+=math.Mod(phi0-pp+1.5,1)-0.5
    }
  }
  return res
}

func ConeWindingsXyz(iy, iz int) float64 {
  size:= engine.Mesh().Size()
  res := 0.
  s:=engine.M.GetCell(0,iy,iz)
  phi0:=math.Atan2(s.Z(),s.Y())/(2*math.Pi)
  pp:=phi0
  for ix:=1; ix<size[0]; ix++ {
    t:=engine.M.GetCell(ix,iy,iz)
    np:=math.Atan2(t.Z(),t.Y())/(2*math.Pi)
    res+=math.Mod(np-pp+1.5,1)-0.5
    pp=np
  }
  res+=math.Mod(phi0-pp+1.5,1)-0.5
  return res
}

func ConeOpeningAngleY() (float64,float64) {
  size := engine.Mesh().Size()
  th:=0.
  vari:=0.
  for ix:=0; ix<size[0]; ix++ {
    for iy:=0; iy<size[1]; iy++ {
      for iz:=0; iz<size[2]; iz++ {
        s:=engine.M.GetCell(ix,iy,iz)
        dummy:=math.Atan2(math.Sqrt(s.X()*s.X()+s.Z()*s.Z()),s.Y())
        th+=dummy
        vari+=dummy*dummy
      }
    }
  }
  th/=float64(size[0]*size[1]*size[2])
  vari/=float64(size[0]*size[1]*size[2])
  vari-=th*th
  return th,math.Sqrt(vari)
}

func ConeWindingsY() [][]float64 {
  size:= engine.Mesh().Size()
  res := make([][]float64, size[0])
  for ix:=0; ix<size[0]; ix++ {
    res[ix]=make([]float64, size[2])
    for iz:=0; iz<size[2]; iz++ {
      s:=engine.M.GetCell(ix,0,iz)
      phi0:=math.Atan2(s.X(),s.Z())/(2*math.Pi)
      pp:=phi0
      for iy:=1; iy<size[1]; iy++ {
        t:=engine.M.GetCell(ix,iy,iz)
        np:=math.Atan2(t.X(),t.Z())/(2*math.Pi)
        res[ix][iz]+=math.Mod(np-pp+1.5,1)-0.5
        pp=np
      }
      res[ix][iz]+=math.Mod(phi0-pp+1.5,1)-0.5
    }
  }
  return res
}

func ConeWindingsYxz(ix, iz int) float64 {
  size:= engine.Mesh().Size()
  res := 0.
  s:=engine.M.GetCell(ix,0,iz)
  phi0:=math.Atan2(s.X(),s.Z())/(2*math.Pi)
  pp:=phi0
  for iy:=1; iy<size[1]; iy++ {
    t:=engine.M.GetCell(ix,iy,iz)
    np:=math.Atan2(t.X(),t.Z())/(2*math.Pi)
    res+=math.Mod(np-pp+1.5,1)-0.5
    pp=np
  }
  res+=math.Mod(phi0-pp+1.5,1)-0.5
  return res
}

func ConeOpeningAngleZ() (float64,float64) {
  size := engine.Mesh().Size()
  th:=0.
  vari:=0.
  for ix:=0; ix<size[0]; ix++ {
    for iy:=0; iy<size[1]; iy++ {
      for iz:=0; iz<size[2]; iz++ {
        s:=engine.M.GetCell(ix,iy,iz)
        dummy:=math.Atan2(math.Sqrt(s.X()*s.X()+s.Y()*s.Y()),s.Z())
        th+=dummy
        vari+=dummy*dummy
      }
    }
  }
  th/=float64(size[0]*size[1]*size[2])
  vari/=float64(size[0]*size[1]*size[2])
  vari-=th*th
  return th,math.Sqrt(vari)
}

func ConeWindingsZ() [][]float64 {
  size:= engine.Mesh().Size()
  res := make([][]float64, size[0])
  for ix:=0; ix<size[0]; ix++ {
    res[ix]=make([]float64, size[1])
    for iy:=0; iy<size[1]; iy++ {
      s:=engine.M.GetCell(ix,iy,0)
      phi0:=math.Atan2(s.Y(),s.X())/(2*math.Pi)
      pp:=phi0
      for iz:=1; iz<size[2]; iz++ {
        t:=engine.M.GetCell(ix,iy,iz)
        np:=math.Atan2(t.Y(),t.X())/(2*math.Pi)
        res[ix][iy]+=math.Mod(np-pp+1.5,1)-0.5
        pp=np
      }
      res[ix][iy]+=math.Mod(phi0-pp+1.5,1)-0.5
    }
  }
  return res
}

func ConeWindingsZxy(ix, iy int) float64 {
  size:= engine.Mesh().Size()
  res := 0.
  s:=engine.M.GetCell(ix,iy,0)
  phi0:=math.Atan2(s.Y(),s.X())/(2*math.Pi)
  pp:=phi0
  for iz:=1; iz<size[2]; iz++ {
    t:=engine.M.GetCell(ix,iy,iz)
    np:=math.Atan2(t.Y(),t.X())/(2*math.Pi)
    res+=math.Mod(np-pp+1.5,1)-0.5
    pp=np
  }
  res+=math.Mod(phi0-pp+1.5,1)-0.5
  return res
}
