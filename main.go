package main

import (
	"fmt"
	"os"

	"github.com/jasonjoo2010/bilibili-download/util"
)

// Currently default to download the video in quality 480p(which doesn't require to login).
//	If it doesn't exist then low down in descending order.

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: bilibili-download <video url>")
	fmt.Fprintln(os.Stderr, "Example: bilibili-download https://www.bilibili.com/video/BV1bs411s7kR")
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	url := os.Args[1]

	var (
		err         error
		output_file string
	)
	info, err := util.GetPlayerInfo(url)
	if err != nil {
		goto OUTPUT_ERROR
	}

	if len(info.VideoUrls) > 0 {
		fmt.Println("Format:", info.Format)
		output_file = "video." + util.GetExt(info.VideoUrls[0].URL)
		err = util.Download(info.VideoUrls[0].URL, output_file)
		if err != nil {
			goto OUTPUT_ERROR
		}
	} else if info.Streams != nil {
		fmt.Println("Detected mixed streams")
		var (
			audioStream = info.Streams.Audios[0]
			videoStream = info.Streams.Videos[0]
		)
		for _, s := range info.Streams.Videos {
			if s.Codec != 7 {
				continue
			}
			if videoStream.Codec != 7 {
				videoStream = s
				continue
			}
			if s.Quality > 64 {
				continue
			}
			if s.Quality > videoStream.Quality || videoStream.Quality > 64 {
				videoStream = s
				continue
			}
		}
		for _, s := range info.Streams.Audios {
			if s.Codec != 30216 {
				continue
			}
			audioStream = s
		}
		fmt.Println("Download video stream ... ")
		output_file = "video." + util.GetExt(videoStream.URL)
		err = util.Download(videoStream.URL, output_file)
		if err != nil {
			goto OUTPUT_ERROR
		}
		fmt.Println("Download audio stream ... ")
		output_file = "audio." + util.GetExt(audioStream.URL)
		err = util.Download(audioStream.URL, output_file)
		if err != nil {
			goto OUTPUT_ERROR
		}
	}

	return

OUTPUT_ERROR:
	fmt.Fprintln(os.Stderr, "Error:", err.Error())
	os.Exit(1)

}
