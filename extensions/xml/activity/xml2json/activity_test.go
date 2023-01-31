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
	"testing"
)

var xmlData = []byte(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:get="http://www.iata.org/IATA/EDIST/2017.2"><soapenv:Body><OrderCreateRQ xmlns="http://www.iata.org/IATA/EDIST/2017.2" Version="17.2" PrimaryLangID="EN" AltLangID="EN"><Document><Name>BA</Name></Document><Party><Sender><CorporateSender><ID>JB000000</ID></CorporateSender></Sender><Participants><Participant><TravelAgencyParticipant SequenceNumber="1"><Contacts><Contact><EmailContact><Address>agent.email@xyz.com</Address></EmailContact></Contact></Contacts><IATA_Number>00000000</IATA_Number><AgencyID>Test_Agency</AgencyID></TravelAgencyParticipant></Participant></Participants></Party><Query><Order><Offer OfferID="OF-44a450e6-75b3-4443-9c18-1024137b082f" Owner="BA" ResponseID="tx-08-201-0d5cafdd-ea25-48ac-9a3b-42480defcad4"><OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-1"><PassengerRefs>SH1 SH2 SH3</PassengerRefs></OfferItem><OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-2"><PassengerRefs>SH4 SH5</PassengerRefs></OfferItem><OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-3"><PassengerRefs>SH6 SH7</PassengerRefs></OfferItem><OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-4"><PassengerRefs>SH8 SH9</PassengerRefs></OfferItem></Offer></Order><Payments><Payment><Type>CC</Type><Method><PaymentCard><CardCode>MD</CardCode><CardNumber>7777884566565720</CardNumber><SeriesCode>123</SeriesCode><CardHolderName>MR MIKE TEST</CardHolderName><CardHolderBillingAddress><Street>Beeches Apartment</Street><Street>200 Lampton Road</Street><CityName>LON</CityName><PostalCode>TW345RT</PostalCode><CountryCode>GB</CountryCode></CardHolderBillingAddress><Surcharge><Amount Code="GBP">0.00</Amount></Surcharge><EffectiveExpireDate><Effective>1212</Effective><Expiration>0219</Expiration></EffectiveExpireDate></PaymentCard></Method><Amount Code="GBP">528.70</Amount><Payer><ContactInfoRefs>Payer</ContactInfoRefs></Payer></Payment></Payments><DataLists><PassengerList><Passenger PassengerID="SH1"><PTC>ADT</PTC><Individual><Birthdate>1982-12-15</Birthdate><Gender>Male</Gender><NameTitle>DR</NameTitle><GivenName>one</GivenName><Surname>TEST</Surname></Individual><ContactInfoRef>ContactInfo-SH1</ContactInfoRef></Passenger><Passenger PassengerID="SH2"><PTC>ADT</PTC><Individual><Birthdate>1983-08-05</Birthdate><Gender>Male</Gender><NameTitle>DR</NameTitle><GivenName>TWO</GivenName><Surname>TEST</Surname></Individual><InfantRef>SH9</InfantRef></Passenger><Passenger PassengerID="SH3"><PTC>ADT</PTC><Individual><Birthdate>1984-12-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>three</GivenName><Surname>TEST</Surname></Individual><InfantRef>SH8</InfantRef></Passenger><Passenger PassengerID="SH4"><PTC>ADT</PTC><Individual><Birthdate>2005-10-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>four</GivenName><Surname>TEST</Surname></Individual></Passenger><Passenger PassengerID="SH5"><PTC>ADT</PTC><Individual><Birthdate>2005-12-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>five</GivenName><Surname>TEST</Surname></Individual></Passenger><Passenger PassengerID="SH6"><PTC>CHD</PTC><Individual><Birthdate>2010-12-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>six</GivenName><Surname>TEST</Surname></Individual></Passenger><Passenger PassengerID="SH7"><PTC>CHD</PTC><Individual><Birthdate>2012-12-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>seven</GivenName><Surname>TEST</Surname></Individual></Passenger><Passenger PassengerID="SH8"><PTC>INF</PTC><Individual><Birthdate>2017-12-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>eight</GivenName><Surname>TEST</Surname></Individual></Passenger><Passenger PassengerID="SH9"><PTC>INF</PTC><Individual><Birthdate>2017-10-15</Birthdate><Gender>Male</Gender><NameTitle>MR</NameTitle><GivenName>nine</GivenName><Surname>TEST</Surname></Individual></Passenger></PassengerList><ContactList><ContactInformation ContactID="ContactInfo-SH1"><!--ContactType>Payment</ContactType--><ContactProvided><EmailAddress><EmailAddressValue>CBD.DBA@BA.COM</EmailAddressValue></EmailAddress></ContactProvided><ContactProvided><Phone><Label>mobile</Label><CountryDialingCode>11</CountryDialingCode><AreaCode>44</AreaCode><PhoneNumber>11122211</PhoneNumber></Phone></ContactProvided></ContactInformation><ContactInformation ContactID="Payer"><ContactType>Payment</ContactType><ContactProvided><EmailAddress><EmailAddressValue>third.party@xyz.com</EmailAddressValue></EmailAddress></ContactProvided><ContactProvided><Phone><Label>mobile</Label><CountryDialingCode>11</CountryDialingCode><AreaCode>44</AreaCode><PhoneNumber>888444444</PhoneNumber></Phone></ContactProvided><Individual><NameTitle>Mr</NameTitle><GivenName>MIKE</GivenName><Surname>TEST</Surname></Individual></ContactInformation></ContactList></DataLists></Query></OrderCreateRQ></soapenv:Body></soapenv:Envelope>`)

var encodedXmlData = []byte(`PHNvYXBlbnY6RW52ZWxvcGUgeG1sbnM6c29hcGVudj0iaHR0cDovL3NjaGVtYXMueG1sc29hcC5vcmcvc29hcC9lbnZlbG9wZS8iIHhtbG5zOmdldD0iaHR0cDovL3d3dy5pYXRhLm9yZy9JQVRBL0VESVNULzIwMTcuMiI+PHNvYXBlbnY6Qm9keT48T3JkZXJDcmVhdGVSUSB4bWxucz0iaHR0cDovL3d3dy5pYXRhLm9yZy9JQVRBL0VESVNULzIwMTcuMiIgVmVyc2lvbj0iMTcuMiIgUHJpbWFyeUxhbmdJRD0iRU4iIEFsdExhbmdJRD0iRU4iPjxEb2N1bWVudD48TmFtZT5CQTwvTmFtZT48L0RvY3VtZW50PjxQYXJ0eT48U2VuZGVyPjxDb3Jwb3JhdGVTZW5kZXI+PElEPkpCMDAwMDAwPC9JRD48L0NvcnBvcmF0ZVNlbmRlcj48L1NlbmRlcj48UGFydGljaXBhbnRzPjxQYXJ0aWNpcGFudD48VHJhdmVsQWdlbmN5UGFydGljaXBhbnQgU2VxdWVuY2VOdW1iZXI9IjEiPjxDb250YWN0cz48Q29udGFjdD48RW1haWxDb250YWN0PjxBZGRyZXNzPmFnZW50LmVtYWlsQHh5ei5jb208L0FkZHJlc3M+PC9FbWFpbENvbnRhY3Q+PC9Db250YWN0PjwvQ29udGFjdHM+PElBVEFfTnVtYmVyPjAwMDAwMDAwPC9JQVRBX051bWJlcj48QWdlbmN5SUQ+VGVzdF9BZ2VuY3k8L0FnZW5jeUlEPjwvVHJhdmVsQWdlbmN5UGFydGljaXBhbnQ+PC9QYXJ0aWNpcGFudD48L1BhcnRpY2lwYW50cz48L1BhcnR5PjxRdWVyeT48T3JkZXI+PE9mZmVyIE9mZmVySUQ9Ik9GLTQ0YTQ1MGU2LTc1YjMtNDQ0My05YzE4LTEwMjQxMzdiMDgyZiIgT3duZXI9IkJBIiBSZXNwb25zZUlEPSJ0eC0wOC0yMDEtMGQ1Y2FmZGQtZWEyNS00OGFjLTlhM2ItNDI0ODBkZWZjYWQ0Ij48T2ZmZXJJdGVtIE9mZmVySXRlbUlEPSJPRi00NGE0NTBlNi03NWIzLTQ0NDMtOWMxOC0xMDI0MTM3YjA4MmYtT0ktMSI+PFBhc3NlbmdlclJlZnM+U0gxIFNIMiBTSDM8L1Bhc3NlbmdlclJlZnM+PC9PZmZlckl0ZW0+PE9mZmVySXRlbSBPZmZlckl0ZW1JRD0iT0YtNDRhNDUwZTYtNzViMy00NDQzLTljMTgtMTAyNDEzN2IwODJmLU9JLTIiPjxQYXNzZW5nZXJSZWZzPlNINCBTSDU8L1Bhc3NlbmdlclJlZnM+PC9PZmZlckl0ZW0+PE9mZmVySXRlbSBPZmZlckl0ZW1JRD0iT0YtNDRhNDUwZTYtNzViMy00NDQzLTljMTgtMTAyNDEzN2IwODJmLU9JLTMiPjxQYXNzZW5nZXJSZWZzPlNINiBTSDc8L1Bhc3NlbmdlclJlZnM+PC9PZmZlckl0ZW0+PE9mZmVySXRlbSBPZmZlckl0ZW1JRD0iT0YtNDRhNDUwZTYtNzViMy00NDQzLTljMTgtMTAyNDEzN2IwODJmLU9JLTQiPjxQYXNzZW5nZXJSZWZzPlNIOCBTSDk8L1Bhc3NlbmdlclJlZnM+PC9PZmZlckl0ZW0+PC9PZmZlcj48L09yZGVyPjxQYXltZW50cz48UGF5bWVudD48VHlwZT5DQzwvVHlwZT48TWV0aG9kPjxQYXltZW50Q2FyZD48Q2FyZENvZGU+TUQ8L0NhcmRDb2RlPjxDYXJkTnVtYmVyPjc3Nzc4ODQ1NjY1NjU3MjA8L0NhcmROdW1iZXI+PFNlcmllc0NvZGU+MTIzPC9TZXJpZXNDb2RlPjxDYXJkSG9sZGVyTmFtZT5NUiBNSUtFIFRFU1Q8L0NhcmRIb2xkZXJOYW1lPjxDYXJkSG9sZGVyQmlsbGluZ0FkZHJlc3M+PFN0cmVldD5CZWVjaGVzIEFwYXJ0bWVudDwvU3RyZWV0PjxTdHJlZXQ+MjAwIExhbXB0b24gUm9hZDwvU3RyZWV0PjxDaXR5TmFtZT5MT048L0NpdHlOYW1lPjxQb3N0YWxDb2RlPlRXMzQ1UlQ8L1Bvc3RhbENvZGU+PENvdW50cnlDb2RlPkdCPC9Db3VudHJ5Q29kZT48L0NhcmRIb2xkZXJCaWxsaW5nQWRkcmVzcz48U3VyY2hhcmdlPjxBbW91bnQgQ29kZT0iR0JQIj4wLjAwPC9BbW91bnQ+PC9TdXJjaGFyZ2U+PEVmZmVjdGl2ZUV4cGlyZURhdGU+PEVmZmVjdGl2ZT4xMjEyPC9FZmZlY3RpdmU+PEV4cGlyYXRpb24+MDIxOTwvRXhwaXJhdGlvbj48L0VmZmVjdGl2ZUV4cGlyZURhdGU+PC9QYXltZW50Q2FyZD48L01ldGhvZD48QW1vdW50IENvZGU9IkdCUCI+NTI4LjcwPC9BbW91bnQ+PFBheWVyPjxDb250YWN0SW5mb1JlZnM+UGF5ZXI8L0NvbnRhY3RJbmZvUmVmcz48L1BheWVyPjwvUGF5bWVudD48L1BheW1lbnRzPjxEYXRhTGlzdHM+PFBhc3Nlbmdlckxpc3Q+PFBhc3NlbmdlciBQYXNzZW5nZXJJRD0iU0gxIj48UFRDPkFEVDwvUFRDPjxJbmRpdmlkdWFsPjxCaXJ0aGRhdGU+MTk4Mi0xMi0xNTwvQmlydGhkYXRlPjxHZW5kZXI+TWFsZTwvR2VuZGVyPjxOYW1lVGl0bGU+RFI8L05hbWVUaXRsZT48R2l2ZW5OYW1lPm9uZTwvR2l2ZW5OYW1lPjxTdXJuYW1lPlRFU1Q8L1N1cm5hbWU+PC9JbmRpdmlkdWFsPjxDb250YWN0SW5mb1JlZj5Db250YWN0SW5mby1TSDE8L0NvbnRhY3RJbmZvUmVmPjwvUGFzc2VuZ2VyPjxQYXNzZW5nZXIgUGFzc2VuZ2VySUQ9IlNIMiI+PFBUQz5BRFQ8L1BUQz48SW5kaXZpZHVhbD48QmlydGhkYXRlPjE5ODMtMDgtMDU8L0JpcnRoZGF0ZT48R2VuZGVyPk1hbGU8L0dlbmRlcj48TmFtZVRpdGxlPkRSPC9OYW1lVGl0bGU+PEdpdmVuTmFtZT5UV088L0dpdmVuTmFtZT48U3VybmFtZT5URVNUPC9TdXJuYW1lPjwvSW5kaXZpZHVhbD48SW5mYW50UmVmPlNIOTwvSW5mYW50UmVmPjwvUGFzc2VuZ2VyPjxQYXNzZW5nZXIgUGFzc2VuZ2VySUQ9IlNIMyI+PFBUQz5BRFQ8L1BUQz48SW5kaXZpZHVhbD48QmlydGhkYXRlPjE5ODQtMTItMTU8L0JpcnRoZGF0ZT48R2VuZGVyPk1hbGU8L0dlbmRlcj48TmFtZVRpdGxlPk1SPC9OYW1lVGl0bGU+PEdpdmVuTmFtZT50aHJlZTwvR2l2ZW5OYW1lPjxTdXJuYW1lPlRFU1Q8L1N1cm5hbWU+PC9JbmRpdmlkdWFsPjxJbmZhbnRSZWY+U0g4PC9JbmZhbnRSZWY+PC9QYXNzZW5nZXI+PFBhc3NlbmdlciBQYXNzZW5nZXJJRD0iU0g0Ij48UFRDPkFEVDwvUFRDPjxJbmRpdmlkdWFsPjxCaXJ0aGRhdGU+MjAwNS0xMC0xNTwvQmlydGhkYXRlPjxHZW5kZXI+TWFsZTwvR2VuZGVyPjxOYW1lVGl0bGU+TVI8L05hbWVUaXRsZT48R2l2ZW5OYW1lPmZvdXI8L0dpdmVuTmFtZT48U3VybmFtZT5URVNUPC9TdXJuYW1lPjwvSW5kaXZpZHVhbD48L1Bhc3Nlbmdlcj48UGFzc2VuZ2VyIFBhc3NlbmdlcklEPSJTSDUiPjxQVEM+QURUPC9QVEM+PEluZGl2aWR1YWw+PEJpcnRoZGF0ZT4yMDA1LTEyLTE1PC9CaXJ0aGRhdGU+PEdlbmRlcj5NYWxlPC9HZW5kZXI+PE5hbWVUaXRsZT5NUjwvTmFtZVRpdGxlPjxHaXZlbk5hbWU+Zml2ZTwvR2l2ZW5OYW1lPjxTdXJuYW1lPlRFU1Q8L1N1cm5hbWU+PC9JbmRpdmlkdWFsPjwvUGFzc2VuZ2VyPjxQYXNzZW5nZXIgUGFzc2VuZ2VySUQ9IlNINiI+PFBUQz5DSEQ8L1BUQz48SW5kaXZpZHVhbD48QmlydGhkYXRlPjIwMTAtMTItMTU8L0JpcnRoZGF0ZT48R2VuZGVyPk1hbGU8L0dlbmRlcj48TmFtZVRpdGxlPk1SPC9OYW1lVGl0bGU+PEdpdmVuTmFtZT5zaXg8L0dpdmVuTmFtZT48U3VybmFtZT5URVNUPC9TdXJuYW1lPjwvSW5kaXZpZHVhbD48L1Bhc3Nlbmdlcj48UGFzc2VuZ2VyIFBhc3NlbmdlcklEPSJTSDciPjxQVEM+Q0hEPC9QVEM+PEluZGl2aWR1YWw+PEJpcnRoZGF0ZT4yMDEyLTEyLTE1PC9CaXJ0aGRhdGU+PEdlbmRlcj5NYWxlPC9HZW5kZXI+PE5hbWVUaXRsZT5NUjwvTmFtZVRpdGxlPjxHaXZlbk5hbWU+c2V2ZW48L0dpdmVuTmFtZT48U3VybmFtZT5URVNUPC9TdXJuYW1lPjwvSW5kaXZpZHVhbD48L1Bhc3Nlbmdlcj48UGFzc2VuZ2VyIFBhc3NlbmdlcklEPSJTSDgiPjxQVEM+SU5GPC9QVEM+PEluZGl2aWR1YWw+PEJpcnRoZGF0ZT4yMDE3LTEyLTE1PC9CaXJ0aGRhdGU+PEdlbmRlcj5NYWxlPC9HZW5kZXI+PE5hbWVUaXRsZT5NUjwvTmFtZVRpdGxlPjxHaXZlbk5hbWU+ZWlnaHQ8L0dpdmVuTmFtZT48U3VybmFtZT5URVNUPC9TdXJuYW1lPjwvSW5kaXZpZHVhbD48L1Bhc3Nlbmdlcj48UGFzc2VuZ2VyIFBhc3NlbmdlcklEPSJTSDkiPjxQVEM+SU5GPC9QVEM+PEluZGl2aWR1YWw+PEJpcnRoZGF0ZT4yMDE3LTEwLTE1PC9CaXJ0aGRhdGU+PEdlbmRlcj5NYWxlPC9HZW5kZXI+PE5hbWVUaXRsZT5NUjwvTmFtZVRpdGxlPjxHaXZlbk5hbWU+bmluZTwvR2l2ZW5OYW1lPjxTdXJuYW1lPlRFU1Q8L1N1cm5hbWU+PC9JbmRpdmlkdWFsPjwvUGFzc2VuZ2VyPjwvUGFzc2VuZ2VyTGlzdD48Q29udGFjdExpc3Q+PENvbnRhY3RJbmZvcm1hdGlvbiBDb250YWN0SUQ9IkNvbnRhY3RJbmZvLVNIMSI+PCEtLUNvbnRhY3RUeXBlPlBheW1lbnQ8L0NvbnRhY3RUeXBlLS0+PENvbnRhY3RQcm92aWRlZD48RW1haWxBZGRyZXNzPjxFbWFpbEFkZHJlc3NWYWx1ZT5DQkQuREJBQEJBLkNPTTwvRW1haWxBZGRyZXNzVmFsdWU+PC9FbWFpbEFkZHJlc3M+PC9Db250YWN0UHJvdmlkZWQ+PENvbnRhY3RQcm92aWRlZD48UGhvbmU+PExhYmVsPm1vYmlsZTwvTGFiZWw+PENvdW50cnlEaWFsaW5nQ29kZT4xMTwvQ291bnRyeURpYWxpbmdDb2RlPjxBcmVhQ29kZT40NDwvQXJlYUNvZGU+PFBob25lTnVtYmVyPjExMTIyMjExPC9QaG9uZU51bWJlcj48L1Bob25lPjwvQ29udGFjdFByb3ZpZGVkPjwvQ29udGFjdEluZm9ybWF0aW9uPjxDb250YWN0SW5mb3JtYXRpb24gQ29udGFjdElEPSJQYXllciI+PENvbnRhY3RUeXBlPlBheW1lbnQ8L0NvbnRhY3RUeXBlPjxDb250YWN0UHJvdmlkZWQ+PEVtYWlsQWRkcmVzcz48RW1haWxBZGRyZXNzVmFsdWU+dGhpcmQucGFydHlAeHl6LmNvbTwvRW1haWxBZGRyZXNzVmFsdWU+PC9FbWFpbEFkZHJlc3M+PC9Db250YWN0UHJvdmlkZWQ+PENvbnRhY3RQcm92aWRlZD48UGhvbmU+PExhYmVsPm1vYmlsZTwvTGFiZWw+PENvdW50cnlEaWFsaW5nQ29kZT4xMTwvQ291bnRyeURpYWxpbmdDb2RlPjxBcmVhQ29kZT40NDwvQXJlYUNvZGU+PFBob25lTnVtYmVyPjg4ODQ0NDQ0NDwvUGhvbmVOdW1iZXI+PC9QaG9uZT48L0NvbnRhY3RQcm92aWRlZD48SW5kaXZpZHVhbD48TmFtZVRpdGxlPk1yPC9OYW1lVGl0bGU+PEdpdmVuTmFtZT5NSUtFPC9HaXZlbk5hbWU+PFN1cm5hbWU+VEVTVDwvU3VybmFtZT48L0luZGl2aWR1YWw+PC9Db250YWN0SW5mb3JtYXRpb24+PC9Db250YWN0TGlzdD48L0RhdGFMaXN0cz48L1F1ZXJ5PjwvT3JkZXJDcmVhdGVSUT48L3NvYXBlbnY6Qm9keT48L3NvYXBlbnY6RW52ZWxvcGU+`)

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

func TestEvalUnorderedEncoded(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	tc.SetInput("contentAsXml", encodedXmlData)
	tc.SetInput("encoded", true)
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
	xml, err := mv.Xml()

	fmt.Println("Inverse  : ", string(xml))

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
