# Paneless

Painless Window Layout Management

![Sample Screenshot](https://i.imgur.com/JOCL93T.jpg)

Stemming from frustrations with current window-snapping utilities in Windows when using an ultrawide monitor, Paneless was written to allow pixel perfect management of semi-complex window arrangements.

# Defining Window Arrangements

Window arrangements are defined using a JSON file.  Multiple arrangments can be specified in the file and can be switched from the taskbar when the application is running.  To achieve the screenshot above, the JSON file would like like...

```json
[
    {
        "name": "Watching Twitch - App",
        "preferences": [
            {
                "NameRegex": "Twitch",
                "NameExlusionRegex": "",
                "X": 1895,
                "Y": 500,
                "Cx": 1545,
                "Cy": 900
            },
            {
                "NameRegex": ".* - Google Chrome",
                "NameExlusionRegex": "",
                "X": -7,
                "Y": 0,
                "Cx": 1909,
                "Cy": 1407
            },
            {
                "NameRegex": "Friends",
                "NameExlusionRegex": "",
                "X": 1895,
                "Y": 0,
                "Cx": 325,
                "Cy": 500
            },
            {
                "NameRegex": ".* - Discord",
                "NameExlusionRegex": "",
                "X": 2220,
                "Y": 0,
                "Cx": 1220,
                "Cy": 500
            }
        ]
    }
]
```

Windows are distinguished via their title bar name and can be found using a regex for both inclusion and exclusion.  For example, before I migrated to the desktop app for Twitch, I used to use a layout containing two Chrome windows...

```json
[
    {
        "name": "Watching Twitch - Web Player",
        "preferences": [
            {
                "NameRegex": "Twitch - Google Chrome",
                "NameExlusionRegex": "",
                "X": 1828,
                "Y": 390,
                "Cx": 1618,
                "Cy": 1017
            },
            {
                "NameRegex": ".* - Google Chrome",
                "NameExlusionRegex": ".*Twitch.*",
                "X": -7,
                "Y": 0,
                "Cx": 1850,
                "Cy": 1407
            },
            {
                "NameRegex": "Friends",
                "NameExlusionRegex": "",
                "X": 1835,
                "Y": 0,
                "Cx": 323,
                "Cy": 500
            },
            {
                "NameRegex": ".* - Discord",
                "NameExlusionRegex": "",
                "X": 2155,
                "Y": 0,
                "Cx": 1285,
                "Cy": 500
            }
        ]
    }
]
```

This snippet also highlights that window layouts are defined by pixel positions, not snap zones.  This layout positioned both Discord and the Steam Friends list ontop of the Chrome tab/URL bar of the Twitch web player resulting in as much screen space from the player as possible.

The "Get Current" option in the menu can help retrieve the current JSON representation of your screen to facilitate creating these arrangements.