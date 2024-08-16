package maze_solver

import (
  "errors"
)

type Point struct {
  x, y int
}

var directions = [4]Point{
  {-1, 0},
  {0, -1},
  {1, 0},
  {0, 1},
}

type Maze [][]rune;

func Solve(maze Maze, start Point, end Point) ([]Point, error) {
  stack := make([]Point, 0)
  visited := make(map[Point]bool)

  if _solve(maze, end, start, &stack, visited) {
    return stack, nil
  }

  return make([]Point, 0, 0), errors.New("no path found")
}

func _solve(maze Maze, end Point, current Point, stack *[]Point, visited map[Point]bool) bool {
  if (current.y >= len(maze) || current.y < 0) || (current.x >= len(maze[0]) || current.x < 0) {
    return false
  }

  if visited[current] {
    return false
  }

  if maze[current.y][current.x] == '#' {
    return false
  }

  if current == end {
    *stack = append(*stack, current)
    return true
  }

  visited[current] = true
  *stack = append(*stack, current)

  for _, d := range directions {
    next := Point{x: current.x + d.x, y: current.y + d.y}

    if _solve(maze, end, next, stack, visited) {
      return true
    }
  }

  *stack = (*stack)[:len(*stack)-1]
  return false
}
