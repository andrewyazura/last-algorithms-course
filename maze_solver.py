class Square:
    WALL = "#"
    PATH = " "

Point = tuple[int, int]
Path = list[Point]

Map = list[list[str]]
SeenMap = list[list[bool]]

directions = (
    (-1, 0),
    (0, -1),
    (1, 0),
    (0, 1),
)

def walk(maze: Map, current: Point, seen: SeenMap, end: Point, path: Path) -> bool:
    # off grid
    if (current[0] < 0 or current[0] > len(maze[0])) or (current[1] < 0 or current[1] > len(maze)):
        return False

    # already visited square
    if seen[current[1]][current[0]] is True:
        return False
    
    # ran into a wall
    if maze[current[1]][current[0]] == Square.WALL:
        return False

    # found end
    if current[0] == end[0] and current[1] == end[1]:
        path.append(current)
        return True

    # pre - push to stack
    seen[current[1]][current[0]] = True
    path.append(current)

    # recurse
    for direction in directions:
        next_point = (current[0] + direction[0], current[1] + direction[1])

        if walk(maze, next_point, seen, end, path):
            return True

    # post - pop from stack
    path.pop()

    # dead-end - return to the previous step
    return False

def solve(maze: Map, start: Point, end: Point) -> Path:
    path = []
    seen = [[False] * len(maze[0]) for _ in range(len(maze))]

    walk(maze, start, seen, end, path)

    return path

if __name__ == "__main__":
    maze = [
        ["#", "#", "#", "#", "#", "E", "#",],
        ["#", " ", " ", " ", " ", " ", "#",],
        ["#", "S", "#", "#", "#", "#", "#",],
    ]
    start = (1, 2)
    end = (5, 0)

    path = solve(maze, start, end)
    print(f"{start=}")
    print(f"{end=}")
    print(f"{path=}")
