// ZooBC Copyright (C) 2020 Quasisoft Limited - Hong Kong
// This file is part of ZooBC <https://github.com/zoobc/zoobc-core>
//
// ZooBC is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ZooBC is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ZooBC.  If not, see <http://www.gnu.org/licenses/>.
//
// Additional Permission Under GNU GPL Version 3 section 7.
// As the special exception permitted under Section 7b, c and e,
// in respect with the Author’s copyright, please refer to this section:
//
// 1. You are free to convey this Program according to GNU GPL Version 3,
//     as long as you respect and comply with the Author’s copyright by
//     showing in its user interface an Appropriate Notice that the derivate
//     program and its source code are “powered by ZooBC”.
//     This is an acknowledgement for the copyright holder, ZooBC,
//     as the implementation of appreciation of the exclusive right of the
//     creator and to avoid any circumvention on the rights under trademark
//     law for use of some trade names, trademarks, or service marks.
//
// 2. Complying to the GNU GPL Version 3, you may distribute
//     the program without any permission from the Author.
//     However a prior notification to the authors will be appreciated.
//
// ZooBC is architected by Roberto Capodieci & Barton Johnston
//             contact us at roberto.capodieci[at]blockchainzoo.com
//             and barton.johnston[at]blockchainzoo.com
//
// Core developers that contributed to the current implementation of the
// software are:
//             Ahmad Ali Abdilah ahmad.abdilah[at]blockchainzoo.com
//             Allan Bintoro allan.bintoro[at]blockchainzoo.com
//             Andy Herman
//             Gede Sukra
//             Ketut Ariasa
//             Nawi Kartini nawi.kartini[at]blockchainzoo.com
//             Stefano Galassi stefano.galassi[at]blockchainzoo.com
//
// IMPORTANT: The above copyright notice and this permission notice
// shall be included in all copies or substantial portions of the Software.
package storage

import (
	"reflect"
	"sync"
	"testing"

	"github.com/zoobc/zoobc-core/common/chaintype"
)

func TestNewReceiptReminderStorage(t *testing.T) {
	tests := []struct {
		name string
		want *ReceiptReminderStorage
	}{
		{
			name: "TestNewReceiptReminderStorage:Success",
			want: &ReceiptReminderStorage{
				RWMutex:   sync.RWMutex{},
				reminders: map[string]chaintype.ChainType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReceiptReminderStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReceiptReminderStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReceiptReminderStorage_ClearCache(t *testing.T) {
	type fields struct {
		RWMutex   sync.RWMutex
		reminders map[string]chaintype.ChainType
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "TestReceiptReminderStorage_ClearCache:Success",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ReceiptReminderStorage{
				RWMutex:   tt.fields.RWMutex,
				reminders: tt.fields.reminders,
			}
			if err := rs.ClearCache(); (err != nil) != tt.wantErr {
				t.Errorf("ClearCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReceiptReminderStorage_GetAllItems(t *testing.T) {
	mockKey := map[string]chaintype.ChainType{}
	type fields struct {
		RWMutex   sync.RWMutex
		reminders map[string]chaintype.ChainType
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestReceiptReminderStorage_GetAllItems:Success",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: map[string]chaintype.ChainType{},
			},
			args: args{
				key: &mockKey,
			},
			wantErr: false,
		},
		{
			name: "TestReceiptReminderStorage_GetAllItems:Fail-WrongKey",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: nil,
			},
			args: args{
				key: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ReceiptReminderStorage{
				RWMutex:   tt.fields.RWMutex,
				reminders: tt.fields.reminders,
			}
			if err := rs.GetAllItems(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("GetAllItems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReceiptReminderStorage_GetItem(t *testing.T) {
	mockItem := chaintype.GetChainType(1)
	type fields struct {
		RWMutex   sync.RWMutex
		reminders map[string]chaintype.ChainType
	}
	type args struct {
		key  interface{}
		item interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestReceiptReminderStorage_GetItem:Success",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: map[string]chaintype.ChainType{},
			},
			args: args{
				key:  "test",
				item: &mockItem,
			},
			wantErr: false,
		},
		{
			name: "TestReceiptReminderStorage_GetItem:Fail-WrongTypeKey",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: map[string]chaintype.ChainType{},
			},
			args: args{
				key:  nil,
				item: nil,
			},
			wantErr: true,
		},
		{
			name: "TestReceiptReminderStorage_GetItem:Fail-WrongTypeItem",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: map[string]chaintype.ChainType{},
			},
			args: args{
				key:  "test",
				item: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ReceiptReminderStorage{
				RWMutex:   tt.fields.RWMutex,
				reminders: tt.fields.reminders,
			}
			if err := rs.GetItem(tt.args.key, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("GetItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReceiptReminderStorage_GetSize(t *testing.T) {
	type fields struct {
		RWMutex   sync.RWMutex
		reminders map[string]chaintype.ChainType
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "TestReceiptReminderStorage_GetSize",
			fields: fields{
				RWMutex: sync.RWMutex{},
				reminders: map[string]chaintype.ChainType{
					"1": chaintype.GetChainType(1),
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ReceiptReminderStorage{
				RWMutex:   tt.fields.RWMutex,
				reminders: tt.fields.reminders,
			}
			if got := rs.GetSize(); got != tt.want {
				t.Errorf("GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReceiptReminderStorage_GetTotalItems(t *testing.T) {
	type fields struct {
		RWMutex   sync.RWMutex
		reminders map[string]chaintype.ChainType
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "TestReceiptReminderStorage_GetTotalItems:Success",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ReceiptReminderStorage{
				RWMutex:   tt.fields.RWMutex,
				reminders: tt.fields.reminders,
			}
			if got := rs.GetTotalItems(); got != tt.want {
				t.Errorf("GetTotalItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReceiptReminderStorage_RemoveItem(t *testing.T) {
	type fields struct {
		RWMutex   sync.RWMutex
		reminders map[string]chaintype.ChainType
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestReceiptReminderStorage_RemoveItem:Success",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: make(map[string]chaintype.ChainType),
			},
			args: args{
				key: "test",
			},
			wantErr: false,
		},
		{
			name: "TestReceiptReminderStorage_RemoveItem:Fail-WrongTypeKey",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: nil,
			},
			args: args{
				key: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ReceiptReminderStorage{
				RWMutex:   tt.fields.RWMutex,
				reminders: tt.fields.reminders,
			}
			if err := rs.RemoveItem(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("RemoveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReceiptReminderStorage_SetItem(t *testing.T) {
	type fields struct {
		RWMutex   sync.RWMutex
		reminders map[string]chaintype.ChainType
	}
	type args struct {
		key  interface{}
		item interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestReceiptReminderStorage_SetItem:Success",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: map[string]chaintype.ChainType{},
			},
			args: args{
				key:  "test",
				item: chaintype.GetChainType(1),
			},
			wantErr: false,
		},
		{
			name: "TestReceiptReminderStorage_SetItem:Fail-WrongTypeKey",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: map[string]chaintype.ChainType{},
			},
			args: args{
				key:  1,
				item: chaintype.GetChainType(1),
			},
			wantErr: true,
		},
		{
			name: "TestReceiptReminderStorage_SetItem:Fail-WrongTypeItem",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: map[string]chaintype.ChainType{},
			},
			args: args{
				key:  "1",
				item: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ReceiptReminderStorage{
				RWMutex:   tt.fields.RWMutex,
				reminders: tt.fields.reminders,
			}
			if err := rs.SetItem(tt.args.key, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("SetItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReceiptReminderStorage_SetItems(t *testing.T) {
	type fields struct {
		RWMutex   sync.RWMutex
		reminders map[string]chaintype.ChainType
	}
	type args struct {
		in0 interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestReceiptReminderStorage_SetItems:Success",
			fields: fields{
				RWMutex:   sync.RWMutex{},
				reminders: nil,
			},
			args: args{
				in0: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ReceiptReminderStorage{
				RWMutex:   tt.fields.RWMutex,
				reminders: tt.fields.reminders,
			}
			if err := rs.SetItems(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("SetItems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
