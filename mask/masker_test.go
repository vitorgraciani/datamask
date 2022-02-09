package mask

import (
	"reflect"
	"testing"
)

func TestMask(t *testing.T) {
	type Address struct {
		Name string `mask:"last"`
	}

	type User struct {
		Name             string  `json:"fullName" mask:"middle"`
		LastName         string  `json:"lastName" mask:"initial"`
		Nickname         string  `show:"last" mask:"last"`
		CreditCardNumber string  `json:"creditCardNumber"`
		Address          Address `mask:"struct"`
		RG               string  `json:"rg" mask:"all"`
		Hobby            string  `json:"hobby" mask:"firstLetter"`
		CPF              string  `mask:"lastLetter"`
		Email            string  `json:"personalEmail" mask:"email"`
	}

	type args struct {
		s interface{}
	}
	tests := []struct {
		name    string
		m       *Mask
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Nil Input",
			m:    NewMask(),
			args: args{
				s: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "String Fields",
			m:    NewMask(),
			args: args{
				s: &User{
					Name:             "Arnaldo Cesar",
					LastName:         "Junqueira",
					Nickname:         "Arnaldinho",
					CreditCardNumber: "1234567890",
					Address: Address{
						Name: "Marginal Pinheiros",
					},
					RG:    "123458769",
					Hobby: "Cervejeiro",
					CPF:   "43578689000",
					Email: "arnaldinho@gmail.com",
				},
			},
			want: &User{
				Name:             "Arn**** **sar",
				LastName:         "****ueira",
				Nickname:         "Arnal*****",
				CreditCardNumber: "1234567890",
				Address: Address{
					Name: "Marginal *********",
				},
				RG:    "*********",
				Hobby: "*****jeiro",
				CPF:   "43578******",
				Email: "ar**********@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "Struct without tag",
			m:    NewMask(),
			args: args{
				s: &User{
					CreditCardNumber: "678909890909",
				},
			},
			want: &User{
				CreditCardNumber: "678909890909",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMask()
			got, err := m.createMask(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Masker.Struct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Masker.Struct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowMask(t *testing.T) {
	type Address struct {
		Name string `show:"last"`
	}

	type User struct {
		Name             string `json:"fullName" show:"middle"`
		LastName         string `json:"lastName" show:"initial"`
		Nickname         string `json:"nickname" show:"last"`
		CreditCardNumber string `json:"creditCardNumber"`
		Address          Address
		RG               string `json:"rg" show:"all"`
		Hobby            string `json:"hobby" show:"firstLetter"`
		CPF              string `show:"lastLetter"`
		Email            string `json:"personalEmail" show:"email"`
	}

	type args struct {
		s interface{}
	}
	tests := []struct {
		name    string
		m       *Show
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Nil Input",
			m:    NewShow(),
			args: args{
				s: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "String Fields",
			m:    NewShow(),
			args: args{
				s: &User{
					Name:             "Arnaldo Cesar",
					LastName:         "Junqueira",
					Nickname:         "Arnaldinho",
					CreditCardNumber: "1234567890",
					Address: Address{
						Name: "Marginal Pinheiros",
					},
					RG:    "123458769",
					Hobby: "Cervejeiro",
					CPF:   "43578689000",
					Email: "arnaldinho@gmail.com",
				},
			},
			want: &User{
				Name:             "***aldo Ce***",
				LastName:         "Junq*****",
				Nickname:         "*****dinho",
				CreditCardNumber: "**********",
				Address: Address{
					Name: "******** Pinheiros",
				},
				RG:    "123458769",
				Hobby: "Cerve*****",
				CPF:   "*****689000",
				Email: "ar**********@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "Struct without tag",
			m:    NewShow(),
			args: args{
				s: &User{
					CreditCardNumber: "678909890909",
				},
			},
			want: &User{
				CreditCardNumber: "************",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewShow()
			got, err := m.createShowData(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Masker.Struct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Masker.Struct() = %v, want %v", got, tt.want)
			}
		})
	}
}
