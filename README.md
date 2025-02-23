[![Go Reference](https://pkg.go.dev/badge/github.com/sjet47/go-curseforge.svg)](https://pkg.go.dev/github.com/sjet47/go-curseforge)

# go-curseforge

[CurseForge API](https://docs.curseforge.com) client for golang

## Supported API

### Mods

- [x] Search Mods
- [x] Get Mod
- [x] Get Mods
- [x] Get Featured Mods
- [x] Get Mod Description

### Files

- [x] Get Mod File
- [x] Get Mod Files
- [x] Get Files
- [x] Get Mod File Changelog
- [x] Get Mod File Download URL

### Minecraft

- [x] Get Minecraft Versions
- [x] Get Specific Minecraft Version
- [x] Get Minecraft ModLoaders
- [x] Get Specific Minecraft ModLoader

### Games

- [x] Get Games
- [x] Get Game
- [x] Get Versions
- [x] Get Version Types
- [x] Get Versions V2

### Categories

- [x] Get Categories

### Fingerprints

- [x] Get Fingerprints Matches By Game Id
- [x] Get Fingerprints Matches
- [x] Get Fingerprints Fuzzy Matches By Game Id
- [x] Get Fingerprints Fuzzy Matches

## Acknowledgements

This client library is developed totally in interest and for non-commercial usage, and may not be updated as frequently as the CurseForge API. If you find any inconsistent defines or error, please open an issue or a pull request.

Some ID types are intentionally leaved as enums since they may be useful while using the API(i.e. `enum.GameID`), but they are not documented in the official documentation. Currently I'm focusing on the APIs so there is only one GameID, which is `432` for Minecraft cause it's the only one I used. After I'm done with the API stuff I'll try to do some code generation for those enums. Also, any PRs are welcome.
