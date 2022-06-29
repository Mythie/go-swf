package swfparser

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/mythie/go-swf/pkg/swfheader"
	"github.com/mythie/go-swf/pkg/swfreader"
	"github.com/mythie/go-swf/pkg/swftag"
)

type ParsedSWF struct {
	reader swfreader.Reader
	Header *swfheader.Header
	Tags   *[]swftag.SWFTag
	parsed bool
}

func FromReader(reader io.Reader) (*ParsedSWF, error) {
	r, err := swfreader.NewReader(reader)

	if err != nil {
		return nil, err
	}

	swf := &ParsedSWF{
		reader: r,
	}

	err = swf.parse()

	if err != nil {
		return nil, err
	}

	return swf, nil
}

func (s *ParsedSWF) parse() error {
	s.Header, _ = swfheader.ParseHeader(s.reader)
	s.Tags, _ = swftag.ParseTags(s.reader)

	v, _ := json.MarshalIndent(s, "", "  ")

	fmt.Println(string(v))

	return nil
}
