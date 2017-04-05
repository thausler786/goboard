package main

import(
  "os"
  "bufio"
)

func main() {
    board:= NewBoard()
    printBoard(board)
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
  move(board, next_move)
  printBoard(board)
}


