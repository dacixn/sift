# Sift
A lightweight CLI tool for sorting files to a user-defined structure
## ⚠️ Work in progress
This is my first "real" Go project, and I'm still learning as I go (no pun intended). Functionality is there but very limited at the moment; you can run Sift against a directory with `sift <dir>` and it will create subdirectories respective to the file extensions of unsorted files.

## Planned functionality

Define groups and wildcards in config.toml
```
[groups]
"video" = ["*.mp4", "*.mkv"]
"video/iPhone" = ["MOV*"]
"video/iPhone/cinematic" = ["*ProRes*"]
"audio" = ["*.mp3"]
```
Sift will apply the structure to unsorted files
```
video/
    └ iPhone/
        └ cinematic/
            └ sunset_ProRes_422.mov
        └ MOV_0842.mov
    └ vacation.mp4
audio/
    └ deftones.mp3
```