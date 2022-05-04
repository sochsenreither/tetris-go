# Tetris

A simple Go implementation of Tetris clone with an AI based on the algorithm described [here](https://codemyroad.wordpress.com/2013/04/14/tetris-ai-the-near-perfect-player/)\
The board size can be changed in `game/board.go`

![prev_normal_board](https://user-images.githubusercontent.com/29070949/166808811-5cc5aa8e-3d37-4517-b26b-f0c9994be216.png)
![preview_large_board](https://user-images.githubusercontent.com/29070949/166808815-3e4db5e6-f45a-433f-a910-65466dba3e18.png)


## Controls
- Up arrow: Rotate piece
- Down/Left/Right: Move piece
- Space: Drop piece
- P: pause
- Esc: exit game

## Running the Game
Use `go run .` from the root directory to run tetris. Use the `-ai` flag to let the ai play `go run . -ai`.
