// Code generated by MockGen. DO NOT EDIT.
// Source: hex-example/ticket (interfaces: TicketRepository,TicketService,TicketHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	ticket "hex-example/ticket"
	http "net/http"
	reflect "reflect"
)

// MockTicketRepository is a mock of TicketRepository interface
type MockTicketRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTicketRepositoryMockRecorder
}

// MockTicketRepositoryMockRecorder is the mock recorder for MockTicketRepository
type MockTicketRepositoryMockRecorder struct {
	mock *MockTicketRepository
}

// NewMockTicketRepository creates a new mock instance
func NewMockTicketRepository(ctrl *gomock.Controller) *MockTicketRepository {
	mock := &MockTicketRepository{ctrl: ctrl}
	mock.recorder = &MockTicketRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTicketRepository) EXPECT() *MockTicketRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTicketRepository) Create(arg0 *ticket.Ticket) error {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockTicketRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTicketRepository)(nil).Create), arg0)
}

// FindAll mocks base method
func (m *MockTicketRepository) FindAll() ([]*ticket.Ticket, error) {
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*ticket.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockTicketRepositoryMockRecorder) FindAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTicketRepository)(nil).FindAll))
}

// FindById mocks base method
func (m *MockTicketRepository) FindById(arg0 string) (*ticket.Ticket, error) {
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*ticket.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById
func (mr *MockTicketRepositoryMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockTicketRepository)(nil).FindById), arg0)
}

// MockTicketService is a mock of TicketService interface
type MockTicketService struct {
	ctrl     *gomock.Controller
	recorder *MockTicketServiceMockRecorder
}

// MockTicketServiceMockRecorder is the mock recorder for MockTicketService
type MockTicketServiceMockRecorder struct {
	mock *MockTicketService
}

// NewMockTicketService creates a new mock instance
func NewMockTicketService(ctrl *gomock.Controller) *MockTicketService {
	mock := &MockTicketService{ctrl: ctrl}
	mock.recorder = &MockTicketServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTicketService) EXPECT() *MockTicketServiceMockRecorder {
	return m.recorder
}

// CreateTicket mocks base method
func (m *MockTicketService) CreateTicket(arg0 *ticket.Ticket) error {
	ret := m.ctrl.Call(m, "CreateTicket", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTicket indicates an expected call of CreateTicket
func (mr *MockTicketServiceMockRecorder) CreateTicket(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicket", reflect.TypeOf((*MockTicketService)(nil).CreateTicket), arg0)
}

// FindAllTickets mocks base method
func (m *MockTicketService) FindAllTickets() ([]*ticket.Ticket, error) {
	ret := m.ctrl.Call(m, "FindAllTickets")
	ret0, _ := ret[0].([]*ticket.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllTickets indicates an expected call of FindAllTickets
func (mr *MockTicketServiceMockRecorder) FindAllTickets() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllTickets", reflect.TypeOf((*MockTicketService)(nil).FindAllTickets))
}

// FindTicketById mocks base method
func (m *MockTicketService) FindTicketById(arg0 string) (*ticket.Ticket, error) {
	ret := m.ctrl.Call(m, "FindTicketById", arg0)
	ret0, _ := ret[0].(*ticket.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTicketById indicates an expected call of FindTicketById
func (mr *MockTicketServiceMockRecorder) FindTicketById(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTicketById", reflect.TypeOf((*MockTicketService)(nil).FindTicketById), arg0)
}

// MockTicketHandler is a mock of TicketHandler interface
type MockTicketHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTicketHandlerMockRecorder
}

// MockTicketHandlerMockRecorder is the mock recorder for MockTicketHandler
type MockTicketHandlerMockRecorder struct {
	mock *MockTicketHandler
}

// NewMockTicketHandler creates a new mock instance
func NewMockTicketHandler(ctrl *gomock.Controller) *MockTicketHandler {
	mock := &MockTicketHandler{ctrl: ctrl}
	mock.recorder = &MockTicketHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTicketHandler) EXPECT() *MockTicketHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockTicketHandler) Create(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.Call(m, "Create", arg0, arg1)
}

// Create indicates an expected call of Create
func (mr *MockTicketHandlerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTicketHandler)(nil).Create), arg0, arg1)
}

// Get mocks base method
func (m *MockTicketHandler) Get(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.Call(m, "Get", arg0, arg1)
}

// Get indicates an expected call of Get
func (mr *MockTicketHandlerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTicketHandler)(nil).Get), arg0, arg1)
}

// GetById mocks base method
func (m *MockTicketHandler) GetById(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.Call(m, "GetById", arg0, arg1)
}

// GetById indicates an expected call of GetById
func (mr *MockTicketHandlerMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockTicketHandler)(nil).GetById), arg0, arg1)
}
