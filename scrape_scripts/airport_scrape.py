import requests
from bs4 import BeautifulSoup
import pandas as pd

# Target page (your example)
url = "https://en.wikipedia.org/wiki/Indianapolis_International_Airport"

# Fetch the page
resp = requests.get(url, headers={"User-Agent": "Mozilla/5.0"})
resp.encoding = "utf-8"
soup = BeautifulSoup(resp.text, "html.parser")

# 1) Find the <h3 id="Passenger"> heading (or fallback if the id sits on a nested span)
passenger_h3 = soup.find("h3", id="Passenger")
if not passenger_h3:
    span = soup.find("span", id="Passenger")
    if span:
        passenger_h3 = span.find_parent("h3")

if not passenger_h3:
    raise RuntimeError("Passenger section <h3> not found.")

# 2) Get the first wikitable after that heading
table = passenger_h3.find_next("table", class_="wikitable")
if table is None:
    raise RuntimeError("No 'wikitable' found after the Passenger heading.")

print(table)
