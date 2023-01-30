/*
 * Copyright © 2023. Mark Mussett.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */

package x2j

import (
	"encoding/base64"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"strings"

	xj "github.com/basgys/goxml2json"
)

/*
 * Copyright © 2023. Mark Mussett.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityLog = log.ChildLogger(log.RootLogger(), "aws-activity-sqssendmessage")

var activityMd = activity.ToMetadata(&Input{}, &Output{})

type Activity struct {
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)

	if err != nil {
		return nil, err
	}

	act := &Activity{}
	return act, nil
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func (a *Activity) Eval(context activity.Context) (done bool, err error) {

	input := &Input{}

	err = context.GetInputObject(input)
	if err != nil {
		return false, err
	}

	activityLog.Info("Executing decode activity")

	if input.ContentAsXml == "" {
		return false, activity.NewError("XML content is empty", "XML-DECODE-4000", nil)
	}

	var xmldata = ""
	if input.Encoded {
		data, err := base64.StdEncoding.DecodeString(input.ContentAsXml)
		if err != nil {
			logger.Debugf("Error decoding string: %s", err.Error())
			return false, activity.NewError("Error decoding base64 encoded string", "XML-DECODE-4002", nil)
		}
		xmldata = string(data)
	} else {
		xmldata = input.ContentAsXml
	}

	json, err := xj.Convert(strings.NewReader(xmldata))
	if err != nil {
		panic("That's embarrassing...")
	}

	activityLog.Debug(json.String())

	output := &Output{}
	output.ContentAsJson = json.String()

	err = context.SetOutputObject(output)
	if err != nil {
		return false, fmt.Errorf("error setting output for Activity [%s]: %s", context.Name(), err.Error())
	}
	return true, nil
}
