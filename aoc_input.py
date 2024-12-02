import requests
from dotenv import load_dotenv
import os

def get_aoc_input(day, year=2024):
    load_dotenv()

    session_cookie = os.getenv("AOC_SESSION_COOKIE")
    
    if session_cookie is None:
        raise ValueError("AOC_SESSION_COOKIE environment variable not set")
    
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