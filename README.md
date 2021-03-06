# runproj

CLI tool to open predefinied sets of Programs for example VS Code - specific project, Github project page)

The configuration should be given in `runproj.json`.
TBD - where to save the `runproj.json`

Program execution: `runproj <set-name>` or `runproj <set1-name> <set2-name>`

---

### Configuration

```json
[
    {
        "name": "<set1-name>",
        "programs": [
            {
                "program": "<program1-url>",
                "arguments": ["<argument1>"]
            },
            {
                "program": "<program2-url>",
                "arguments": ["<argument2>", "<argument3>"]
            }
        ]
    },
    {
        "name": "<set2-name>",
        "programs": [
            {
                "program": "<program3-url>",
                "arguments": ["<argument4>"]
            },
            {
                "program": "<program4-url>",
                "arguments": ["<argument5>", "<argument6>"]
            }
        ]
    }
]
```

#### Example:

```json
[
    {
        "name": "timeo flutter",
        "programs": [
            {
                "program": "C:\\Users\\Feka\\AppData\\Local\\Programs\\Microsoft VS Code\\Code.exe",
                "arguments": ["D:\\2_Projekte\\Flutter\\timeo"]
            },
            {
                "program": "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
                "arguments": [
                    "--new-window",
                    "https://github.com/FlorianFeka/timeo-app"
                ]
            }
        ]
    },
    {
        "name": "timeo angular",
        "programs": [
            {
                "program": "C:\\Users\\Feka\\AppData\\Local\\Programs\\Microsoft VS Code\\Code.exe",
                "arguments": ["D:\\2_Projekte\\Javascript\\Angular\\timeo"]
            },
            {
                "program": "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
                "arguments": [
                    "--new-window",
                    "https://github.com/FlorianFeka/timeo"
                ]
            }
        ]
    }
]
```
