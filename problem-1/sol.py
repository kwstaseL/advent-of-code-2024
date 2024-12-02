import sys
import os
import heapq

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..'))) 
from aoc_input import get_aoc_input


def calculate_total_distance():
    input_data = get_aoc_input(day=1)

    left_list = []
    right_list = []

    # Prepare the left and right list
    for line in input_data.splitlines():
        left, right = map(int, line.split())
        left_list.append(left)
        right_list.append(right)

    heapq.heapify(left_list)
    heapq.heapify(right_list)

    res = 0

    while left_list and right_list:
        left_min = heapq.heappop(left_list)
        right_min = heapq.heappop(right_list)
        res += abs(left_min - right_min)

    print(f"Total distance: {res}")

calculate_total_distance()