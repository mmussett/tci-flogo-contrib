package json2xml

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
}

type Input struct {
	Ordered       bool   `md:ordered`
	ContentAsJson string `md:"contentAsXml"`
}

// ToMap conversion
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ordered":       i.Ordered,
		"contentAsJson": i.ContentAsJson,
	}
}

// FromMap conversion
func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	i.ContentAsJson, err = coerce.ToString(values["contentAsJson"])
	if err != nil {
		return err
	}

	i.Ordered, err = coerce.ToBool(values["ordered"])
	if err != nil {
		return err
	}

	return nil
}

// Output struct for activity output
type Output struct {
	ContentAsXml string `md:"contentAsXml"`
}

// ToMap conversion
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"contentAsXml": o.ContentAsXml,
	}
}

// FromMap conversion
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.ContentAsXml, err = coerce.ToString(values["contentAsXml"])
	if err != nil {
		return err
	}

	return nil
}
