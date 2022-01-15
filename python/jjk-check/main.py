# This is a dead simple script which checks for new Jujutsu Kaisen chapters

import json
import os

import requests
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.firefox.options import Options
from selenium.webdriver.firefox.service import Service
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.ui import WebDriverWait

WEBHOOK = os.environ["WEBHOOK_URL"]

options = Options()
options.headless = True

driver = webdriver.Firefox(options=options, service=Service(log_path="/dev/null"))
driver.get("https://mangaplus.shueisha.co.jp/titles/100034")


# Asserts that data files are written on the script's folder
os.chdir(os.path.dirname(os.path.abspath(__file__)))

try:
    with open("data.json", "r") as data:
        last_chapter = json.load(data)
except FileNotFoundError:
    last_chapter = {"chapter": None}

try:
    element = WebDriverWait(driver, 0).until(
        EC.presence_of_element_located(
            (
                By.CSS_SELECTOR,
                "div.ChapterListItem-module_chapterListItem_ykICp:nth-child(9) > div:nth-child(1) > p:nth-child(2)",
            )
        )
    )

    # All good, no new chapter.
    if last_chapter["chapter"] == element.text:
        exit(0)

    resp = requests.post(
        WEBHOOK,
        json={"content": f"Chapter {element.text} of Jujutsu Kaisen has been released"},
    )
    assert resp.status_code == 204

    # Only writes on data file if the webhook was successfully posted

    last_chapter["chapter"] = element.text
    with open("data.json", "w") as data:
        data.write(json.dumps(last_chapter))

except Exception as e:
    msg = str(e)
    requests.post(
        WEBHOOK,
        json={"content": f"Failed to check for a new Jujutsu Kaisen chapter: {msg}"},
    )
