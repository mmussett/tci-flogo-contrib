/*
 * Copyright © 2023. Mark Mussett.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */

package json2xml

import (
	"fmt"
	"github.com/mmussett/mxj"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityLog = log.ChildLogger(log.RootLogger(), "activity-xml2json")

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

	activityLog.Info("Executing json2xml activity")

	if input.ContentAsJson == "" {
		return false, activity.NewError("JSON content is empty", "JSON-DECODE-4000", nil)
	}

	var jsondata = input.ContentAsJson

	mv, err := mxj.NewMapJson([]byte(jsondata))
	mv.Json(true)

	if err != nil {
		return false, err
	}

	var xml []byte
	if input.Ordered {
		msv := mxj.MapSeq(mv)
		xml, err = msv.Xml()
		if err != nil {
			return false, err
		}

	} else {
		xml, err = mv.Xml()
		if err != nil {
			return false, err
		}

	}

	output := &Output{}
	output.ContentAsXml = string(xml)

	activityLog.Debug(output.ContentAsXml)

	err = context.SetOutputObject(output)
	if err != nil {
		return false, fmt.Errorf("error setting output for Activity [%s]: %s", context.Name(), err.Error())
	}
	return true, nil
}
