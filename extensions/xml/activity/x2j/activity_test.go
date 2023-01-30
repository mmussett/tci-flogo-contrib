/*
 * Copyright © 2023. Mark Mussett.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */

package x2j

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

var xmlData = `<?xml version="1.0" encoding="UTF-8"?>
	<note>
	  <to>Tove</to>
	  <from>Jani</from>
	  <heading>Reminder</heading>
	  <body>Don't forget me this weekend!</body>
	</note>`

var encodedXmlData = `PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KCTxub3RlPgoJICA8dG8+VG92ZTwvdG8+CgkgIDxmcm9tPkphbmk8L2Zyb20+CgkgIDxoZWFkaW5nPlJlbWluZGVyPC9oZWFkaW5nPgoJICA8Ym9keT5Eb24ndCBmb3JnZXQgbWUgdGhpcyB3ZWVrZW5kITwvYm9keT4KCTwvbm90ZT4=`

var expected = `{"note": {"to": "Tove", "from": "Jani", "heading": "Reminder", "body": "Don't forget me this weekend!"}}`

func TestCreate(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval1(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	tc.SetInput("contentAsXml", xmlData)
	tc.SetInput("encoded", false)
	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	require.JSONEq(t, expected, fmt.Sprint(tc.GetOutput("contentAsJson")))

}

func TestEval2(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	tc.SetInput("contentAsXml", encodedXmlData)
	tc.SetInput("encoded", true)
	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	require.JSONEq(t, expected, fmt.Sprint(tc.GetOutput("contentAsJson")))

}
