# bilibili-download

Download videos from bilibili.com supporting flv/dash(need post processing)/mp4.

支持从bilibili下载视频，支持flv/dash(需要用ffmpeg后处理)/mp4格式。

## Usage

```shell
bilibili-download <video url>
```

Example:

```
bilibili-download https://www.bilibili.com/video/BV1bs411s7kR
```

## FLV

Video will be downloaded into current directory named `video.flv`.

直接在当前目录生成`video.flv`文件。

## MP4

Video will be downloaded into current directory named `video.mp4`.

直接在当前目录生成`video.mp4`文件。

## Dash

* Video will be downloaded into current directory named `video.m4s`.
* Audio will be downloaded into current directory named `audio.m4s`.

If you want to merge them together using `ffmpeg` you can try:

```shell
ffmpeg -i audio.m4s -i video.m4s -c:v copy output.mp4
```

直接在当前目录生成`video.m4s`/`audio.m4s`文件，分别代表无声视频和音频。

如果系统安装有`ffmpeg`那么可以使用以下命令来合并为一个文件：

```shell
ffmpeg -i audio.m4s -i video.m4s -c:v copy output.mp4
```
