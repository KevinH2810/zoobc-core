package util

import (
	"reflect"
	"strings"
	"testing"

	"github.com/zoobc/zoobc-core/common/model"
)

func TestNewHost(t *testing.T) {

	type args struct {
		address    string
		port       uint32
		knownPeers []*model.Peer
	}
	tests := []struct {
		name string
		args args
		want *model.Host
	}{
		// TODO: Add test cases.
		{
			name: "NewHostTest:NewHost",
			args: args{
				address: "127.0.0.1",
				port:    3000,
				knownPeers: []*model.Peer{
					{
						Info: &model.Node{
							Address:       "127.0.0.1",
							Port:          3001,
							SharedAddress: "127.0.0.1",
						},
					},
					{
						Info: &model.Node{
							Address:       "192.168.5.1",
							Port:          3002,
							SharedAddress: "192.168.5.1",
						},
					},
				},
			},
			want: &model.Host{
				Info: &model.Node{
					Address: "127.0.0.1",
					Port:    3000,
				},
				KnownPeers: map[string]*model.Peer{
					"127.0.0.1:3001": {
						Info: &model.Node{
							SharedAddress: "127.0.0.1",
							Address:       "127.0.0.1",
							Port:          3001,
						},
					},
					"192.168.5.1:3002": {
						Info: &model.Node{
							SharedAddress: "192.168.5.1",
							Address:       "192.168.5.1",
							Port:          3002,
						},
					},
				},
				UnresolvedPeers: map[string]*model.Peer{
					"127.0.0.1:3001": {
						Info: &model.Node{
							SharedAddress: "127.0.0.1",
							Address:       "127.0.0.1",
							Port:          3001,
						},
					},
					"192.168.5.1:3002": {
						Info: &model.Node{
							SharedAddress: "192.168.5.1",
							Address:       "192.168.5.1",
							Port:          3002,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHost(tt.args.address, tt.args.port, tt.args.knownPeers)
			if strings.Compare(got.String(), tt.want.String()) != 0 {
				t.Errorf("\n%v \n%v", got.String(), tt.want.String())
			}
		})
	}
}

func TestNewKnownPeer(t *testing.T) {
	type args struct {
		address string
		port    int
	}
	tests := []struct {
		name string
		args args
		want *model.Peer
	}{
		// TODO: Add test cases.
		{
			name: "NewKnownPeer:success",
			args: args{
				address: "127.0.0.1",
				port:    8001,
			},
			want: &model.Peer{
				Info: &model.Node{
					Address: "127.0.0.1",
					Port:    8001,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKnownPeer(tt.args.address, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKnownPeer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFullAddressPeer(t *testing.T) {
	type args struct {
		peer *model.Peer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "GetFullAddressPeer:success",
			args: args{
				peer: &model.Peer{
					Info: &model.Node{
						Address: "127.0.0.1",
						Port:    8001,
					},
				},
			},
			want: "127.0.0.1:8001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFullAddressPeer(tt.args.peer); got != tt.want {
				t.Errorf("GetFullAddressPeer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseKnownPeers(t *testing.T) {
	type args struct {
		peers []string
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.Peer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ParseKnownPeersTest:success",
			args: args{
				peers: []string{"192.168.1.2:2001", "192.168.5.123:3000"},
			},
			want:    append([]*model.Peer{}, NewKnownPeer("192.168.1.2", 2001), NewKnownPeer("192.168.5.123", 3000)),
			wantErr: false,
		},
		{
			name: "ParseKnownPeersTest:true",
			args: args{
				peers: []string{"192.168.1.2:2001xa", "192.168.5.123:3000a"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseKnownPeers(tt.args.peers)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseKnownPeers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ParseKnownPeers() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

// func TestAddToResolvedPeer(t *testing.T) {
// 	type args struct {
// 		host *model.Host
// 		peer *model.Peer
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *model.Host
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "AddToResolvedPeer:success",
// 			args: args{
// 				host: &model.Host{
// 					Info:            &model.Node{},
// 					KnownPeers:      make(map[string]*model.Peer),
// 					Peers:           make(map[string]*model.Peer),
// 					UnresolvedPeers: make(map[string]*model.Peer),
// 				},
// 				peer: &model.Peer{
// 					Info: &model.Node{
// 						Address: "127.0.0.1",
// 						Port:    8001,
// 					},
// 				},
// 			},
// 			want: &model.Host{
// 				Info:       &model.Node{},
// 				KnownPeers: make(map[string]*model.Peer),
// 				Peers: map[string]*model.Peer{
// 					"127.0.0.1:8001": {
// 						Info: &model.Node{
// 							Address: "127.0.0.1",
// 							Port:    8001,
// 						},
// 					},
// 				},
// 				UnresolvedPeers: make(map[string]*model.Peer),
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := AddToResolvedPeer(tt.args.host, tt.args.peer); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("AddToResolvedPeer() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestAddToUnresolvedPeers(t *testing.T) {
// 	type args struct {
// 		host     *model.Host
// 		newNodes []*model.Node
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *model.Host
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "AddToUnresolvedPeers:success",
// 			args: args{
// 				host: &model.Host{
// 					Info:            &model.Node{},
// 					KnownPeers:      make(map[string]*model.Peer),
// 					Peers:           make(map[string]*model.Peer),
// 					UnresolvedPeers: make(map[string]*model.Peer),
// 				},
// 				newNodes: []*model.Node{
// 					{
// 						Address: "127.0.0.1",
// 						Port:    8001,
// 					},
// 					{
// 						Address: "192.168.1.5",
// 						Port:    8001,
// 					},
// 				},
// 			},
// 			want: &model.Host{
// 				Info:       &model.Node{},
// 				KnownPeers: make(map[string]*model.Peer),
// 				Peers:      make(map[string]*model.Peer),
// 				UnresolvedPeers: map[string]*model.Peer{
// 					"127.0.0.1:8001": {
// 						Info: &model.Node{
// 							Address: "127.0.0.1",
// 							Port:    8001,
// 						},
// 					},
// 					"192.168.1.5:8001": {
// 						Info: &model.Node{
// 							Address: "192.168.1.5",
// 							Port:    8001,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := AddToUnresolvedPeers(tt.args.host, tt.args.newNodes)
// 			if strings.Compare(got.String(), tt.want.String()) != 0 {
// 				t.Errorf("AddToUnresolvedPeers() = \n%v, want \n%v", got.String(), tt.want.String())
// 			}
// 		})
// 	}
// }

// func TestPeerUnblacklist(t *testing.T) {
// 	type args struct {
// 		peer *model.Peer
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *model.Peer
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "TestPeerUnblacklist:success",
// 			args: args{
// 				peer: &model.Peer{
// 					BlacklistingCause: "not connected",
// 					BlacklistingTime:  1234,
// 					State:             model.PeerState_BLACKLISTED,
// 				},
// 			},
// 			want: &model.Peer{
// 				BlacklistingCause: "",
// 				BlacklistingTime:  0,
// 				State:             model.PeerState_NON_CONNECTED,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := PeerUnblacklist(tt.args.peer); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("PeerUnblacklist() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }