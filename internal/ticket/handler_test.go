package ticket_test

import (
	"bytes"
	"encoding/json"
	"hex-example/internal/mocksnal/mocks"
	"hex-example/internal/ticketal/ticket"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func TestTicketHandlerSuite(t *testing.T) {
	suite.Run(t, new(TicketHandlerTestSuite))
}

type TicketHandlerTestSuite struct {
	suite.Suite
	ticketService *mocks.MockTicketService
	underTest     ticket.TicketHandler
}

func (suite *TicketHandlerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.ticketService = mocks.NewMockTicketService(mockCtrl)
	suite.underTest = ticket.NewTicketHandler(suite.ticketService)
}

func (suite *TicketHandlerTestSuite) TestCreate() {
	t := &ticket.Ticket{
		Creator: "Joel",
	}
	suite.ticketService.EXPECT().CreateTicket(gomock.Eq(t)).Return(nil)

	body, _ := json.Marshal(t)
	r, _ := http.NewRequest("POST", "/tickets", bytes.NewBuffer(body))

	w := httptest.NewRecorder()
	suite.underTest.Create(w, r)

	response := w.Result()
	suite.Equal("201 Created", response.Status)

	defer response.Body.Close()
	result := new(ticket.Ticket)
	json.NewDecoder(response.Body).Decode(result)

	suite.Equal("Joel", result.Creator)
}

func (suite *TicketHandlerTestSuite) TestFindTicketById() {
	t := &ticket.Ticket{
		Creator: "Joel",
	}
	suite.ticketService.EXPECT().FindTicketById("test").Return(t, nil)

	vars := map[string]string{
		"id": "test",
	}

	r, _ := http.NewRequest("GET", "/tickets/test", nil)
	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()
	suite.underTest.GetById(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)

	defer response.Body.Close()
	result := new(ticket.Ticket)
	json.NewDecoder(response.Body).Decode(result)

	suite.Equal("Joel", result.Creator)
}

func (suite *TicketHandlerTestSuite) TestFindAll() {
	ts := []*ticket.Ticket{
		&ticket.Ticket{
			ID:      "test1",
			Creator: "Joel",
		},
		&ticket.Ticket{
			ID:      "test2",
			Creator: "Other",
		},
	}
	suite.ticketService.EXPECT().FindAllTickets().Return(ts, nil)

	r, _ := http.NewRequest("GET", "/tickets", nil)

	w := httptest.NewRecorder()
	suite.underTest.Get(w, r)

	response := w.Result()
	suite.Equal("200 OK", response.Status)

	defer response.Body.Close()
	result := new([]ticket.Ticket)
	json.NewDecoder(response.Body).Decode(result)
	suite.Len(*result, 2, "Should get two results")
}
