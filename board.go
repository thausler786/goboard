package main 

const(
  EMPTY="+"
  BLACK="0"
  WHITE="X"
)

type Board interface {
  Print() Board
  Move(Move) Board
  StoneAt(int, int) string
}


type memoryBoard struct {
  grid [19][19]string
  black_turn bool
}



func NewBoard() Board {
  board := new(memoryBoard)
  for y,row := range board.grid {
    for x,_ := range row {
      board.grid[y][x] = EMPTY
    }
  }
  return *board
}

func (board memoryBoard) Print() Board {
  for _,row := range board.grid {
    for _,stone := range row {
      print(stone)
    }
    println()
  }
  return board
}

func (board memoryBoard) StoneAt(x int, y int) string {
  return board.grid[y][x]
}

func (board memoryBoard) Move(move Move) Board {
  board.grid[move.y][move.x] = move.color
  return board
}
