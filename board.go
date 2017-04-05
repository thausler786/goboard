package main 

const(
  EMPTY="+"
  BLACK="0"
  WHITE="X"
)


type Board struct {
  grid [19][19]string
  black_turn bool
}



func NewBoard() Board {
  board := new(Board)
  for y,row := range board.grid {
    for x,_ := range row {
      board.grid[y][x] = EMPTY
    }
  }
  return *board
}

func printBoard(board Board) Board {
  for _,row := range board.grid {
    for _,stone := range row {
      print(stone)
    }
    println()
  }
  return board
}

func move(board Board, move Move) Board {
  board.grid[move.y][move.x] = move.color
  return board
}
