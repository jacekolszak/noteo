package yml

import (
	"time"

	"github.com/jacekolszak/noteo/notes"
	"github.com/jacekolszak/noteo/output"
	"gopkg.in/yaml.v2"
)

type Formatter struct{}

func (f Formatter) Header() string {
	return ""
}

func (f Formatter) Footer() string {
	return ""
}

func (f Formatter) Note(note notes.Note) string {
	text, err := note.Text()
	if err != nil {
		return err.Error()
	}
	created, err := note.Created()
	if err != nil {
		return err.Error()
	}
	tags, err := output.StringTags(note)
	if err != nil {
		return err.Error()
	}
	n := noteToMarshal{
		File:     note.Path(),
		Modified: note.Modified(),
		Created:  created,
		Text:     text,
		Tags:     tags,
	}
	bytes, err := yaml.Marshal(n)
	if err != nil {
		return "error marshalling note: " + err.Error()
	}
	return string(bytes) + "\n"
}

type noteToMarshal struct {
	File     string    `yaml:"file"`
	Modified time.Time `yaml:"modified"`
	Created  time.Time `yaml:"created"`
	Tags     []string  `yaml:"tags"`
	Text     string    `yaml:"text"`
}
