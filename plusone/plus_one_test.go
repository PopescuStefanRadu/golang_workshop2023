package plusone_test

import (
	"esol/must"
	"esol/plusone"
	mock_plusone "esol/plusone/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestPlusDataFromDB_AddWithDataFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	dbMock := mock_plusone.NewMockDbRepo(ctrl)
	subject := plusone.PlusDataFromDB{
		DbRepo: dbMock,
	}
	dbMock.EXPECT().GetNumber().Return(1).Times(1)

	res := subject.AddWithDataFromDB(1)
	require.Equal(t, 2, res)
}

func TestImpoundCar(t *testing.T) {
	ctrl := gomock.NewController(t)
	dbMock := mock_plusone.NewMockDbRepo(ctrl)
	subject := plusone.PlusDataFromDB{
		DbRepo: dbMock,
	}
	dbMock.EXPECT().FindCarByVIN("12345").Return(plusone.Car{
		VIN:             "12345",
		Model:           "Honda Civic",
		FabricationYear: must.ParseDate("2021-03-04"),
		Impounded:       false,
	}).Times(1)
	dbMock.EXPECT().UpdateCar(plusone.Car{
		VIN:             "12345",
		Model:           "Honda Civic",
		FabricationYear: must.ParseDate("2021-03-04"),
		Impounded:       true,
	}).Times(1)

	subject.ImpoundCar("12345")
}

func TestSendFunctionality(t *testing.T) {
	ctrl := gomock.NewController(t)
	dbMock := mock_plusone.NewMockDbRepo(ctrl)
	somethingMock := mock_plusone.NewMockSomethingThatTakesFunctionality(ctrl)
	subject := plusone.PlusDataFromDB{
		DbRepo:    dbMock,
		Something: somethingMock,
	}

	somethingMock.EXPECT().DoSomething("1234", gomock.Any()).Do(func(_ string, f func(s string)) {
		f("B-24-UDV")
	})

	got := subject.SendFunctionality("1234")
	require.Equal(t, "B-24-UDV", got)
}
