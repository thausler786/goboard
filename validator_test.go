package main_test

import (
	. "github.com/thausler786/goboard"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validator", func() {

  var (
    board Board
    validator Validator
  )

  Describe("when it validates", func() {
    BeforeEach(func() {
      validator = NewValidator()
      board = NewBoard(validator)
    })

    It("does not allow moves out of bounds", func() {
      move1 := NewMoveString("{\"x\": 0, \"y\": 2, \"color\": \"X\"}")
      move2 := NewMoveString("{\"x\": 20, \"y\": 2, \"color\": \"X\"}")
      move3 := NewMoveString("{\"x\": 2, \"y\": 0, \"color\": \"X\"}")
      move4 := NewMoveString("{\"x\": 2, \"y\": 20, \"color\": \"X\"}")
      validMove := NewMoveString("{\"x\": 2, \"y\": 2, \"color\": \"X\"}")
      Expect(validator.IsValid(board, move1)).To(BeFalse())
      Expect(validator.IsValid(board, move2)).To(BeFalse())
      Expect(validator.IsValid(board, move3)).To(BeFalse())
      Expect(validator.IsValid(board, move4)).To(BeFalse())
      Expect(validator.IsValid(board, validMove)).To(BeTrue())

    })

    It("does not allow moves with the wrong stone color", func() {
      move1 := NewMoveString("{\"x\": 2, \"y\": 2, \"color\": \"O\"}")
      Expect(validator.IsValid(board, move1)).To(BeFalse())
    })

    It("does not allow placing a piece on another piece", func() {
      move1 := NewMoveString("{\"x\": 2, \"y\": 2, \"color\": \"X\"}")
      board := board.Move(move1)
      move2 := NewMoveString("{\"x\": 2, \"y\": 2, \"color\": \"O\"}")
      Expect(validator.IsValid(board, move2)).To(BeFalse())
    })
  })
})
