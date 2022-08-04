package db

import (
	"Police-Check-Api/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomResident(t *testing.T) Resident {
	arg := CreateResidentParams{
		Nic:       util.RandomNic(),
		Fullname:  util.RandomName(),
		Raddress:  util.RandomAddress(),
		Clearance: util.RandomClear(),
	}

	resident, err := testQueries.CreateResident(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, resident)

	require.Equal(t, arg.Nic, resident.Nic)
	require.Equal(t, arg.Fullname, resident.Fullname)
	require.Equal(t, arg.Raddress, resident.Raddress)
	require.Equal(t, arg.Clearance, resident.Clearance)

	return resident
}

func TestCreateResident(t *testing.T) {
	createRandomResident(t)
}

func TestGetResident(t *testing.T) {
	resident1 := createRandomResident(t)
	resident2, err := testQueries.GetResident(context.Background(), resident1.Nic)
	require.NoError(t, err)
	require.NotEmpty(t, resident2)

	require.Equal(t, resident1.Nic, resident2.Nic)
	require.Equal(t, resident1.Fullname, resident2.Fullname)
	require.Equal(t, resident1.Raddress, resident2.Raddress)
	require.Equal(t, resident1.Clearance, resident2.Clearance)
}
