from selenium import webdriver
import argparse
import json
import time

parser= argparse.ArgumentParser(description='Kakaostory Session Tooker')
parser.add_argument("-o", "--dest", help="session file save destination")
parser.add_argument("-d", "--driver-dest", help="set driver path")
args = parser.parse_args()

if not args.dest or not args.driver_dest:
    parser.print_usage()
    exit(-1)

driver = webdriver.Chrome(args.driver_dest)
driver.get('https://accounts.kakao.com/login/kakaostory')


while True:
    if driver.current_url == 'https://story.kakao.com/':
        with open(args.dest, 'w') as f:
            data = json.dumps(driver.get_cookies())
            f.write(data)
            driver.quit()
            exit(0)
        time.sleep(0.3)
