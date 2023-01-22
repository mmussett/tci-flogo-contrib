/*
 * Copyright © 2023. Mark Mussett.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */

package decode

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

/*
 * Copyright © 2023. Mark Mussett.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */

func init() {
	_ = activity.Register(&DecodeActivity{}, New)
}

var activityLog = log.ChildLogger(log.RootLogger(), "aws-activity-sqssendmessage")

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

type DecodeActivity struct {
	settings *Settings
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)

	if err != nil {
		return nil, err
	}

	act := &DecodeActivity{settings: s}
	return act, nil
}

func (a *DecodeActivity) Metadata() *activity.Metadata {
	return activityMd
}

func (a *DecodeActivity) Eval(context activity.Context) (done bool, err error) {

	input := &Input{}

	err = context.GetInputObject(input)
	if err != nil {
		return false, err
	}

	activityLog.Info("Executing decode activity")

	if input.contentAsXml == "" {
		return false, activity.NewError("XML content is empty", "XML-DECODE-4000", nil)
	}

	var data interface{}

	if a.settings.Encoded {
		contentAsXml, err := base64.StdEncoding.DecodeString(input.contentAsXml)
		if err != nil {
			logger.Debugf("Error decoding string: %s", err.Error())
			return false, activity.NewError("Error decoding base64 encoded string", "XML-DECODE-4002", nil)
		}
		activityLog.Debug(contentAsXml)
		xml.Unmarshal([]byte(contentAsXml), &data)

	} else {
		activityLog.Debug(input.contentAsXml)
		xml.Unmarshal([]byte(input.contentAsXml), &data)
	}
	
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return false, activity.NewError("Error marshalling JSON output", "XML-DECODE-4001", nil)
	}

	activityLog.Debug(string(jsonData))

	//Set  ID in the output
	output := &Output{}
	output.contentAsJson = string(jsonData)

	err = context.SetOutputObject(output)
	if err != nil {
		return false, fmt.Errorf("error setting output for Activity [%s]: %s", context.Name(), err.Error())
	}
	return true, nil
}
