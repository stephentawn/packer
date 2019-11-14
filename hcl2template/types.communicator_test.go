package hcl2template

import (
	"testing"
)

func TestParse_communicator(t *testing.T) {
	defaultParser := getBasicParser()

	tests := []parseTest{
		{"basic build",
			defaultParser,
			parseTestArgs{"testdata/communicator/basic.pkr.hcl"},
			&PackerConfig{
				Communicators: map[CommunicatorRef]*Communicator{
					CommunicatorRef{
						Type: "ssh",
						Name: "vagrant",
					}: &Communicator{
						Type:         "ssh",
						Name:         "vagrant",
						Communicator: basicMockCommunicator,
					},
				},
			},
			false,
		},
	}
	testParse(t, tests)
}
