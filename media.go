package main

type Media struct {
	Media_key   string
	Type        string
	Url         string
	Duration_ms int
	Height      int
	// TODO: non_public_metrics
	// TODO: organic_metrics
	Preview_image_url string
	// TODO: promoted_metrics
	// TODO: public_metrics
	Width    int
	Alt_text string
	Variants []MediaVariants
}

type MediaVariants struct {
	Bit_rate     int
	Content_type string
	Url          string
}
