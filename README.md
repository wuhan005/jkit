# 🔧 jkit [![Go](https://github.com/wuhan005/jkit/actions/workflows/go.yml/badge.svg)](https://github.com/wuhan005/jkit/actions/workflows/go.yml)

JSON CLI Tool

## Install

```bash
git clone https://github.com/wuhan005/jkit.git
cd jkit
go install
```

## Usage

Copy your JSON first, then enjoy it.

------

### `jkit f` Format JSON

[Input JSON](https://i.apicon.cn/chunzhen/?ip=52.68.96.58)

```bash
> jkit f

{
    "data": {
        "area": "东京Amazon数据中心",
        "country": "日本",
        "ip": "52.68.96.58"
    },
    "error": 0,
    "msg": "success"
}
```

### `jkit c <deep>` Fold JSON

[Input JSON](https://api.bilibili.com/pgc/web/season/section?season_id=34412)

```bash
> jkit c 4

{
    "code": 0.000000,
    "message": "success",
    "result": {
        "main_section": {
            "type": 0.000000,
            "episodes": [
                { 14 items dict },
                { 14 items dict },
                { 14 items dict },
                { 14 items dict },
                { 14 items dict },
            ],
            "id": 49914.000000,
            "title": "正片"
        },
        "section": [
            {
                "id": 50112.000000,
                "title": "PV",
                "type": 2.000000,
                "episodes": [ 3 items array ]
            }
        ]
    }
}
```

### `jkit g` Get JSON element

[Input JSON](https://api.bilibili.com/pgc/web/season/section?season_id=34412)

```bash
> jkit g result.main_section.episodes.0

{
    "badge": "会员",
    "badge_info": {
        "bg_color": "#FB7299",
        "bg_color_night": "#BB5B76",
        "text": "会员"
    },
    "cover": "http://i0.hdslb.com/bfs/archive/34a44b83a9c657a9d096d85771334f545e32dd17.jpg",
    "from": "bangumi",
    "id": 341208,
    "is_premiere": 0,
    "status": 13,
    "title": "1",
    "vid": "",
    "badge_type": 0,
    "long_title": "见习魔女伊蕾娜",
    "aid": 329827456,
    "cid": 281989571,
    "share_url": "https://www.bilibili.com/bangumi/play/ep341208"
}
```

### `jkit m` Generate JSON form list

Input from my [GitHub profile README](https://github.com/wuhan005/wuhan005):

```bash
> jkit m

[
    "Hi, I'm E99p1ant. 🍆",
    "🐭 Focus on Golang.",
    "🏠 Blog at github.red.",
    "💬 Ask me something?",
    "🤤 Buy me a cup of coffee.",
    "Some cool gadgets I made:",
    "",
    "NekoBox - 匿名提问箱 / Anonymous Question Box",
    "Apicon - API 热爱者"
]
```