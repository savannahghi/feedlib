package go_utils_test

import (
	"testing"

	base "github.com/savannahghi/go_utils"
)

func TestIsMSISDNValid(t *testing.T) {

	tests := []struct {
		name   string
		msisdn string
		want   bool
	}{
		{
			name:   "valid : kenyan with code",
			msisdn: "+254722000000",
			want:   true,
		},
		{
			name:   "valid : kenyan without code",
			msisdn: "0722000000",
			want:   true,
		},
		{
			name:   "valid : kenyan without code and spaces",
			msisdn: "0722 000 000",
			want:   true,
		},
		{
			name:   "valid : kenyan without plus sign",
			msisdn: "+254722000000",
			want:   true,
		},
		{
			name:   "valid : kenyan without plus sign and spaces",
			msisdn: "+254 722 000 000",
			want:   true,
		},
		{
			name:   "invalid : kenyan with alphanumeric1",
			msisdn: "+25472abc0000",
			want:   false,
		},
		{
			name:   "invalid : kenyan with alphanumeric2",
			msisdn: "072abc0000",
			want:   false,
		},
		{
			name:   "invalid : kenyan short length",
			msisdn: "0720000",
			want:   false,
		},
		{
			name:   "invalid : kenyan with unwanted characters : asterisk",
			msisdn: "072*120000",
			want:   false,
		},
		{
			name:   "invalid : kenyan without code with plus sign as prefix",
			msisdn: "+0722000000",
			want:   false,
		},
		{
			name:   "ivalid : international with alphanumeric",
			msisdn: "90191919qwe",
			want:   false,
		},
		{
			name:   "invalid : international with unwanted characters : asterisk",
			msisdn: "(+351) 282 *3 50 50",
			want:   false,
		},
		{
			name:   "invalid : international with unwanted characters : assorted",
			msisdn: "(+351) $82 *3 50 50",
			want:   false,
		},
		{
			name:   "valid : usa number",
			msisdn: "+12028569601",
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base.IsMSISDNValid(tt.msisdn); got != tt.want {
				t.Errorf("IsMSISDNValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSliceContains(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "slice which contains the int",
			args: args{
				s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				e: 7,
			},
			want: true,
		},
		{
			name: "slice which does NOT contain the int",
			args: args{
				s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				e: 79,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base.IntSliceContains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("IntSliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSliceContains(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string found in slice",
			args: args{
				s: []string{"a", "b", "c", "d", "e"},
				e: "a",
			},
			want: true,
		},
		{
			name: "string not found in slice",
			args: args{
				s: []string{"a", "b", "c", "d", "e"},
				e: "z",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base.StringSliceContains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("StringSliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeMSISDN(t *testing.T) {
	type args struct {
		msisdn string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "good Kenyan number, full E164 format",
			args: args{
				"+254723002959",
			},
			want:    "+254723002959",
			wantErr: false,
		},
		{
			name: "good Kenyan number, no + prefix",
			args: args{
				"254723002959",
			},
			want:    "+254723002959",
			wantErr: false,
		},
		{
			name: "good Kenyan number, no international dialling code",
			args: args{
				"0723002959",
			},
			want:    "+254723002959",
			wantErr: false,
		},
		{
			name: "good US number, full E164 format",
			args: args{
				"+16125409037",
			},
			want:    "+16125409037",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := base.NormalizeMSISDN(tt.args.msisdn)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeMSISDN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if *got != tt.want {
				t.Errorf("NormalizeMSISDN() = %v, want %v", got, tt.want)
			}
		})
	}
}
