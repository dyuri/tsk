// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kakengloh/tsk/repository (interfaces: TaskRepository)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/kakengloh/tsk/entity"
	repository "github.com/kakengloh/tsk/repository"
)

// MockTaskRepository is a mock of TaskRepository interface.
type MockTaskRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepositoryMockRecorder
}

// MockTaskRepositoryMockRecorder is the mock recorder for MockTaskRepository.
type MockTaskRepositoryMockRecorder struct {
	mock *MockTaskRepository
}

// NewMockTaskRepository creates a new mock instance.
func NewMockTaskRepository(ctrl *gomock.Controller) *MockTaskRepository {
	mock := &MockTaskRepository{ctrl: ctrl}
	mock.recorder = &MockTaskRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRepository) EXPECT() *MockTaskRepositoryMockRecorder {
	return m.recorder
}

// AddComment mocks base method.
func (m *MockTaskRepository) AddComment(arg0 int, arg1 string) (entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", arg0, arg1)
	ret0, _ := ret[0].(entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment.
func (mr *MockTaskRepositoryMockRecorder) AddComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockTaskRepository)(nil).AddComment), arg0, arg1)
}

// BulkDeleteTasks mocks base method.
func (m *MockTaskRepository) BulkDeleteTasks(arg0 ...int) map[int]error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BulkDeleteTasks", varargs...)
	ret0, _ := ret[0].(map[int]error)
	return ret0
}

// BulkDeleteTasks indicates an expected call of BulkDeleteTasks.
func (mr *MockTaskRepositoryMockRecorder) BulkDeleteTasks(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteTasks", reflect.TypeOf((*MockTaskRepository)(nil).BulkDeleteTasks), arg0...)
}

// CreateTask mocks base method.
func (m *MockTaskRepository) CreateTask(arg0 string, arg1 entity.TaskPriority, arg2 entity.TaskStatus, arg3 string) (entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskRepositoryMockRecorder) CreateTask(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskRepository)(nil).CreateTask), arg0, arg1, arg2, arg3)
}

// DeleteTask mocks base method.
func (m *MockTaskRepository) DeleteTask(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockTaskRepositoryMockRecorder) DeleteTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTaskRepository)(nil).DeleteTask), arg0)
}

// GetTaskByID mocks base method.
func (m *MockTaskRepository) GetTaskByID(arg0 int) (entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByID", arg0)
	ret0, _ := ret[0].(entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByID indicates an expected call of GetTaskByID.
func (mr *MockTaskRepositoryMockRecorder) GetTaskByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByID", reflect.TypeOf((*MockTaskRepository)(nil).GetTaskByID), arg0)
}

// ListTasks mocks base method.
func (m *MockTaskRepository) ListTasks() (entity.TaskList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTasks")
	ret0, _ := ret[0].(entity.TaskList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks.
func (mr *MockTaskRepositoryMockRecorder) ListTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockTaskRepository)(nil).ListTasks))
}

// SearchTasks mocks base method.
func (m *MockTaskRepository) SearchTasks(arg0 string) (entity.TaskList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTasks", arg0)
	ret0, _ := ret[0].(entity.TaskList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchTasks indicates an expected call of SearchTasks.
func (mr *MockTaskRepositoryMockRecorder) SearchTasks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTasks", reflect.TypeOf((*MockTaskRepository)(nil).SearchTasks), arg0)
}

// UpdateTask mocks base method.
func (m *MockTaskRepository) UpdateTask(arg0 int, arg1 string, arg2 entity.TaskPriority, arg3 entity.TaskStatus) (entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockTaskRepositoryMockRecorder) UpdateTask(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockTaskRepository)(nil).UpdateTask), arg0, arg1, arg2, arg3)
}

// UpdateTaskStatus mocks base method.
func (m *MockTaskRepository) UpdateTaskStatus(arg0 entity.TaskStatus, arg1 ...int) []repository.UpdateTaskStatusResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTaskStatus", varargs...)
	ret0, _ := ret[0].([]repository.UpdateTaskStatusResult)
	return ret0
}

// UpdateTaskStatus indicates an expected call of UpdateTaskStatus.
func (mr *MockTaskRepositoryMockRecorder) UpdateTaskStatus(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTaskStatus", reflect.TypeOf((*MockTaskRepository)(nil).UpdateTaskStatus), varargs...)
}
