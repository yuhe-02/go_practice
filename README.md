# Image Converter

This package provides a command-line tool to convert JPG images to PNG format. 

## desciption


## Features

- Convert JPG images to PNG format, including files in subdirectories.
- Skip invalid files and report errors.
- Easy-to-use CLI interface.

## Installation

1. Ensure you have Go installed.
2. Clone this repository.
3. Run `go mod tidy` to install dependencies.

## Usage

```bash
go build
./convert <directory>
```

## memo
go run *.go = ビルドと実行
go build -o ### *.go = ビルド
go mod init "**/~" = go modules作成

## わからないこと
・出力ファイル名をデフォルトは、mainが書かれているファイル名になる？
・画像の変換方法
・go modulesとは、どう書くのが普通
・go doc？
・使っていいパッケージ（自作、標準、準標準）はどう判断？
・ユーザー定義型？
・goはOOPS的に書くのか、関数型的に書くのかわからん
