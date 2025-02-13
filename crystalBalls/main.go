package main

import (
	"fmt"
	"math"
)

func findFloorWhereTheyBreak(floors []bool) (int, error){

  numFloors := len(floors)
  jumpOffset := int(math.Sqrt(float64(numFloors)))


  floor := jumpOffset

  for floor < numFloors{
    if floors[floor]{
      break
    }
    floor += jumpOffset
  }

  floor -= jumpOffset
  for i:=floor; i<=floor+jumpOffset && i < numFloors; i++{
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
  fmt.Printf("last floor where they won't break: %d", res)

}
