package main

import (
	"fmt"
	"math"
)

func findFloorWhereTheyBreak(floors []bool) (int, error){

  numFloors := len(floors)
  jumpOffset := int(math.Sqrt(float64(numFloors)))

  floor := 0
  jump := 0

  for{
    floor = floor + jump*jumpOffset
    if floors[floor]{
      //broke, backtrack...
      break
    }
    jump += 1
  }

  for i:=jump*jumpOffset+1; i<floor; i++{
    if floors[i]{
      return i-1, nil
    }
  }

  return -1, fmt.Errorf("not found")
}

func main(){

  floors := []bool{false, false, false, false, false, false, true, true, true, true};
  res, err := findFloorWhereTheyBreak(floors)
  if err != nil{
    fmt.Errorf("error: %s", err)
  }
  fmt.Printf("floor where they break: %d", res)

}
