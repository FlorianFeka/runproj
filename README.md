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
				"path": "<path1>"
			},
			{
				"program": "<program2-url>",
				"path": "<path2>"
			}
		]
	},
	{
		"name": "<set2-name>",
		"programs": [
			{
				"program": "<program3-url>",
				"path": "<path3>"
			},
			{
				"program": "<program4-url>",
				"path": "<path4>"
			}
		]
	}
]
```

#### Example:
TBD