package btypes

type EntityMeta struct {
	Href         string `json:"href"`
	MetadataHref string `json:"metadataHref"`
	Type         string `json:"type"`
	MediaType    string `json:"mediaType"`
}
