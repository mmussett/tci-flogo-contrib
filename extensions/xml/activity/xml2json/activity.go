/*
 * Copyright © 2023. Mark Mussett.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */

package xml2json

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/logger"
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

func mapseqToJson(msv mxj.MapSeq, safeEncoding ...bool) ([]byte, error) {
	var s bool
	if len(safeEncoding) == 1 {
		s = safeEncoding[0]
	}

	b, err := json.Marshal(msv)

	if !s {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func (a *Activity) Eval(context activity.Context) (done bool, err error) {

	input := &Input{}

	err = context.GetInputObject(input)
	if err != nil {
		return false, err
	}

	activityLog.Info("Executing xml2json activity")

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

	// Preserve ordering of XML elements
	var json []byte

	if input.Ordered {
		msv, err := mxj.NewMapXmlSeq([]byte(xmldata), true)
		if err != nil {
			return false, err
		}

		json, err = mapseqToJson(msv, false)

		if err != nil {
			return false, err
		}

	} else {
		mv, err := mxj.NewMapXml([]byte(xmldata), false)
		if err != nil {
			return false, err
		}
		json, err = mv.Json(false)
		if err != nil {
			return false, err
		}
	}

	output := &Output{}
	output.ContentAsJson = string(json)

	activityLog.Debug(output.ContentAsJson)

	err = context.SetOutputObject(output)
	if err != nil {
		return false, fmt.Errorf("error setting output for Activity [%s]: %s", context.Name(), err.Error())
	}
	return true, nil
}
