package main_test

import (
	. "github.com/thausler786/goboard"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {

  var (
    board Board
    move Move
  )

  Describe("Moving on the board", func() {
    BeforeEach(func() {
      board = NewBoard(NewValidator())
      move = NewMoveString("{\"x\": 1, \"y\": 2, \"color\": \"H\"}")
    })

    It("Modfies the board", func() {
      Expect(board.Move(move).StoneAt(1, 2)).To(Equal("H"))
    })
  })

  Describe("Validating the move", func() {
    It("Calls the validator to validate its move", func() {
    })
  })

  Describe("Adjudicating the move", func() {

    It("removes individual stones that are dead", func() {
    })

    It("adjudicates the movee's stones before the mover's", func () {})
  })


})
