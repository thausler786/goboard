package main

import "encoding/json"

type Move struct {
  X int        `json:x`
  Y int        `json:y`
  Color string `json:color`
}

func NewMoveString(s string) Move {
  return NewMove([]byte(s))
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
