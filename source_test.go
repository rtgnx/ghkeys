package ghkey

import (
	"testing"
)

func TestKeys(t *testing.T) {
	type args struct {
		name string
		cfg  Config
	}
	tests := []struct {
		name     string
		args     args
		wantKeys bool
		wantErr  bool
	}{
		{
			name:     "test allowed user",
			args:     args{name: "rtgnx", cfg: Config{AllowedUsers: []string{"rtgnx"}}},
			wantErr:  false,
			wantKeys: true,
		},
		{
			name:     "test not allowd user",
			args:     args{name: "rtgnx", cfg: Config{AllowedUsers: []string{"abc"}}},
			wantErr:  true,
			wantKeys: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKeys, err := Keys(tt.args.name, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (tt.wantKeys && len(gotKeys) == 0) || (!tt.wantKeys && len(gotKeys) > 0) {
				t.Errorf("Keys() = %v, want %v", gotKeys, tt.wantKeys)
			}

		})
	}
}
