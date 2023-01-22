package decode

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
}

type Input struct {
	Encoded      bool   `md:encoded`
	ContentAsXml string `md:"contentAsXml"`
}

// ToMap conversion
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"encoded":      i.Encoded,
		"contentAsXml": i.ContentAsXml,
	}
}

// FromMap conversion
func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	i.ContentAsXml, err = coerce.ToString(values["contentAsXml"])
	if err != nil {
		return err
	}

	i.Encoded, err = coerce.ToBool(values["encoded"])
	if err != nil {
		return err
	}

	return nil
}

// Output struct for activity output
type Output struct {
	ContentAsJson string `md:"contentAsJson"`
}

// ToMap conversion
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"contentAsJson": o.ContentAsJson,
	}
}

// FromMap conversion
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.ContentAsJson, err = coerce.ToString(values["contentAsJson"])
	if err != nil {
		return err
	}

	return nil
}
