package text

import (
	"github.com/golang/freetype/truetype"
	"io/ioutil"
	"log"
)

func Load(filepath string) (*truetype.Font, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(b)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return font, nil
}
