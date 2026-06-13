# Sift
A lightweight CLI tool for sorting files to a user-defined structure

## Building
In order to build this project, you need to [install Go](https://go.dev/doc/install)

```bash
git clone https://github.com/dacixn/sift
cd sift
go build ./cmd/sift
```
## Usage
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



