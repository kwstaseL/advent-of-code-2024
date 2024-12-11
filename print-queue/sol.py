import requests
from dotenv import load_dotenv
import os

load_dotenv()

def solve(input_text):

    def parse_rules(rules):
        graph = {}
        for rule in rules:
            from_page, to_page = map(int, rule.split("|"))
            if from_page not in graph:
                graph[from_page] = []
            graph[from_page].append(to_page)
        return graph

    def is_topological_order(update, graph):
        position = {page: idx for idx, page in enumerate(update)}

        for from_page, neighbors in graph.items():
            if from_page in position:
                for to_page in neighbors:
                    if to_page in position and position[from_page] >= position[to_page]:
                        return False
        return True

    def parse_updates(updates):
        return [list(map(int, update.split(","))) for update in updates]
        
    sections = input_text.strip().split("\n\n")
    rules = sections[0].split("\n")
    updates = sections[1].split("\n")

    graph = parse_rules(rules)
    update_sequences = parse_updates(updates)

    total_sum = 0
    for update in update_sequences:
        if is_topological_order(update, graph):
            middle = update[len(update) // 2]
            total_sum += middle

    return total_sum

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

input_text = get_aoc_input(5)
result = solve(input_text)
print("Sum of middle pages of valid updates:", result)
