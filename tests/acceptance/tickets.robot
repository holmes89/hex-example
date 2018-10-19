*** Settings ***
Library	     Collections
Library	     RequestsLibrary
Library	     String

*** Variables ***
${HOST}	      http://localhost:3000
${ticket_id}

*** Test Cases ***
Create Ticket
       [Tags]	create	tickets
       ${json_string}=	catenate
       ...  {
       ...    "creator": "Joel",
       ...    "title": "Test Title",
       ...    "description": "A Test ticket",
       ...    "points": 5
       ...  }
       Create Session	ticket	${HOST}
       &{headers}=	Create Dictionary	Content-Type=application/json
       ${json}=		Evaluate		json.loads('''${json_string}''')		json
       ${resp}=		Post Request		ticket					/tickets	data=${json}	headers=${headers}
       Should Be Equal As Strings		${resp.status_code}			201
       Dictionary Should Contain Value		${resp.json()}				Joel
       Set Suite Variable	 ${ticket_id}	${resp.json()['id']}

Get Ticket
	[Tags]	get	tickets
	Create Session	ticket	${HOST}
	${resp}=	Get Request		ticket		/tickets/${ticket_id}
	Should Be Equal As Strings		${resp.status_code}	200
	Dictionary Should Contain Value		${resp.json()}		${ticket_id}
	Dictionary Should Contain Value		${resp.json()}		Joel
