package gogiphy

type Images struct {
	FixedHeight            Fixed            `json:"fixed_height"`
	FixedHeightStill       StillImage       `json:"fixed_height_still"`
	FixedHeightDownsampled FixedDownsampled `json:"fixed_height_downsampled"`

	FixedWidth            Fixed            `json:"fixed_width"`
	FixedWidthStill       StillImage       `json:"fixed_width_still"`
	FixedWidthDownsampled FixedDownsampled `json:"fixed_width_downsampled"`

	FixedHeightSmall      Fixed      `json:"fixed_height_small"`
	FixedHeightSmallStill StillImage `json:"fixed_height_small_still"`
	FixedWidthSmall       Fixed      `json:"fixed_width_small"`
	FixedWidthSmallStill  StillImage `json:"fixed_width_small_still"`

	Downsized       Downsized      `json:"downsized"`
	DownsizedStill  StillImage     `json:"downsized_still"`
	DownsizedLarge  Downsized      `json:"downsized_large"`
	DownsizedMedium Downsized      `json:"downsized_medium"`
	DownsizedSmall  DownsizedSmall `json:"downsized_small"`

	Original      Original   `json:"original"`
	OriginalStill StillImage `json:"original_still"`

	Looping    Looping    `json:"looping"`
	Preview    Preview    `json:"preview"`
	PreviewGif StillImage `json:"preview_gif"`
}

type StillImage struct {
	url    string
	width  string
	height string
}

type FixedDownsampled struct {
	URL      string `json:"url"`
	Width    string `json:"width"`
	Height   string `json:"height"`
	Size     string `json:"size"`
	Webp     string `json:"webp"`
	WebpSize string `json:"webp_size"`
}

type Fixed struct {
	URL      string `json:"url"`
	Width    string `json:"width"`
	Height   string `json:"height"`
	Size     string `json:"size"`
	Mp4      string `json:"mp4"`
	Mp4Size  string `json:"mp4_size"`
	Webp     string `json:"webp"`
	WebpSize string `json:"webp_size"`
}

type Downsized struct {
	URL    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Size   string `json:"size"`
}

type DownsizedSmall struct {
	Mp4     string `json:"mp4"`
	Width   string `json:"width"`
	Height  string `json:"height"`
	Size    string `json:"size"`
	Mp4Size string `json:"mp4_size"`
}

type Original struct {
	Width    string `json:"width"`
	Height   string `json:"height"`
	Size     string `json:"size"`
	Frames   string `json:"frames"`
	Mp4      string `json:"mp4"`
	Mp4Size  string `json:"mp4_size"`
	Webp     string `json:"webp"`
	WebpSize string `json:"webp_size"`
}

type Looping struct {
	Mp4 string `json:"mp4"`
}

type Preview struct {
	Mp4     string `json:"mp4"`
	Mp4Size string `json:"mp4_size"`
	Width   string `json:"width"`
	Height  string `json:"height"`
}
