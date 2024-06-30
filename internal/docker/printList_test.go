package docker

import (
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
)

func Test_getIPAddress(t *testing.T) {
	type args struct {
		c types.Container
	}

	sumNetSettingsNoAddresses := make(map[string]*network.EndpointSettings)

	sumNetSettingsOneAddresses := make(map[string]*network.EndpointSettings)
	sumNetSettingsOneAddresses["0"] = &network.EndpointSettings{IPAddress: "127.0.0.1", IPPrefixLen: 32}

	sumNetSettingsTwoAddresses := make(map[string]*network.EndpointSettings)
	sumNetSettingsTwoAddresses["0"] = &network.EndpointSettings{IPAddress: "127.0.0.1", IPPrefixLen: 32}
	sumNetSettingsTwoAddresses["1"] = &network.EndpointSettings{IPAddress: "127.0.1.1", IPPrefixLen: 8}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get ip address - no address",
			args: args{c: types.Container{
				NetworkSettings: &types.SummaryNetworkSettings{
					Networks: sumNetSettingsNoAddresses,
				},
			}},
			want: "",
		},
		{
			name: "get ip address - one address",
			args: args{c: types.Container{
				NetworkSettings: &types.SummaryNetworkSettings{
					Networks: sumNetSettingsOneAddresses,
				},
			}},
			want: "127.0.0.1/32",
		},
		{
			name: "get ip address - two addresses",
			args: args{c: types.Container{
				NetworkSettings: &types.SummaryNetworkSettings{
					Networks: sumNetSettingsTwoAddresses,
				},
			}},
			want: "127.0.0.1/32, 127.0.1.1/8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIPAddress(tt.args.c); got != tt.want {
				t.Errorf("getIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getContainerMounts(t *testing.T) {
	type args struct {
		c types.Container
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get mount points - no mount points",
			args: args{c: types.Container{
				Mounts: []types.MountPoint{},
			}},
			want: "",
		},
		{
			name: "get mount points - one mount point",
			args: args{c: types.Container{
				Mounts: []types.MountPoint{
					{Source: "/tmp", Destination: "/tmp"},
				},
			}},
			want: "/tmp:/tmp",
		},
		{
			name: "get mount points - many mount points",
			args: args{c: types.Container{
				Mounts: []types.MountPoint{
					{Source: "/tmp", Destination: "/tmp"},
					{Source: "/mnt", Destination: "/mnt"},
				},
			}},
			want: "/tmp:/tmp, /mnt:/mnt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getContainerMounts(tt.args.c); got != tt.want {
				t.Errorf("getContainerMounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
