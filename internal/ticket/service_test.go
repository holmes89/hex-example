package ticket_test

import (
	"hex-example/internal/mocksnal/mocks"
	"hex-example/internal/ticketal/ticket"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func TestTicketServiceSuite(t *testing.T) {
	suite.Run(t, new(TicketServiceTestSuite))
}

type TicketServiceTestSuite struct {
	suite.Suite
	ticketRepo *mocks.MockTicketRepository
	underTest  ticket.TicketService
}

func (suite *TicketServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.ticketRepo = mocks.NewMockTicketRepository(mockCtrl)
	suite.underTest = ticket.NewTicketService(suite.ticketRepo)
}

func (suite *TicketServiceTestSuite) TestCreate() {
	//Arrange
	t := &ticket.Ticket{
		Creator: "Joel",
	}
	suite.ticketRepo.EXPECT().Create(gomock.Any()).Return(nil)

	//Act
	err := suite.underTest.CreateTicket(t)

	//Assert
	suite.NoError(err, "Shouldn't error")
	suite.NotNil(t.ID, "should not be null")
	suite.NotNil(t.Created, "should not be null")
	suite.NotNil(t.Updated, "should not be null")

}

func (suite *TicketServiceTestSuite) TestFindTicketById() {
	t := &ticket.Ticket{
		ID:      "test",
		Creator: "Joel",
	}
	suite.ticketRepo.EXPECT().FindById("test").Return(t, nil)

	result, err := suite.underTest.FindTicketById("test")

	suite.NoError(err, "Shouldn't error")
	suite.Equal(t, result, "should be pushing value returned from repo")
}

func (suite *TicketServiceTestSuite) TestFindAll() {
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
	suite.ticketRepo.EXPECT().FindAll().Return(ts, nil)

	result, err := suite.underTest.FindAllTickets()

	suite.NoError(err, "Shouldn't error")
	suite.Len(result, 2, "Should get two results")
}
