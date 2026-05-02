# Sift
A lightweight CLI tool for sorting files to a user-defined structure
## ⚠️ Work in progress
This is my first "real" Go project, and I'm still learning as I go (no pun intended). Functionality is there but very limited at the moment; you can run Sift against a directory with `sift <dir>` and it will sort files to subdirectories by extension.

```yaml
exe/:
    GitHubDesktopSetup-x64.exe
    VC_redist.x64.exe
jpg/:
    flowers.jpg
pdf/:
    2025_Report.pdf
```

## Planned functionality

Define groups and wildcards in config.toml:
```toml
[groups]
"video" = ["*.mp4", "*.mkv"]
"video/iPhone" = ["MOV*"]
"video/iPhone/cinematic" = ["*ProRes*"]
"audio" = ["*.mp3"]
```
Sift will apply the structure to unsorted files:
```yaml
video/:
    vacation.mp4
    iPhone/:
        MOV_0842.mov
        cinematic/:
            sunset_ProRes_422.mov
audio/:
    deftones.mp3
```

### Other features
* **Dry run**: Run without changing any files, possibly in combination with tree or log file
* **File tree**: Export a `tree`-like overview of file changes
* **Config flag**: Specify which config file to use with a command argument. (As of now, Sift checks the working directory then user configuration directory independent of OS.)

### Notes and to-do
In order of relevance:
* (Re)write group-parsing logic 
* Update sorting logic to support wildcards
* Refactor code to include an app struct in order to pass flags to functions more easily (once I implement flags)
* ignoreDirs flag in config



