package xml

import (
	"bytes"
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/stillwondering/minar"
)

// defaultIndent is used as the indentation for our XML documents.
const defaultIdent = "  "

// Encode returns the XML representation of the given minutes object.
func Encode(m minar.Minutes) ([]byte, error) {
	xmlMin := toXmlMinutes(m)

	buf := &bytes.Buffer{}
	buf.Write([]byte(xml.Header))

	encoded, err := xml.MarshalIndent(&xmlMin, "", defaultIdent)
	if err != nil {
		return nil, errors.Wrap(err, "marshal XML")
	}

	buf.Write(encoded)

	return buf.Bytes(), nil
}

// Decode returns the struct that's represented by the given XML input.
func Decode(b []byte) (minar.Minutes, error) {
	var decoded xmlMinutes

	err := xml.Unmarshal(b, &decoded)
	if err != nil {
		return minar.Minutes{}, errors.Wrap(err, "unmarshal XML")
	}

	return decoded.toMinutes(), nil
}

type xmlMinutes struct {
	XMLName      xml.Name   `xml:"minutes"`
	ID           string     `xml:"id,attr"`
	Title        string     `xml:"title"`
	Participants []string   `xml:"participant"`
	Topics       []xmlTopic `xml:"topic"`
}

type xmlTopic struct {
	Title   string `xml:"title"`
	Content string `xml:"content"`
}

func toXmlMinutes(m minar.Minutes) xmlMinutes {
	var topics []xmlTopic
	for _, topic := range m.Topics {
		topics = append(topics, xmlTopic{
			Title:   topic.Title,
			Content: topic.Content,
		})
	}

	return xmlMinutes{
		ID:           string(m.ID),
		Title:        m.Title,
		Participants: m.Participants,
		Topics:       topics,
	}
}

func (m xmlMinutes) toMinutes() minar.Minutes {
	var topics []minar.Topic
	for _, topic := range m.Topics {
		topics = append(topics, minar.Topic{
			Title:   topic.Title,
			Content: topic.Content,
		})
	}

	return minar.Minutes{
		ID:           minar.MinutesID(m.ID),
		Title:        m.Title,
		Participants: m.Participants,
		Topics:       topics,
	}
}
