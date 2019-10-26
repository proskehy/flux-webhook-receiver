package handlers

type GitChange struct {
	Kind   string
	Source Source
}

type Source struct {
	URL    string
	Branch string
}
