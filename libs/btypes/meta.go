package btypes

type Meta struct {
	Href         string `json:"href"`
	MetadataHref string `json:"metadataHref"`
	Type         string `json:"type"`
	MediaType    string `json:"mediaType"`
	Size         int    `json:"size"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}
