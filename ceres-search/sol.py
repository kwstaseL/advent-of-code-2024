import requests
from dotenv import load_dotenv
import os

load_dotenv()

def countXMAS(grid):
    ROWS, COLS = len(grid), len(grid[0])
    word = "XMAS"
    word_len = len(word)
    res = 0
    
    directions = [(0, 1), (0, -1), (1, 0), (-1, 0), (1, 1), (1, -1), (-1, 1), (-1, -1)]
    
    def check_direction(r, c, dr, dc):
        for i in range(word_len):
            nr, nc = r + dr * i, c + dc * i
            if nr < 0 or nr >= ROWS or nc < 0 or nc >= COLS or grid[nr][nc] != word[i]:
                return False
        return True
    
    for r in range(ROWS):
        for c in range(COLS):
            for dr, dc in directions:
                if check_direction(r, c, dr, dc):
                    res += 1
    
    return res

def get_aoc_input(day, year=2024):
    session_cookie = os.getenv("AOC_SESSION_COOKIE")
    cookies = {
        "session": session_cookie
    }
    url = f"https://adventofcode.com/{year}/day/{day}/input"
    response = requests.get(url, cookies=cookies)
    if response.status_code == 200:
        puzzle_input = response.text.strip()
        return puzzle_input
    else:
        raise ValueError(f"Failed to get input. Status code: {response.status_code}")

grid = get_aoc_input(4, 2024).split("\n")
grid = [list(row) for row in grid]

res = countXMAS(grid)
print(f"Number of occurrences of XMAS: {res}")

