/*
 * Copyright © 2023. Mark Mussett.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */

package xml2json

import (
	"fmt"
	"github.com/clbanning/mxj"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var xmlData = []byte(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:get="http://www.iata.org/IATA/EDIST/2017.2"><soapenv:Body><OrderCreateRQ xmlns="http://www.iata.org/IATA/EDIST/2017.2" Version="17.2" PrimaryLangID="EN" AltLangID="EN"><Document><Name>BA</Name></Document><Party><Sender><CorporateSender><ID>JB000000</ID></CorporateSender></Sender><Participants><Participant><TravelAgencyParticipant SequenceNumber="1"><Contacts><Contact><EmailContact><Address>agent.email@xyz.com</Address></EmailContact></Contact></Contacts><IATA_Number>00000000</IATA_Number><AgencyID>Test_Agency</AgencyID></TravelAgencyParticipant></Participant></Participants></Party><Query><Order><Offer OfferID="OF-44a450e6-75b3-4443-9c18-1024137b082f" Owner="BA" ResponseID="tx-08-201-0d5cafdd-ea25-48ac-9a3b-42480defcad4"><OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-1"><PassengerRefs>SH1 SH2 SH3</PassengerRefs></OfferItem><OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-2"><PassengerRefs>SH4 SH5</PassengerRefs></OfferItem><OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-3"><PassengerRefs>SH6 SH7</PassengerRefs></OfferItem><OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-4"><PassengerRefs>SH8 SH9</PassengerRefs></OfferItem></Offer></Order><Payments><Payment><Type>CC</Type><Method><PaymentCard><CardCode>MD</CardCode><CardNumber>7777884566565720</CardNumber><SeriesCode>123</SeriesCode><CardHolderName>MR MIKE TEST</CardHolderName><CardHolderBillingAddress><Street>Beeches Apartment</Street><Street>200 Lampton Road</Street><CityName>LON</CityName><PostalCode>TW345RT</PostalCode><CountryCode>GB</CountryCode></CardHolderBillingAddress><Surcharge><Amount Code="GBP">0.00</Amount></Surcharge><EffectiveExpireDate><Effective>1212</Effective><Expiration>0219</Expiration></EffectiveExpireDate></PaymentCard></Method><Amount Code="GBP">528.70</Amount><Payer><ContactInfoRefs>Payer</ContactInfoRefs></Payer></Payment></Payments><DataLists><PassengerList><Passenger PassengerID="SH1"><PTC>ADT</PTC><Individual><Birthdate>1982-12-15</Birthdate><Gender>Male</Gender><NameTitle>DR</NameTitle><GivenName>one</GivenName><Surname>TEST</Surname></Individual><ContactInfoRef>ContactInfo-SH1</ContactInfoRef></Passenger><Passenger PassengerID="SH2"><PTC>ADT</PTC><Individual><Birthdate>1983-08-05</Birthdate><Gender>Male</Gender><NameTitle>DR</NameTitle><GivenName>TWO</GivenName><Surname>TEST</Surname></Individual><InfantRef>SH9</InfantRef></Passenger><Passenger PassengerID="SH3"><PTC>ADT</PTC><Individual><Birthdate>1984-12-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>three</GivenName><Surname>TEST</Surname></Individual><InfantRef>SH8</InfantRef></Passenger><Passenger PassengerID="SH4"><PTC>ADT</PTC><Individual><Birthdate>2005-10-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>four</GivenName><Surname>TEST</Surname></Individual></Passenger><Passenger PassengerID="SH5"><PTC>ADT</PTC><Individual><Birthdate>2005-12-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>five</GivenName><Surname>TEST</Surname></Individual></Passenger><Passenger PassengerID="SH6"><PTC>CHD</PTC><Individual><Birthdate>2010-12-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>six</GivenName><Surname>TEST</Surname></Individual></Passenger><Passenger PassengerID="SH7"><PTC>CHD</PTC><Individual><Birthdate>2012-12-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>seven</GivenName><Surname>TEST</Surname></Individual></Passenger><Passenger PassengerID="SH8"><PTC>INF</PTC><Individual><Birthdate>2017-12-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>eight</GivenName><Surname>TEST</Surname></Individual></Passenger><Passenger PassengerID="SH9"><PTC>INF</PTC><Individual><Birthdate>2017-10-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>nine</GivenName><Surname>TEST</Surname></Individual></Passenger></PassengerList><ContactList><ContactInformation ContactID="ContactInfo-SH1"><!--ContactType>Payment</ContactType--><ContactProvided><EmailAddress><EmailAddressValue>CBD.DBA@BA.COM</EmailAddressValue></EmailAddress></ContactProvided><ContactProvided><Phone><Label>mobile</Label><CountryDialingCode>11</CountryDialingCode><AreaCode>44</AreaCode><PhoneNumber>11122211</PhoneNumber></Phone></ContactProvided></ContactInformation><ContactInformation ContactID="Payer"><ContactType>Payment</ContactType><ContactProvided><EmailAddress><EmailAddressValue>third.party@xyz.com</EmailAddressValue></EmailAddress></ContactProvided><ContactProvided><Phone><Label>mobile</Label><CountryDialingCode>11</CountryDialingCode><AreaCode>44</AreaCode><PhoneNumber>888444444</PhoneNumber></Phone></ContactProvided><Individual><NameTitle>Mr</NameTitle><GivenName>MIKE</GivenName><Surname>TEST</Surname></Individual></ContactInformation></ContactList></DataLists></Query></OrderCreateRQ></soapenv:Body></soapenv:Envelope>`)

var encodedXmlData = []byte(`PHNvYXBlbnY6RW52ZWxvcGUgeG1sbnM6c29hcGVudj0iIiB4bWxuczpnZXQ9IiI+PHNvYXBlbnY6Qm9keT48T3JkZXJDcmVhdGVSUSB4bWxucz0iIiBWZXJzaW9uPSIiIFByaW1hcnlMYW5nSUQ9IiIgQWx0TGFuZ0lEPSIiPjwvT3JkZXJDcmVhdGVSUT48L3NvYXBlbnY6Qm9keT48L3NvYXBlbnY6RW52ZWxvcGU+`)

var expectedOrdered = []byte(`{"soapenv:Envelope":{"#attr":{"xmlns:get":{"#seq":1,"#text":""},"xmlns:soapenv":{"#seq":0,"#text":""}},"soapenv:Body":{"#seq":0,"OrderCreateRQ":{"#attr":{"AltLangID":{"#seq":3,"#text":""},"PrimaryLangID":{"#seq":2,"#text":""},"Version":{"#seq":1,"#text":""},"xmlns":{"#seq":0,"#text":""}},"#seq":0}}}}`)
var expectedUnordered = []byte(``)

func TestCreate(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEvalOrdered(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	tc.SetInput("contentAsXml", xmlData)
	tc.SetInput("encoded", false)
	tc.SetInput("ordered", true)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	var output = fmt.Sprint(tc.GetOutput("contentAsJson"))
	fmt.Println("Input    : ", string(xmlData))
	//fmt.Println("Expected : ", string(expectedOrdered))
	fmt.Println("Output   : ", output)

	// Inverse

	mv, err := mxj.NewMapJson([]byte(output))
	if err != nil {
		assert.Error(t, err)
	}

	mv.Json(true)
	xml, err := mv.XmlSeq()

	fmt.Println("Inverse  : ", string(xml))

	//require.JSONEq(t, string(expectedOrdered), fmt.Sprint(tc.GetOutput("contentAsJson")))

}

func TestEvalUnordered(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	tc.SetInput("contentAsXml", xmlData)
	tc.SetInput("encoded", false)
	tc.SetInput("ordered", false)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	var output = fmt.Sprint(tc.GetOutput("contentAsJson"))
	fmt.Println("Input    : ", string(xmlData))
	//fmt.Println("Expected : ", string(expectedOrdered))
	fmt.Println("Output   : ", output)

	// Inverse

	mv, err := mxj.NewMapJson([]byte(output))
	if err != nil {
		assert.Error(t, err)
	}

	mv.Json(true)
	xml, err := mv.Xml()

	fmt.Println("Inverse  : ", string(xml))

}

func TestEvalOrderedEncoded(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	tc.SetInput("contentAsXml", encodedXmlData)
	tc.SetInput("encoded", true)
	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	require.JSONEq(t, string(expectedUnordered), fmt.Sprint(tc.GetOutput("contentAsJson")))

}

//func TestAnyXML(t *testing.T) {
//
//	jsonObj, err := xml2json.XmlToJson(xmlData, true)
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	fmt.Println(string(jsonObj[:len(jsonObj)]))
//
//	xmlObj, err := json2xml.JsonToXml(jsonObj)
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	fmt.Println(string(xmlObj[:len(xmlObj)]))
//
//}
