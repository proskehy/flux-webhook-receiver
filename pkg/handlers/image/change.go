package image

type ImageChange struct {
	Kind   string
	Source Source
}

type Source struct {
	Name struct {
		Domain string
		Image  string
	}
}
