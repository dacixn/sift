# Sift
A lightweight CLI tool for sorting files to a user-defined structure
## ⚠️ Work in progress
This is my first "real" Go project, and I'm still learning as I go (no pun intended). Functionality is there but very limited at the moment; you can run Sift against a directory with `sift <dir>` and it will sort files to subdirectories by extension.

```yaml
exe/:
    VC_redist.x64.exe
    GitHubDesktopSetup-x64.exe
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