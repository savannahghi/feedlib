package go_utils_test

import (
	"bytes"
	"os"
	"strconv"
	"testing"

	base "github.com/savannahghi/go_utils"
	"github.com/stretchr/testify/assert"
)

func TestGender_String(t *testing.T) {
	tests := []struct {
		name string
		e    base.Gender
		want string
	}{
		{
			name: "male",
			e:    base.GenderMale,
			want: "male",
		},
		{
			name: "female",
			e:    base.GenderFemale,
			want: "female",
		},
		{
			name: "unknown",
			e:    base.GenderUnknown,
			want: "unknown",
		},
		{
			name: "other",
			e:    base.GenderOther,
			want: "other",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Gender.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGender_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    base.Gender
		want bool
	}{
		{
			name: "valid male",
			e:    base.GenderMale,
			want: true,
		},
		{
			name: "invalid gender",
			e:    base.Gender("this is not a real gender"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("Gender.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGender_UnmarshalGQL(t *testing.T) {
	female := base.GenderFemale
	invalid := base.Gender("")
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *base.Gender
		args    args
		wantErr bool
	}{
		{
			name: "valid female gender",
			e:    &female,
			args: args{
				v: "female",
			},
			wantErr: false,
		},
		{
			name: "invalid gender",
			e:    &invalid,
			args: args{
				v: "this is not a real gender",
			},
			wantErr: true,
		},
		{
			name: "non string gender",
			e:    &invalid,
			args: args{
				v: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Gender.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGender_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     base.Gender
		wantW string
	}{
		{
			name:  "valid unknown gender enum",
			e:     base.GenderUnknown,
			wantW: strconv.Quote("unknown"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Gender.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestContentType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    base.ContentType
		want bool
	}{
		{
			name: "good case",
			e:    base.ContentTypeJpg,
			want: true,
		},
		{
			name: "bad case",
			e:    base.ContentType("not a real content type"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("ContentType.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentType_String(t *testing.T) {
	tests := []struct {
		name string
		e    base.ContentType
		want string
	}{
		{
			name: "default case",
			e:    base.ContentTypePdf,
			want: "PDF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("ContentType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContentType_UnmarshalGQL(t *testing.T) {
	var sc base.ContentType
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *base.ContentType
		args    args
		wantErr bool
	}{
		{
			name: "valid unmarshal",
			e:    &sc,
			args: args{
				v: "PDF",
			},
			wantErr: false,
		},
		{
			name: "invalid unmarshal",
			e:    &sc,
			args: args{
				v: "this is not a valid scalar value",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ContentType.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContentType_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     base.ContentType
		wantW string
	}{
		{
			name:  "default case",
			e:     base.ContentTypePdf,
			wantW: strconv.Quote("PDF"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ContentType.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestLanguage_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    base.Language
		want bool
	}{
		{
			name: "good case",
			e:    base.LanguageEn,
			want: true,
		},
		{
			name: "bad case",
			e:    base.Language("not a real language"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("Language.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanguage_String(t *testing.T) {
	tests := []struct {
		name string
		e    base.Language
		want string
	}{
		{
			name: "default case",
			e:    base.LanguageEn,
			want: "en",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Language.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanguage_UnmarshalGQL(t *testing.T) {
	var sc base.Language

	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *base.Language
		args    args
		wantErr bool
	}{
		{
			name: "valid unmarshal",
			e:    &sc,
			args: args{
				v: "en",
			},
			wantErr: false,
		},
		{
			name: "invalid unmarshal",
			e:    &sc,
			args: args{
				v: "this is not a valid scalar value",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Language.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLanguage_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     base.Language
		wantW string
	}{
		{
			name:  "default case",
			e:     base.LanguageEn,
			wantW: strconv.Quote("en"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Language.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestPractitionerSpecialty_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    base.PractitionerSpecialty
		want bool
	}{
		{
			name: "good case",
			e:    base.PractitionerSpecialtyAnaesthesia,
			want: true,
		},
		{
			name: "bad case",
			e:    base.PractitionerSpecialty("not a real specialty"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("PractitionerSpecialty.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPractitionerSpecialty_String(t *testing.T) {
	tests := []struct {
		name string
		e    base.PractitionerSpecialty
		want string
	}{
		{
			name: "default case",
			e:    base.PractitionerSpecialtyAnaesthesia,
			want: "ANAESTHESIA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("PractitionerSpecialty.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPractitionerSpecialty_UnmarshalGQL(t *testing.T) {
	var sc base.PractitionerSpecialty
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *base.PractitionerSpecialty
		args    args
		wantErr bool
	}{
		{
			name: "valid unmarshal",
			e:    &sc,
			args: args{
				v: "ANAESTHESIA",
			},
			wantErr: false,
		},
		{
			name: "invalid unmarshal",
			e:    &sc,
			args: args{
				v: "this is not a valid scalar value",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("PractitionerSpecialty.UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPractitionerSpecialty_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     base.PractitionerSpecialty
		wantW string
	}{
		{
			name:  "default case",
			e:     base.PractitionerSpecialtyAnaesthesia,
			wantW: strconv.Quote("ANAESTHESIA"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("PractitionerSpecialty.MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestContentType(t *testing.T) {
	type expects struct {
		isValid      bool
		canUnmarshal bool
	}

	cases := []struct {
		name        string
		args        base.ContentType
		convert     interface{}
		expectation expects
	}{
		{
			name:    "invalid_string",
			args:    "testcontent",
			convert: "testcontent",
			expectation: expects{
				isValid:      false,
				canUnmarshal: false,
			},
		},
		{
			name:    "invalid_int_convert",
			args:    "testaddres",
			convert: 101,
			expectation: expects{
				isValid:      false,
				canUnmarshal: false,
			},
		},
		{
			name:    "valid",
			args:    base.ContentTypePng,
			convert: base.ContentTypePng,
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
		{
			name:    "valid_no_convert",
			args:    base.ContentTypePdf,
			convert: "testaddress",
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
		{
			name:    "valid_can_convert",
			args:    base.ContentTypePdf,
			convert: base.ContentTypePdf,
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation.isValid, tt.args.IsValid())
			assert.NotEmpty(t, tt.args.String())
			err := tt.args.UnmarshalGQL(tt.convert)
			assert.NotNil(t, err)
			tt.args.MarshalGQL(os.Stdout)

		})
	}

}

func TestLanguage(t *testing.T) {
	type expects struct {
		isValid      bool
		canUnmarshal bool
	}

	cases := []struct {
		name        string
		args        base.Language
		convert     interface{}
		expectation expects
	}{
		{
			name:    "invalid_string",
			args:    "testcontent",
			convert: "testcontent",
			expectation: expects{
				isValid:      false,
				canUnmarshal: false,
			},
		},
		{
			name:    "invalid_int_convert",
			args:    "testaddres",
			convert: 101,
			expectation: expects{
				isValid:      false,
				canUnmarshal: false,
			},
		},
		{
			name:    "valid",
			args:    base.LanguageEn,
			convert: base.LanguageEn,
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
		{
			name:    "valid_no_convert",
			args:    base.LanguageSw,
			convert: "testaddress",
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
		{
			name:    "valid_can_convert",
			args:    base.LanguageSw,
			convert: base.LanguageSw,
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation.isValid, tt.args.IsValid())
			assert.NotEmpty(t, tt.args.String())
			err := tt.args.UnmarshalGQL(tt.convert)
			assert.NotNil(t, err)
			tt.args.MarshalGQL(os.Stdout)

		})
	}

}

func TestCalendarView(t *testing.T) {
	type expects struct {
		isValid      bool
		canUnmarshal bool
	}

	cases := []struct {
		name        string
		args        base.CalendarView
		convert     interface{}
		expectation expects
	}{
		{
			name:    "invalid_string",
			args:    "testcontent",
			convert: "testcontent",
			expectation: expects{
				isValid:      false,
				canUnmarshal: false,
			},
		},
		{
			name:    "invalid_int_convert",
			args:    "testaddres",
			convert: 101,
			expectation: expects{
				isValid:      false,
				canUnmarshal: false,
			},
		},
		{
			name:    "valid",
			args:    base.CalendarViewDay,
			convert: base.CalendarViewDay,
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
		{
			name:    "valid_no_convert",
			args:    base.CalendarViewWeek,
			convert: "testaddress",
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
		{
			name:    "valid_can_convert",
			args:    base.CalendarViewWeek,
			convert: base.CalendarViewWeek,
			expectation: expects{
				isValid:      true,
				canUnmarshal: true,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation.isValid, tt.args.IsValid())
			assert.NotEmpty(t, tt.args.String())
			err := tt.args.UnmarshalGQL(tt.convert)
			assert.NotNil(t, err)
			tt.args.MarshalGQL(os.Stdout)

		})
	}
}

func TestIdentificationDocType_String(t *testing.T) {
	tests := []struct {
		name string
		e    base.IdentificationDocType
		want string
	}{
		{
			name: "NATIONALID",
			e:    base.IdentificationDocTypeNationalid,
			want: "NATIONALID",
		},
		{
			name: "PASSPORT",
			e:    base.IdentificationDocTypePassport,
			want: "PASSPORT",
		},
		{
			name: "MILITARY",
			e:    base.IdentificationDocTypeMilitary,
			want: "MILITARY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdentificationDocType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		e    base.IdentificationDocType
		want bool
	}{
		{
			name: "valid",
			e:    base.IdentificationDocTypeMilitary,
			want: true,
		},
		{
			name: "invalid",
			e:    base.IdentificationDocType("this is not real"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdentificationDocType_UnmarshalGQL(t *testing.T) {
	valid := base.IdentificationDocTypeNationalid
	invalid := base.IdentificationDocType("this is not real")
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		e       *base.IdentificationDocType
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			e:    &valid,
			args: args{
				v: "NATIONALID",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			e:    &invalid,
			args: args{
				v: "this is not real",
			},
			wantErr: true,
		},
		{
			name: "non string",
			e:    &invalid,
			args: args{
				v: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalGQL(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIdentificationDocType_MarshalGQL(t *testing.T) {
	tests := []struct {
		name  string
		e     base.IdentificationDocType
		wantW string
	}{
		{
			name:  "valid",
			e:     base.IdentificationDocTypeNationalid,
			wantW: strconv.Quote("NATIONALID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.e.MarshalGQL(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("MarshalGQL() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
