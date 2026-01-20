# yst-img

Fast, content-aware image converter & compressor for Yeast.

## Features

- Convert and compress images
- Automatic batch mode (just pass a directory)
- Parallel processing (default 4 workers)
- Smart auto-quality (content-aware)
- AVIF (AV1) encode & decode
- Clean CLI output

## Installation

```bash
yst plugins install kaustubha-chaturvedi:yst-img
```

## Usage

### Convert
```bash
yst-img convert input.png output.avif
```

### Compress
```bash
yst-img compress ./images ./out
```

### Batch options
```bash
yst-img compress ./images ./out --workers 8 --quality 60
```

### Auto mode(default)
```bash
yst-img convert input.jpg output.avif
```


### Supported formats

- Input: PNG, JPEG, WEBP
- Output: PNG, JPEG, WEBP, AVIF

#### THIS README WAS WRITTEN BY [QWEN](https://chat.qwen.ai/) 