package main

type Validator interface {
  IsValid(Board, Move) bool
}

type validator struct {}

func NewValidator() Validator {
  return *new(validator)
}

func (v validator) IsValid(b Board, m Move) bool {
  if m.X > b.XDimen() || m.Y > b.YDimen() || m.X < 1 || m.Y < 1 {
    return false
  }
  if b.StoneAt(m.X, m.Y) != EMPTY {
    return false
  }
  return true
}
