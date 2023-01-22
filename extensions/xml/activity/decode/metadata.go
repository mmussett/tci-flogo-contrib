package decode

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	Encoded bool `md:"encoded"` // Ignore content-type header and treat as string
}

type Input struct {
	contentAsXml string `md:"ContentAsXml"`
}

// ToMap conversion
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ContentAsXml": i.contentAsXml,
	}
}

// FromMap conversion
func (i *Input) FromMap(values map[string]interface{}) error {
	var err error

	i.contentAsXml, err = coerce.ToString(values["ContentAsXml"])
	if err != nil {
		return err
	}

	return nil
}

// Output struct for activity output
type Output struct {
	contentAsJson string `md:"ContentAsJson"`
}

// ToMap conversion
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ContentAsJson": o.contentAsJson,
	}
}

// FromMap conversion
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.contentAsJson, err = coerce.ToString(values["ContentAsJson"])
	if err != nil {
		return err
	}

	return nil
}
