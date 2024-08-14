package utils

import (
	"bytes"
	"io"
	"testing"
)

func MockReader(data string) io.Reader {
	return bytes.NewBufferString(data)
}

func TestGetTCPStates(t *testing.T) {
	mockData := `sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode
0: 3500007F:0035 00000000:0000 0A 00000000:00000000 00:00000000 00000000   101        0 21804 1 0000000000000000 100 0 0 10 5
1: 0100007F:0277 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 2370 1 0000000000000000 100 0 0 10 0
2: 00000000:9329 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 18878 1 0000000000000000 100 0 0 10 0
3: 4700A8C0:E842 BA70C923:01BB 06 00000000:00000000 03:00001486 00000000     0        0 0 3 0000000000000000
4: 4700A8C0:D30A B3DA1203:01BB 01 00000000:00000000 02:00000D00 00000000  1000        0 49990 2 0000000000000000 24 4 28 10 -1
5: 4700A8C0:D312 B3DA1203:01BB 01 00000000:00000000 02:0000094C 00000000  1000        0 45613 2 0000000000000000 24 4 30 10 -1
6: 4700A8C0:E794 01358F03:01BB 01 00000000:00000000 02:00000AD6 00000000  1000        0 335391 2 0000000000000000 24 4 28 10 -1
7: 4700A8C0:C036 0A834D63:01BB 01 00000000:00000000 02:00000A5B 00000000  1000        0 57850 3 0000000000000000 23 4 0 4 4
8: 4700A8C0:D31C B3DA1203:01BB 01 00000000:00000000 02:00000979 00000000  1000        0 45618 2 0000000000000000 24 4 30 10 -1
9: 4700A8C0:D804 41138F03:01BB 01 00000000:00000000 02:00000749 00000000  1000        0 48888 2 0000000000000000 24 4 29 10 -1
10: 4700A8C0:B476 C0DDC322:01BB 01 00000000:00000000 02:00000891 0000000000000000  1000        0 53539 2 0000000000000000 25 4 19 44 -1
11: 4700A8C0:CA96 1670528C:01BB 01 00000000:00000000 02:000005FD 0000000000000000  1000        0 413944 2 0000000000000000 25 4 11 10 -1
12: 4700A8C0:C5B2 C2F93A0D:01BB 01 00000000:00000000 02:00000696 0000000000000000  1000        0 53183 2 0000000000000000 24 4 31 10 -1
13: 4700A8C0:DB1A 368E9D6C:0050 06 00000000:00000000 03:00000BA2 0000000000000000
14: 4700A8C0:A1B8 3AC2BA23:01BB 01 00000000:00000000 02:00000D90 00000000  1000        0 320391 2 0000000000000000 21 4 19 10 24
15: 4700A8C0:C35A 1A72528C:01BB 01 00000000:00000000 02:00000BDE 0000000000000000  1000        0 336976 2 0000000000000000 24 4 31 10 -1
`

	expectedStates := []uint64{
		0x0A, 0x0A, 0x0A, 0x06, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x06, 0x01, 0x01,
	}

	tests := []struct {
		name        string
		mockContent string
		expected    []uint64
		expectError bool
	}{
		{
			name:        "Valid TCP states",
			mockContent: mockData,
			expected:    expectedStates,
			expectError: false,
		},
		{
			name:        "Empty content",
			mockContent: "",
			expected:    []uint64{},
			expectError: false,
		},
		{
			name:        "Invalid content",
			mockContent: "invalid data",
			expected:    []uint64{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := MockReader(tt.mockContent)
			states, err := GetTCPStates(file)
			if (err != nil) != tt.expectError {
				t.Errorf("GetTCPStates() error = %v, wantErr %v", err, tt.expectError)
				return
			}
			if !compareSlices(states, tt.expected) {
				t.Errorf("GetTCPStates() = %v, want %v", states, tt.expected)
			}
		})
	}
}

func compareSlices(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// func TestIsTCPReady(t *testing.T) {
// 	mockData := `sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode
// 0: 3500007F:0035 00000000:0000 0A 00000000:00000000 00:00000000 00000000   101        0 21804 1 0000000000000000 100 0 0 10 5
// 1: 0100007F:0277 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 2370 1 0000000000000000 100 0 0 10 0
// 2: 00000000:9329 00000000:0000 0A 00000000:00000000 00:00000000 00000000     0        0 18878 1 0000000000000000 100 0 0 10 0
// 3: 4700A8C0:E842 BA70C923:01BB 06 00000000:00000000 03:00001486 00000000     0        0 0 3 0000000000000000
// 4: 4700A8C0:D30A B3DA1203:01BB 01 00000000:00000000 02:00000D00 00000000  1000        0 49990 2 0000000000000000 24 4 28 10 -1
// 5: 4700A8C0:D312 B3DA1203:01BB 01 00000000:00000000 02:0000094C 00000000  1000        0 45613 2 0000000000000000 24 4 30 10 -1
// 6: 4700A8C0:E794 01358F03:01BB 01 00000000:00000000 02:00000AD6 00000000  1000        0 335391 2 0000000000000000 24 4 28 10 -1
// 7: 4700A8C0:C036 0A834D63:01BB 01 00000000:00000000 02:00000A5B 00000000  1000        0 57850 3 0000000000000000 23 4 0 4 4
// 8: 4700A8C0:D31C B3DA1203:01BB 01 00000000:00000000 02:00000979 00000000  1000        0 45618 2 0000000000000000 24 4 30 10 -1
// 9: 4700A8C0:D804 41138F03:01BB 01 00000000:00000000 02:00000749 00000000  1000        0 48888 2 0000000000000000 24 4 29 10 -1
// 10: 4700A8C0:B476 C0DDC322:01BB 01 00000000:00000000 02:00000891 0000000000000000  1000        0 53539 2 0000000000000000 25 4 19 44 -1
// 11: 4700A8C0:CA96 1670528C:01BB 01 00000000:00000000 02:000005FD 0000000000000000  1000        0 413944 2 0000000000000000 25 4 11 10 -1
// 12: 4700A8C0:C5B2 C2F93A0D:01BB 01 00000000:00000000 02:00000696 0000000000000000  1000        0 53183 2 0000000000000000 24 4 31 10 -1
// 13: 4700A8C0:DB1A 368E9D6C:0050 06 00000000:00000000 03:00000BA2 0000000000000000
// 14: 4700A8C0:A1B8 3AC2BA23:01BB 01 00000000:00000000 02:00000D90 00000000  1000        0 320391 2 0000000000000000 21 4 19 10 24
// 15: 4700A8C0:C35A 1A72528C:01BB 01 00000000:00000000 02:00000BDE 0000000000000000  1000        0 336976 2 0000000000000000 24 4 31 10 -1
// `

// 	tests := []struct {
// 		name        string
// 		mockContent string
// 		mockStates  []uint64
// 		mockError   error
// 		iteration   int
// 		timeout     time.Duration
// 		expected    bool
// 	}{
// 		{
// 			name:        "TCP SYN_RECV state present",
// 			mockStates:  []uint64{TCP_SYN_RECV},
// 			mockContent: mockData,
// 			mockError:   nil,
// 			iteration:   1,
// 			timeout:     10 * time.Millisecond,
// 			expected:    false,
// 		},
// 		{
// 			name:        "TCP SYN_SENT state present",
// 			mockStates:  []uint64{TCP_SYN_SENT},
// 			mockContent: mockData,
// 			mockError:   nil,
// 			iteration:   1,
// 			timeout:     10 * time.Millisecond,
// 			expected:    false,
// 		},
// 		{
// 			name:        "TCP states ready",
// 			mockStates:  []uint64{TCP_ESTABLISHED, TCP_CLOSE},
// 			mockContent: mockData,
// 			mockError:   nil,
// 			iteration:   1,
// 			timeout:     10 * time.Millisecond,
// 			expected:    true,
// 		},
// 		{
// 			name:        "TCP states not ready after multiple iterations",
// 			mockStates:  []uint64{TCP_SYN_RECV},
// 			mockContent: mockData,
// 			mockError:   nil,
// 			iteration:   2,
// 			timeout:     10 * time.Millisecond,
// 			expected:    false,
// 		},
// 		{
// 			name:        "getTCPStates returns an error",
// 			mockStates:  nil,
// 			mockContent: mockData,
// 			mockError:   errSome,
// 			iteration:   1,
// 			timeout:     10 * time.Millisecond,
// 			expected:    false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			file := MockReader(tt.mockContent)
// 			isReady, err := IsTCPReady(GetTCPStates(file), tt.iteration, tt.timeout, 1234)
// 			if isReady != tt.expected {
// 				t.Errorf("IsTCPReady() = %v, want %v", isReady, tt.expected)
// 			}
// 			if err != nil && tt.mockError == nil {
// 				t.Errorf("IsTCPReady() unexpected error = %v", err)
// 			}
// 			if err == nil && tt.mockError != nil {
// 				t.Errorf("IsTCPReady() expected error = %v, got none", tt.mockError)
// 			}
// 		})
// 	}
// }
