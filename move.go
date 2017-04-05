package main

import "encoding/json"

type Move struct {
  x int        `json:x`
  y int        `json:y`
  color string `json:color`
}

func NewMove(b []byte) Move {
  var dat map[string]interface{}
  if err := json.Unmarshal(b, &dat); err != nil {
    println("you wasted your turn")
  }
  y := int(dat["y"].(float64))
  x := int(dat["x"].(float64))
  color := dat["color"].(string)
  return Move{x, y, color}
}
