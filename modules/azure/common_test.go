package azure

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTargetAzureSubscription(t *testing.T) {
	t.Parallel()

	type args struct {
		subID string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "subIDProvidedAsArg", args: args{subID: "test"}, want: "test", wantErr: false},
		{name: "subIDNotProvided", args: args{subID: ""}, want: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTargetAzureSubscription(tt.args.subID)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGetTargetAzureResourceGroupName(t *testing.T) {
	t.Parallel()

	type args struct {
		rgName string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "rgNameProvidedAsArg", args: args{rgName: "test"}, want: "test", wantErr: false},
		{name: "rgNameNotProvided", args: args{rgName: ""}, want: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTargetAzureResourceGroupName(tt.args.rgName)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
