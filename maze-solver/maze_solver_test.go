package maze_solver

import "testing"

func TestMazeSolver(t *testing.T) {
	maze := [][]rune{
		{'#', '#', '#', '#', '#', 'X', '#'},
		{'#', ' ', ' ', ' ', '#', ' ', '#'},
		{'#', ' ', '#', ' ', ' ', ' ', '#'},
		{'#', ' ', '#', '#', '#', '#', '#'},
	}
  start := Point{1, 3}
  end := Point{5, 0}

  path, err := Solve(maze, start, end)

  if err != nil {
    t.Errorf("Solve(%v, %v, %v) == (%v, %v), expected ([]Point, nil)", maze, start, end, path, err)
  }

  correct_path := []Point{
    {1, 3},
    {1, 2},
    {1, 1},
    {2, 1},
    {3, 1},
    {3, 2},
    {4, 2},
    {5, 2},
    {5, 1},
    {5, 0},
  }

  for i, point := range correct_path {
    if point != path[i] {
      t.Errorf("")
    }
  }
}
