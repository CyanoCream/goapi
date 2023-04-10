package service


import (
"errors"
"sesi_8/model"
"sesi_8/repository/mocks"
"testing"

"github.com/golang/mock/gomock"
"github.com/stretchr/testify/assert"
)

func Test_EmployeeService_GetEmployeeByID(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult model.Employee
		expectedError  error
		onEmployeeRepo func(mock *mocks.MockEmployeeRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onEmployeeRepo: func(mock *mocks.MockEmployeeRepo) {
			mock.EXPECT().GetEmployeeByID(gomock.Any()).Return(model.Employee{
				ID:       1,
				Fullname: "Rizqi",
				Email:    "rizqi@gmail.com",
				Age:      17,
				Division: "Tech",
			}, nil).Times(1)
		},
		expectedResult: model.Employee{
			ID:       1,
			Fullname: "Rizqi",
			Email:    "rizqi@gmail.com",
			Age:      17,
			Division: "Tech",
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onEmployeeRepo: func(mock *mocks.MockEmployeeRepo) {
			mock.EXPECT().GetEmployeeByID(gomock.Any()).Return(model.Employee{}, errors.New("record not found")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:          "unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onEmployeeRepo: func(mock *mocks.MockEmployeeRepo) {
			mock.EXPECT().GetEmployeeByID(gomock.Any()).Return(model.Employee{}, errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			employeeRepo := mocks.NewMockEmployeeRepo(mockCtrl)

			if testCase.onEmployeeRepo != nil {
				testCase.onEmployeeRepo(employeeRepo)
			}

			service := Service{
				repo: employeeRepo,
			}

			res, err := service.GetEmployeeByID(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_EmployeeService_CreateEmployee(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          model.Employee
		expectedResult model.Employee
		expectedError  error
		onEmployeeRepo func(mock *mocks.MockEmployeeRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: model.Employee{
			Fullname: "Rizqi",
			Email:    "rizqi@gmail.com",
			Age:      17,
			Division: "Tech",
		},
		onEmployeeRepo: func(mock *mocks.MockEmployeeRepo) {
			mock.EXPECT().CreateEmployee(gomock.Any()).Return(model.Employee{
				ID:       1,
				Fullname: "Rizqi",
				Email:    "rizqi@gmail.com",
				Age:      17,
				Division: "Tech",
			}, nil).Times(1)
		},
		expectedResult: model.Employee{
			ID:       1,
			Fullname: "Rizqi",
			Email:    "rizqi@gmail.com",
			Age:      17,
			Division: "Tech",
		},
	})

	testTable = append(testTable, testCase{
		name:      "unexpected error",
		wantError: true,
		input: model.Employee{
			Fullname: "Rizqi",
			Email:    "rizqi@gmail.com",
			Age:      17,
			Division: "Tech",
		},
		expectedError: errors.New("unexpected error"),
		onEmployeeRepo: func(mock *mocks.MockEmployeeRepo) {
			mock.EXPECT().CreateEmployee(gomock.Any()).Return(model.Employee{}, errors.New("unexpected error")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:      "invalid fullname length",
		wantError: true,
		input: model.Employee{
			Fullname: "Ade", // case negative
			Email:    "ade@gmail.com",
			Age:      17,
			Division: "Tech",
		},
		expectedError: errors.New("invalid fullname length"),
	})

	testTable = append(testTable, testCase{
		name:      "invalid division",
		wantError: true,
		input: model.Employee{
			Fullname: "Anang",
			Email:    "anang@gmail.com",
			Age:      17,
			Division: "Marketing", // case negative
		},
		expectedError: errors.New("invalid division"),
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			employeeRepo := mocks.NewMockEmployeeRepo(mockCtrl)

			if testCase.onEmployeeRepo != nil {
				testCase.onEmployeeRepo(employeeRepo)
			}

			service := Service{
				repo: employeeRepo,
			}

			res, err := service.CreateEmployee(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}
