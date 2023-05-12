package models

type NoteFormat string

const (
	FORMAT_HTML       NoteFormat = "html"
	FORMAT_MD         NoteFormat = "md"
	FORMAT_PLAIN_TEXT NoteFormat = "plain_text"
)
