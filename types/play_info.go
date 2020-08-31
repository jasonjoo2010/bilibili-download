package types

type VideoURL struct {
	Size    int      `json:"size"`
	Order   int      `json:"order"`
	URL     string   `json:"url"`
	Backups []string `json:"backup_url"`
}

type MediaStream struct {
	Quality   int      `json:"id"`
	Codec     int      `json:"codecid"`
	Codecs    string   `json:"codecs"`
	MimeType  string   `json:"mime_type"`
	FrameRate string   `json:"frame_rate"`
	Width     int      `json:"width"`
	Height    int      `json:"height"`
	URL       string   `json:"baseUrl"`
	Backups   []string `json:"backupUrl"`
}

type MediaStreams struct {
	Videos []MediaStream `json:"video"`
	Audios []MediaStream `json:"audio"`
}

type PlayInfo struct {
	Quality     []int         `json:"accept_quality"`
	QualityDesc []string      `json:"accept_description"`
	Length      int           `json:"timelength"`
	Codec       int           `json:"video_codecid"`
	Format      string        `json:"format"`
	VideoUrls   []VideoURL    `json:"durl"`
	Streams     *MediaStreams `json:"dash"`
}

type Response struct {
	Code    int      `json:"code"`
	Msg     string   `json:"message"`
	Session string   `json:"session"`
	TTL     int      `json:"ttl"`
	Data    PlayInfo `json:"data"`
}
