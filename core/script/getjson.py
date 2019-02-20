import requests
import json
import uuid
import os

data = {
    "GetProfileFeed":"https://story.kakao.com/a/profiles/lproject",
    "GetProfile":"https://story.kakao.com/a/profiles/lproject?with=activities",
    "GetProfileRelationship": "https://story.kakao.com/a/profiles/lproject?profile_only=true",
    "GetFeed" : "https://story.kakao.com/a/feeds",
    "GetBookmark":"https://story.kakao.com/a/profiles/lproject/sections/bookmark",
    "GetFriendData": "https://story.kakao.com/a/friends/",
    "GetComment": "https://story.kakao.com/a/activities/_fVJf69.hWksWZ1aKMA/comments?lpp=30&order=desc",
    "GetProfileData" :"https://story.kakao.com/a/settings/profile",
    "GetSpecificFriends":"https://story.kakao.com/a/activities/_fVJf69.hWksWZ1aKMA/specific_friends"
}

def main():
    for name in data.keys():
        jar = requests.cookies.RequestsCookieJar()
        with open(r"C:\Users\Dev\go\src\github.com\git-soo\Kakaostory-Manager-Cli\core\script\session_test.json") as f:
            cookies = json.loads(f.read())
            

        headers = {
            "Content-Type":"application/x-www-form-urlencoded; charset=utf-8",
            "X-Kakao-ApiLevel":"46",
            "X-Kakao-VC":str(uuid.uuid4()).replace("-", "")[:20],
            "X-Kakao-DeviceInfo":"web:d;-;-",
            "X-Requested-With": "XMLHttpRequest",
            "DNT":"1",
            "Accept":"application/json",
            "Accept-Encoding":"gzip, deflate, br",
            "Accept-Language":"ko",
            "Cache-Control":"no-cache",
            "Referer":"https://story.kakao.com/",
            "Host":"story.kakao.com",
            "User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Safari/537.36"
        }

        session = requests.Session()
        for cookie in cookies: session.cookies.set(cookie['name'], cookie['value'])

        r = session.get(data[name], headers=headers)
        with open(os.path.join(r"C:\Users\Dev\go\src\github.com\git-soo\Kakaostory-Manager-Cli\core\json", name + ".json"), "w") as f:
            f.write(json.dumps(r.json()))
            print("Done", name)

if __name__ == "__main__":
    main()