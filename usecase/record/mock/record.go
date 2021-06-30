// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/record/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	entity "github.com/erbilsilik/getir-go-challange/entity"
	record "github.com/erbilsilik/getir-go-challange/usecase/record"
	gomock "github.com/golang/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange mocks base method.
func (m *MockReader) GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(query *record.RecordsFilteredByTimeAndTotalCountInGivenNumberRangeQuery) ([]*entity.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange", query)
	ret0, _ := ret[0].([]*entity.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange indicates an expected call of GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange.
func (mr *MockReaderMockRecorder) GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange", reflect.TypeOf((*MockReader)(nil).GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange), query)
}

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange mocks base method.
func (m *MockRepository) GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(query *record.RecordsFilteredByTimeAndTotalCountInGivenNumberRangeQuery) ([]*entity.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange", query)
	ret0, _ := ret[0].([]*entity.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange indicates an expected call of GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange.
func (mr *MockRepositoryMockRecorder) GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange", reflect.TypeOf((*MockRepository)(nil).GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange), query)
}

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange mocks base method.
func (m *MockUseCase) GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(query *record.RecordsFilteredByTimeAndTotalCountInGivenNumberRangeQuery) ([]*entity.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange", query)
	ret0, _ := ret[0].([]*entity.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange indicates an expected call of GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange.
func (mr *MockUseCaseMockRecorder) GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange", reflect.TypeOf((*MockUseCase)(nil).GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange), query)
}