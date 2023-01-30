/*
 * Copyright © 2023. Mark Mussett.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */

package x2j

import (
	"fmt"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var xmlData = []byte(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:get="http://www.iata.org/IATA/EDIST/2017.2">
	<soapenv:Body>
		<OrderCreateRQ Version="17.2" PrimaryLangID="EN" AltLangID="EN" xmlns="http://www.iata.org/IATA/EDIST/2017.2">
			<Document>
				<Name>BA</Name>
			</Document>
			<Party>
				<Sender>
					<CorporateSender>
						<ID>JB000000</ID>
					</CorporateSender>
				</Sender>
				<Participants>
					<Participant>
						<TravelAgencyParticipant SequenceNumber="1">
							<Contacts>
								<Contact>
									<EmailContact>
										<Address>agent.email@xyz.com</Address>
									</EmailContact>
								</Contact>
							</Contacts>
							<IATA_Number>00000000</IATA_Number>
							<AgencyID>Test_Agency</AgencyID>
						</TravelAgencyParticipant>
					</Participant>
				</Participants>
			</Party>
			<Query>
				<Order>
					<Offer OfferID="OF-44a450e6-75b3-4443-9c18-1024137b082f" Owner="BA" ResponseID="tx-08-201-0d5cafdd-ea25-48ac-9a3b-42480defcad4">
						<OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-1">
							<PassengerRefs>SH1 SH2 SH3</PassengerRefs>
						</OfferItem>
						<OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-2">
							<PassengerRefs>SH4 SH5</PassengerRefs>
						</OfferItem>
						<OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-3">
							<PassengerRefs>SH6 SH7</PassengerRefs>
						</OfferItem>
						<OfferItem OfferItemID="OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-4">
							<PassengerRefs>SH8 SH9</PassengerRefs>
						</OfferItem>
					</Offer>
				</Order>
				<Payments>
					<Payment>
						<Type>CC</Type>
						<Method>
							<PaymentCard>
								<CardCode>MD</CardCode>
								<CardNumber>7777884566565720</CardNumber>
								<SeriesCode>123</SeriesCode>
								<CardHolderName>MR MIKE TEST</CardHolderName>
								<CardHolderBillingAddress>
									<Street>Beeches Apartment</Street>
									<Street>200 Lampton Road</Street>
									<CityName>LON</CityName>
									<PostalCode>TW345RT</PostalCode>
									<CountryCode>GB</CountryCode>
								</CardHolderBillingAddress>
								<Surcharge>
									<Amount Code="GBP">0.00</Amount>
								</Surcharge>
								<EffectiveExpireDate>
									<Effective>1212</Effective>
									<Expiration>0219</Expiration>
								</EffectiveExpireDate>
							</PaymentCard>
						</Method>
						<Amount Code="GBP">528.70</Amount>
						<Payer>
							<ContactInfoRefs>Payer</ContactInfoRefs>
						</Payer>
					</Payment>
				</Payments>
				<DataLists>
					<PassengerList>
						<Passenger PassengerID="SH1">
							<PTC>ADT</PTC>
							<Individual>
								<Birthdate>1982-12-15</Birthdate>
								<Gender>Male</Gender>
								<NameTitle>DR</NameTitle>
								<GivenName>one</GivenName>
								<Surname>TEST</Surname>
							</Individual>
							<ContactInfoRef>ContactInfo-SH1</ContactInfoRef>
						</Passenger>
						<Passenger PassengerID="SH2">
							<PTC>ADT</PTC>
							<Individual>
								<Birthdate>1983-08-05</Birthdate>
								<Gender>Male</Gender>
								<NameTitle>DR</NameTitle>
								<GivenName>TWO</GivenName>
								<Surname>TEST</Surname>
							</Individual>
							<InfantRef>SH9</InfantRef>
						</Passenger>
						<Passenger PassengerID="SH3">
							<PTC>ADT</PTC>
							<Individual>
								<Birthdate>1984-12-15</Birthdate>
								<Gender>Male</Gender>
								<NameTitle>MR</NameTitle>
								<GivenName>three</GivenName>
								<Surname>TEST</Surname>
							</Individual>
							<InfantRef>SH8</InfantRef>
						</Passenger>
						<Passenger PassengerID="SH4">
							<PTC>ADT</PTC>
							<Individual>
								<Birthdate>2005-10-15</Birthdate>
								<Gender>Male</Gender>
								<NameTitle>MR</NameTitle>
								<GivenName>four</GivenName>
								<Surname>TEST</Surname>
							</Individual>
						</Passenger>
						<Passenger PassengerID="SH5">
							<PTC>ADT</PTC>
							<Individual>
								<Birthdate>2005-12-15</Birthdate>
								<Gender>Male</Gender>
								<NameTitle>MR</NameTitle>
								<GivenName>five</GivenName>
								<Surname>TEST</Surname>
							</Individual>
						</Passenger>
						<Passenger PassengerID="SH6">
							<PTC>CHD</PTC>
							<Individual>
								<Birthdate>2010-12-15</Birthdate>
								<Gender>Male</Gender>
								<NameTitle>MR</NameTitle>
								<GivenName>six</GivenName>
								<Surname>TEST</Surname>
							</Individual>
						</Passenger>
						<Passenger PassengerID="SH7">
							<PTC>CHD</PTC>
							<Individual>
								<Birthdate>2012-12-15</Birthdate>
								<Gender>Male</Gender>
								<NameTitle>MR</NameTitle>
								<GivenName>seven</GivenName>
								<Surname>TEST</Surname>
							</Individual>
						</Passenger>
						<Passenger PassengerID="SH8">
							<PTC>INF</PTC>
							<Individual>
								<Birthdate>2017-12-15</Birthdate>
								<Gender>Male</Gender>
								<NameTitle>MR</NameTitle>
								<GivenName>eight</GivenName>
								<Surname>TEST</Surname>
							</Individual>
						</Passenger>
						<Passenger PassengerID="SH9">
							<PTC>INF</PTC>
							<Individual>
								<Birthdate>2017-10-15</Birthdate>
								<Gender>Male</Gender>
								<NameTitle>MR</NameTitle>
								<GivenName>nine</GivenName>
								<Surname>TEST</Surname>
							</Individual>
						</Passenger>
					</PassengerList>
					<ContactList>
						<ContactInformation ContactID="ContactInfo-SH1">
							<!--ContactType>Payment</ContactType-->
							<ContactProvided>
								<EmailAddress>
									<EmailAddressValue>CBD.DBA@BA.COM</EmailAddressValue>
								</EmailAddress>
							</ContactProvided>
							<ContactProvided>
								<Phone>
									<Label>mobile</Label>
									<CountryDialingCode>11</CountryDialingCode>
									<AreaCode>44</AreaCode>
									<PhoneNumber>11122211</PhoneNumber>
								</Phone>
							</ContactProvided>
						</ContactInformation>
						<ContactInformation ContactID="Payer">
							<ContactType>Payment</ContactType>
							<ContactProvided>
								<EmailAddress>
									<EmailAddressValue>third.party@xyz.com</EmailAddressValue>
								</EmailAddress>
							</ContactProvided>
							<ContactProvided>
								<Phone>
									<Label>mobile</Label>
									<CountryDialingCode>11</CountryDialingCode>
									<AreaCode>44</AreaCode>
									<PhoneNumber>888444444</PhoneNumber>
								</Phone>
							</ContactProvided>
							<Individual>
								<NameTitle>Mr</NameTitle>
								<GivenName>MIKE</GivenName>
								<Surname>TEST</Surname>
							</Individual>
						</ContactInformation>
					</ContactList>
				</DataLists>
			</Query>
		</OrderCreateRQ>
	</soapenv:Body>
</soapenv:Envelope>`)

var encodedXmlData = []byte(`PHNvYXBlbnY6RW52ZWxvcGUgeG1sbnM6c29hcGVudj0iaHR0cDovL3NjaGVtYXMueG1sc29hcC5vcmcvc29hcC9lbnZlbG9wZS8iIHhtbG5zOmdldD0iaHR0cDovL3d3dy5pYXRhLm9yZy9JQVRBL0VESVNULzIwMTcuMiI+Cgk8c29hcGVudjpCb2R5PgoJCTxPcmRlckNyZWF0ZVJRIFZlcnNpb249IjE3LjIiIFByaW1hcnlMYW5nSUQ9IkVOIiBBbHRMYW5nSUQ9IkVOIiB4bWxucz0iaHR0cDovL3d3dy5pYXRhLm9yZy9JQVRBL0VESVNULzIwMTcuMiI+CgkJCTxEb2N1bWVudD4KCQkJCTxOYW1lPkJBPC9OYW1lPgoJCQk8L0RvY3VtZW50PgoJCQk8UGFydHk+CgkJCQk8U2VuZGVyPgoJCQkJCTxDb3Jwb3JhdGVTZW5kZXI+CgkJCQkJCTxJRD5KQjAwMDAwMDwvSUQ+CgkJCQkJPC9Db3Jwb3JhdGVTZW5kZXI+CgkJCQk8L1NlbmRlcj4KCQkJCTxQYXJ0aWNpcGFudHM+CgkJCQkJPFBhcnRpY2lwYW50PgoJCQkJCQk8VHJhdmVsQWdlbmN5UGFydGljaXBhbnQgU2VxdWVuY2VOdW1iZXI9IjEiPgoJCQkJCQkJPENvbnRhY3RzPgoJCQkJCQkJCTxDb250YWN0PgoJCQkJCQkJCQk8RW1haWxDb250YWN0PgoJCQkJCQkJCQkJPEFkZHJlc3M+YWdlbnQuZW1haWxAeHl6LmNvbTwvQWRkcmVzcz4KCQkJCQkJCQkJPC9FbWFpbENvbnRhY3Q+CgkJCQkJCQkJPC9Db250YWN0PgoJCQkJCQkJPC9Db250YWN0cz4KCQkJCQkJCTxJQVRBX051bWJlcj4wMDAwMDAwMDwvSUFUQV9OdW1iZXI+CgkJCQkJCQk8QWdlbmN5SUQ+VGVzdF9BZ2VuY3k8L0FnZW5jeUlEPgoJCQkJCQk8L1RyYXZlbEFnZW5jeVBhcnRpY2lwYW50PgoJCQkJCTwvUGFydGljaXBhbnQ+CgkJCQk8L1BhcnRpY2lwYW50cz4KCQkJPC9QYXJ0eT4KCQkJPFF1ZXJ5PgoJCQkJPE9yZGVyPgoJCQkJCTxPZmZlciBPZmZlcklEPSJPRi00NGE0NTBlNi03NWIzLTQ0NDMtOWMxOC0xMDI0MTM3YjA4MmYiIE93bmVyPSJCQSIgUmVzcG9uc2VJRD0idHgtMDgtMjAxLTBkNWNhZmRkLWVhMjUtNDhhYy05YTNiLTQyNDgwZGVmY2FkNCI+CgkJCQkJCTxPZmZlckl0ZW0gT2ZmZXJJdGVtSUQ9Ik9GLTQ0YTQ1MGU2LTc1YjMtNDQ0My05YzE4LTEwMjQxMzdiMDgyZi1PSS0xIj4KCQkJCQkJCTxQYXNzZW5nZXJSZWZzPlNIMSBTSDIgU0gzPC9QYXNzZW5nZXJSZWZzPgoJCQkJCQk8L09mZmVySXRlbT4KCQkJCQkJPE9mZmVySXRlbSBPZmZlckl0ZW1JRD0iT0YtNDRhNDUwZTYtNzViMy00NDQzLTljMTgtMTAyNDEzN2IwODJmLU9JLTIiPgoJCQkJCQkJPFBhc3NlbmdlclJlZnM+U0g0IFNINTwvUGFzc2VuZ2VyUmVmcz4KCQkJCQkJPC9PZmZlckl0ZW0+CgkJCQkJCTxPZmZlckl0ZW0gT2ZmZXJJdGVtSUQ9Ik9GLTQ0YTQ1MGU2LTc1YjMtNDQ0My05YzE4LTEwMjQxMzdiMDgyZi1PSS0zIj4KCQkJCQkJCTxQYXNzZW5nZXJSZWZzPlNINiBTSDc8L1Bhc3NlbmdlclJlZnM+CgkJCQkJCTwvT2ZmZXJJdGVtPgoJCQkJCQk8T2ZmZXJJdGVtIE9mZmVySXRlbUlEPSJPRi00NGE0NTBlNi03NWIzLTQ0NDMtOWMxOC0xMDI0MTM3YjA4MmYtT0ktNCI+CgkJCQkJCQk8UGFzc2VuZ2VyUmVmcz5TSDggU0g5PC9QYXNzZW5nZXJSZWZzPgoJCQkJCQk8L09mZmVySXRlbT4KCQkJCQk8L09mZmVyPgoJCQkJPC9PcmRlcj4KCQkJCTxQYXltZW50cz4KCQkJCQk8UGF5bWVudD4KCQkJCQkJPFR5cGU+Q0M8L1R5cGU+CgkJCQkJCTxNZXRob2Q+CgkJCQkJCQk8UGF5bWVudENhcmQ+CgkJCQkJCQkJPENhcmRDb2RlPk1EPC9DYXJkQ29kZT4KCQkJCQkJCQk8Q2FyZE51bWJlcj43Nzc3ODg0NTY2NTY1NzIwPC9DYXJkTnVtYmVyPgoJCQkJCQkJCTxTZXJpZXNDb2RlPjEyMzwvU2VyaWVzQ29kZT4KCQkJCQkJCQk8Q2FyZEhvbGRlck5hbWU+TVIgTUlLRSBURVNUPC9DYXJkSG9sZGVyTmFtZT4KCQkJCQkJCQk8Q2FyZEhvbGRlckJpbGxpbmdBZGRyZXNzPgoJCQkJCQkJCQk8U3RyZWV0PkJlZWNoZXMgQXBhcnRtZW50PC9TdHJlZXQ+CgkJCQkJCQkJCTxTdHJlZXQ+MjAwIExhbXB0b24gUm9hZDwvU3RyZWV0PgoJCQkJCQkJCQk8Q2l0eU5hbWU+TE9OPC9DaXR5TmFtZT4KCQkJCQkJCQkJPFBvc3RhbENvZGU+VFczNDVSVDwvUG9zdGFsQ29kZT4KCQkJCQkJCQkJPENvdW50cnlDb2RlPkdCPC9Db3VudHJ5Q29kZT4KCQkJCQkJCQk8L0NhcmRIb2xkZXJCaWxsaW5nQWRkcmVzcz4KCQkJCQkJCQk8U3VyY2hhcmdlPgoJCQkJCQkJCQk8QW1vdW50IENvZGU9IkdCUCI+MC4wMDwvQW1vdW50PgoJCQkJCQkJCTwvU3VyY2hhcmdlPgoJCQkJCQkJCTxFZmZlY3RpdmVFeHBpcmVEYXRlPgoJCQkJCQkJCQk8RWZmZWN0aXZlPjEyMTI8L0VmZmVjdGl2ZT4KCQkJCQkJCQkJPEV4cGlyYXRpb24+MDIxOTwvRXhwaXJhdGlvbj4KCQkJCQkJCQk8L0VmZmVjdGl2ZUV4cGlyZURhdGU+CgkJCQkJCQk8L1BheW1lbnRDYXJkPgoJCQkJCQk8L01ldGhvZD4KCQkJCQkJPEFtb3VudCBDb2RlPSJHQlAiPjUyOC43MDwvQW1vdW50PgoJCQkJCQk8UGF5ZXI+CgkJCQkJCQk8Q29udGFjdEluZm9SZWZzPlBheWVyPC9Db250YWN0SW5mb1JlZnM+CgkJCQkJCTwvUGF5ZXI+CgkJCQkJPC9QYXltZW50PgoJCQkJPC9QYXltZW50cz4KCQkJCTxEYXRhTGlzdHM+CgkJCQkJPFBhc3Nlbmdlckxpc3Q+CgkJCQkJCTxQYXNzZW5nZXIgUGFzc2VuZ2VySUQ9IlNIMSI+CgkJCQkJCQk8UFRDPkFEVDwvUFRDPgoJCQkJCQkJPEluZGl2aWR1YWw+CgkJCQkJCQkJPEJpcnRoZGF0ZT4xOTgyLTEyLTE1PC9CaXJ0aGRhdGU+CgkJCQkJCQkJPEdlbmRlcj5NYWxlPC9HZW5kZXI+CgkJCQkJCQkJPE5hbWVUaXRsZT5EUjwvTmFtZVRpdGxlPgoJCQkJCQkJCTxHaXZlbk5hbWU+b25lPC9HaXZlbk5hbWU+CgkJCQkJCQkJPFN1cm5hbWU+VEVTVDwvU3VybmFtZT4KCQkJCQkJCTwvSW5kaXZpZHVhbD4KCQkJCQkJCTxDb250YWN0SW5mb1JlZj5Db250YWN0SW5mby1TSDE8L0NvbnRhY3RJbmZvUmVmPgoJCQkJCQk8L1Bhc3Nlbmdlcj4KCQkJCQkJPFBhc3NlbmdlciBQYXNzZW5nZXJJRD0iU0gyIj4KCQkJCQkJCTxQVEM+QURUPC9QVEM+CgkJCQkJCQk8SW5kaXZpZHVhbD4KCQkJCQkJCQk8QmlydGhkYXRlPjE5ODMtMDgtMDU8L0JpcnRoZGF0ZT4KCQkJCQkJCQk8R2VuZGVyPk1hbGU8L0dlbmRlcj4KCQkJCQkJCQk8TmFtZVRpdGxlPkRSPC9OYW1lVGl0bGU+CgkJCQkJCQkJPEdpdmVuTmFtZT5UV088L0dpdmVuTmFtZT4KCQkJCQkJCQk8U3VybmFtZT5URVNUPC9TdXJuYW1lPgoJCQkJCQkJPC9JbmRpdmlkdWFsPgoJCQkJCQkJPEluZmFudFJlZj5TSDk8L0luZmFudFJlZj4KCQkJCQkJPC9QYXNzZW5nZXI+CgkJCQkJCTxQYXNzZW5nZXIgUGFzc2VuZ2VySUQ9IlNIMyI+CgkJCQkJCQk8UFRDPkFEVDwvUFRDPgoJCQkJCQkJPEluZGl2aWR1YWw+CgkJCQkJCQkJPEJpcnRoZGF0ZT4xOTg0LTEyLTE1PC9CaXJ0aGRhdGU+CgkJCQkJCQkJPEdlbmRlcj5NYWxlPC9HZW5kZXI+CgkJCQkJCQkJPE5hbWVUaXRsZT5NUjwvTmFtZVRpdGxlPgoJCQkJCQkJCTxHaXZlbk5hbWU+dGhyZWU8L0dpdmVuTmFtZT4KCQkJCQkJCQk8U3VybmFtZT5URVNUPC9TdXJuYW1lPgoJCQkJCQkJPC9JbmRpdmlkdWFsPgoJCQkJCQkJPEluZmFudFJlZj5TSDg8L0luZmFudFJlZj4KCQkJCQkJPC9QYXNzZW5nZXI+CgkJCQkJCTxQYXNzZW5nZXIgUGFzc2VuZ2VySUQ9IlNINCI+CgkJCQkJCQk8UFRDPkFEVDwvUFRDPgoJCQkJCQkJPEluZGl2aWR1YWw+CgkJCQkJCQkJPEJpcnRoZGF0ZT4yMDA1LTEwLTE1PC9CaXJ0aGRhdGU+CgkJCQkJCQkJPEdlbmRlcj5NYWxlPC9HZW5kZXI+CgkJCQkJCQkJPE5hbWVUaXRsZT5NUjwvTmFtZVRpdGxlPgoJCQkJCQkJCTxHaXZlbk5hbWU+Zm91cjwvR2l2ZW5OYW1lPgoJCQkJCQkJCTxTdXJuYW1lPlRFU1Q8L1N1cm5hbWU+CgkJCQkJCQk8L0luZGl2aWR1YWw+CgkJCQkJCTwvUGFzc2VuZ2VyPgoJCQkJCQk8UGFzc2VuZ2VyIFBhc3NlbmdlcklEPSJTSDUiPgoJCQkJCQkJPFBUQz5BRFQ8L1BUQz4KCQkJCQkJCTxJbmRpdmlkdWFsPgoJCQkJCQkJCTxCaXJ0aGRhdGU+MjAwNS0xMi0xNTwvQmlydGhkYXRlPgoJCQkJCQkJCTxHZW5kZXI+TWFsZTwvR2VuZGVyPgoJCQkJCQkJCTxOYW1lVGl0bGU+TVI8L05hbWVUaXRsZT4KCQkJCQkJCQk8R2l2ZW5OYW1lPmZpdmU8L0dpdmVuTmFtZT4KCQkJCQkJCQk8U3VybmFtZT5URVNUPC9TdXJuYW1lPgoJCQkJCQkJPC9JbmRpdmlkdWFsPgoJCQkJCQk8L1Bhc3Nlbmdlcj4KCQkJCQkJPFBhc3NlbmdlciBQYXNzZW5nZXJJRD0iU0g2Ij4KCQkJCQkJCTxQVEM+Q0hEPC9QVEM+CgkJCQkJCQk8SW5kaXZpZHVhbD4KCQkJCQkJCQk8QmlydGhkYXRlPjIwMTAtMTItMTU8L0JpcnRoZGF0ZT4KCQkJCQkJCQk8R2VuZGVyPk1hbGU8L0dlbmRlcj4KCQkJCQkJCQk8TmFtZVRpdGxlPk1SPC9OYW1lVGl0bGU+CgkJCQkJCQkJPEdpdmVuTmFtZT5zaXg8L0dpdmVuTmFtZT4KCQkJCQkJCQk8U3VybmFtZT5URVNUPC9TdXJuYW1lPgoJCQkJCQkJPC9JbmRpdmlkdWFsPgoJCQkJCQk8L1Bhc3Nlbmdlcj4KCQkJCQkJPFBhc3NlbmdlciBQYXNzZW5nZXJJRD0iU0g3Ij4KCQkJCQkJCTxQVEM+Q0hEPC9QVEM+CgkJCQkJCQk8SW5kaXZpZHVhbD4KCQkJCQkJCQk8QmlydGhkYXRlPjIwMTItMTItMTU8L0JpcnRoZGF0ZT4KCQkJCQkJCQk8R2VuZGVyPk1hbGU8L0dlbmRlcj4KCQkJCQkJCQk8TmFtZVRpdGxlPk1SPC9OYW1lVGl0bGU+CgkJCQkJCQkJPEdpdmVuTmFtZT5zZXZlbjwvR2l2ZW5OYW1lPgoJCQkJCQkJCTxTdXJuYW1lPlRFU1Q8L1N1cm5hbWU+CgkJCQkJCQk8L0luZGl2aWR1YWw+CgkJCQkJCTwvUGFzc2VuZ2VyPgoJCQkJCQk8UGFzc2VuZ2VyIFBhc3NlbmdlcklEPSJTSDgiPgoJCQkJCQkJPFBUQz5JTkY8L1BUQz4KCQkJCQkJCTxJbmRpdmlkdWFsPgoJCQkJCQkJCTxCaXJ0aGRhdGU+MjAxNy0xMi0xNTwvQmlydGhkYXRlPgoJCQkJCQkJCTxHZW5kZXI+TWFsZTwvR2VuZGVyPgoJCQkJCQkJCTxOYW1lVGl0bGU+TVI8L05hbWVUaXRsZT4KCQkJCQkJCQk8R2l2ZW5OYW1lPmVpZ2h0PC9HaXZlbk5hbWU+CgkJCQkJCQkJPFN1cm5hbWU+VEVTVDwvU3VybmFtZT4KCQkJCQkJCTwvSW5kaXZpZHVhbD4KCQkJCQkJPC9QYXNzZW5nZXI+CgkJCQkJCTxQYXNzZW5nZXIgUGFzc2VuZ2VySUQ9IlNIOSI+CgkJCQkJCQk8UFRDPklORjwvUFRDPgoJCQkJCQkJPEluZGl2aWR1YWw+CgkJCQkJCQkJPEJpcnRoZGF0ZT4yMDE3LTEwLTE1PC9CaXJ0aGRhdGU+CgkJCQkJCQkJPEdlbmRlcj5NYWxlPC9HZW5kZXI+CgkJCQkJCQkJPE5hbWVUaXRsZT5NUjwvTmFtZVRpdGxlPgoJCQkJCQkJCTxHaXZlbk5hbWU+bmluZTwvR2l2ZW5OYW1lPgoJCQkJCQkJCTxTdXJuYW1lPlRFU1Q8L1N1cm5hbWU+CgkJCQkJCQk8L0luZGl2aWR1YWw+CgkJCQkJCTwvUGFzc2VuZ2VyPgoJCQkJCTwvUGFzc2VuZ2VyTGlzdD4KCQkJCQk8Q29udGFjdExpc3Q+CgkJCQkJCTxDb250YWN0SW5mb3JtYXRpb24gQ29udGFjdElEPSJDb250YWN0SW5mby1TSDEiPgoJCQkJCQkJPCEtLUNvbnRhY3RUeXBlPlBheW1lbnQ8L0NvbnRhY3RUeXBlLS0+CgkJCQkJCQk8Q29udGFjdFByb3ZpZGVkPgoJCQkJCQkJCTxFbWFpbEFkZHJlc3M+CgkJCQkJCQkJCTxFbWFpbEFkZHJlc3NWYWx1ZT5DQkQuREJBQEJBLkNPTTwvRW1haWxBZGRyZXNzVmFsdWU+CgkJCQkJCQkJPC9FbWFpbEFkZHJlc3M+CgkJCQkJCQk8L0NvbnRhY3RQcm92aWRlZD4KCQkJCQkJCTxDb250YWN0UHJvdmlkZWQ+CgkJCQkJCQkJPFBob25lPgoJCQkJCQkJCQk8TGFiZWw+bW9iaWxlPC9MYWJlbD4KCQkJCQkJCQkJPENvdW50cnlEaWFsaW5nQ29kZT4xMTwvQ291bnRyeURpYWxpbmdDb2RlPgoJCQkJCQkJCQk8QXJlYUNvZGU+NDQ8L0FyZWFDb2RlPgoJCQkJCQkJCQk8UGhvbmVOdW1iZXI+MTExMjIyMTE8L1Bob25lTnVtYmVyPgoJCQkJCQkJCTwvUGhvbmU+CgkJCQkJCQk8L0NvbnRhY3RQcm92aWRlZD4KCQkJCQkJPC9Db250YWN0SW5mb3JtYXRpb24+CgkJCQkJCTxDb250YWN0SW5mb3JtYXRpb24gQ29udGFjdElEPSJQYXllciI+CgkJCQkJCQk8Q29udGFjdFR5cGU+UGF5bWVudDwvQ29udGFjdFR5cGU+CgkJCQkJCQk8Q29udGFjdFByb3ZpZGVkPgoJCQkJCQkJCTxFbWFpbEFkZHJlc3M+CgkJCQkJCQkJCTxFbWFpbEFkZHJlc3NWYWx1ZT50aGlyZC5wYXJ0eUB4eXouY29tPC9FbWFpbEFkZHJlc3NWYWx1ZT4KCQkJCQkJCQk8L0VtYWlsQWRkcmVzcz4KCQkJCQkJCTwvQ29udGFjdFByb3ZpZGVkPgoJCQkJCQkJPENvbnRhY3RQcm92aWRlZD4KCQkJCQkJCQk8UGhvbmU+CgkJCQkJCQkJCTxMYWJlbD5tb2JpbGU8L0xhYmVsPgoJCQkJCQkJCQk8Q291bnRyeURpYWxpbmdDb2RlPjExPC9Db3VudHJ5RGlhbGluZ0NvZGU+CgkJCQkJCQkJCTxBcmVhQ29kZT40NDwvQXJlYUNvZGU+CgkJCQkJCQkJCTxQaG9uZU51bWJlcj44ODg0NDQ0NDQ8L1Bob25lTnVtYmVyPgoJCQkJCQkJCTwvUGhvbmU+CgkJCQkJCQk8L0NvbnRhY3RQcm92aWRlZD4KCQkJCQkJCTxJbmRpdmlkdWFsPgoJCQkJCQkJCTxOYW1lVGl0bGU+TXI8L05hbWVUaXRsZT4KCQkJCQkJCQk8R2l2ZW5OYW1lPk1JS0U8L0dpdmVuTmFtZT4KCQkJCQkJCQk8U3VybmFtZT5URVNUPC9TdXJuYW1lPgoJCQkJCQkJPC9JbmRpdmlkdWFsPgoJCQkJCQk8L0NvbnRhY3RJbmZvcm1hdGlvbj4KCQkJCQk8L0NvbnRhY3RMaXN0PgoJCQkJPC9EYXRhTGlzdHM+CgkJCTwvUXVlcnk+CgkJPC9PcmRlckNyZWF0ZVJRPgoJPC9zb2FwZW52OkJvZHk+Cjwvc29hcGVudjpFbnZlbG9wZT4=`)

var expected = []byte(`{"Envelope":{"-get":"http://www.iata.org/IATA/EDIST/2017.2","-soapenv":"http://schemas.xmlsoap.org/soap/envelope/","Body":{"OrderCreateRQ":{"-AltLangID":"EN","-PrimaryLangID":"EN","-Version":17.2,"-xmlns":"http://www.iata.org/IATA/EDIST/2017.2","Document":{"Name":"BA"},"Party":{"Participants":{"Participant":{"TravelAgencyParticipant":{"-SequenceNumber":1,"AgencyID":"Test_Agency","Contacts":{"Contact":{"EmailContact":{"Address":"agent.email@xyz.com"}}},"IATA_Number":0}}},"Sender":{"CorporateSender":{"ID":"JB000000"}}},"Query":{"DataLists":{"ContactList":{"ContactInformation":[{"-ContactID":"ContactInfo-SH1","ContactProvided":[{"EmailAddress":{"EmailAddressValue":"CBD.DBA@BA.COM"}},{"Phone":{"AreaCode":44,"CountryDialingCode":11,"Label":"mobile","PhoneNumber":11122211}}]},{"-ContactID":"Payer","ContactProvided":[{"EmailAddress":{"EmailAddressValue":"third.party@xyz.com"}},{"Phone":{"AreaCode":44,"CountryDialingCode":11,"Label":"mobile","PhoneNumber":888444444}}],"ContactType":"Payment","Individual":{"GivenName":"MIKE","NameTitle":"Mr","Surname":"TEST"}}]},"PassengerList":{"Passenger":[{"-PassengerID":"SH1","ContactInfoRef":"ContactInfo-SH1","Individual":{"Birthdate":"1982-12-15","Gender":"Male","GivenName":"one","NameTitle":"DR","Surname":"TEST"},"PTC":"ADT"},{"-PassengerID":"SH2","Individual":{"Birthdate":"1983-08-05","Gender":"Male","GivenName":"TWO","NameTitle":"DR","Surname":"TEST"},"InfantRef":"SH9","PTC":"ADT"},{"-PassengerID":"SH3","Individual":{"Birthdate":"1984-12-15","Gender":"Male","GivenName":"three","NameTitle":"MR","Surname":"TEST"},"InfantRef":"SH8","PTC":"ADT"},{"-PassengerID":"SH4","Individual":{"Birthdate":"2005-10-15","Gender":"Male","GivenName":"four","NameTitle":"MR","Surname":"TEST"},"PTC":"ADT"},{"-PassengerID":"SH5","Individual":{"Birthdate":"2005-12-15","Gender":"Male","GivenName":"five","NameTitle":"MR","Surname":"TEST"},"PTC":"ADT"},{"-PassengerID":"SH6","Individual":{"Birthdate":"2010-12-15","Gender":"Male","GivenName":"six","NameTitle":"MR","Surname":"TEST"},"PTC":"CHD"},{"-PassengerID":"SH7","Individual":{"Birthdate":"2012-12-15","Gender":"Male","GivenName":"seven","NameTitle":"MR","Surname":"TEST"},"PTC":"CHD"},{"-PassengerID":"SH8","Individual":{"Birthdate":"2017-12-15","Gender":"Male","GivenName":"eight","NameTitle":"MR","Surname":"TEST"},"PTC":"INF"},{"-PassengerID":"SH9","Individual":{"Birthdate":"2017-10-15","Gender":"Male","GivenName":"nine","NameTitle":"MR","Surname":"TEST"},"PTC":"INF"}]}},"Order":{"Offer":{"-OfferID":"OF-44a450e6-75b3-4443-9c18-1024137b082f","-Owner":"BA","-ResponseID":"tx-08-201-0d5cafdd-ea25-48ac-9a3b-42480defcad4","OfferItem":[{"-OfferItemID":"OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-1","PassengerRefs":"SH1 SH2 SH3"},{"-OfferItemID":"OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-2","PassengerRefs":"SH4 SH5"},{"-OfferItemID":"OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-3","PassengerRefs":"SH6 SH7"},{"-OfferItemID":"OF-44a450e6-75b3-4443-9c18-1024137b082f-OI-4","PassengerRefs":"SH8 SH9"}]}},"Payments":{"Payment":{"Amount":{"#text":528.7,"-Code":"GBP"},"Method":{"PaymentCard":{"CardCode":"MD","CardHolderBillingAddress":{"CityName":"LON","CountryCode":"GB","PostalCode":"TW345RT","Street":["Beeches Apartment","200 Lampton Road"]},"CardHolderName":"MR MIKE TEST","CardNumber":7777884566565720,"EffectiveExpireDate":{"Effective":1212,"Expiration":219},"SeriesCode":123,"Surcharge":{"Amount":{"#text":0,"-Code":"GBP"}}}},"Payer":{"ContactInfoRefs":"Payer"},"Type":"CC"}}}}}}}`)

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

	fmt.Println(tc.GetOutput("contentAsJson"))
	require.JSONEq(t, string(expected), fmt.Sprint(tc.GetOutput("contentAsJson")))

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

	require.JSONEq(t, string(expected), fmt.Sprint(tc.GetOutput("contentAsJson")))

}

//func TestAnyXML(t *testing.T) {
//
//	jsonObj, err := x2j.XmlToJson(xmlData, true)
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	fmt.Println(string(jsonObj[:len(jsonObj)]))
//
//	xmlObj, err := j2x.JsonToXml(jsonObj)
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	fmt.Println(string(xmlObj[:len(xmlObj)]))
//
//}
