package main

import(
  "os"
  "bufio"
)

func main() {
    board:= NewBoard()
    board.Print()
    println("hello!")
    for 1 < 2 {
      doMove(board)
      println("next move")
    }
}

func doMove(board Board) {
  println("Enter text:")
  reader := bufio.NewReader(os.Stdin)
  text,_ := reader.ReadBytes('\n')
  next_move := NewMove(text)
  board.Move(next_move)
  board.Print()
}


