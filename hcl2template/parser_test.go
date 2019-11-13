package hcl2template

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/provisioner/file"
	"github.com/hashicorp/packer/provisioner/shell"
	"github.com/zclconf/go-cty/cty"
)

func getBasicParser() *Parser {
	return &Parser{
		Parser: hclparse.NewParser(),
		BuilderSchemas: mapOfBuilder(map[string]packer.Builder{
			"amazon-ebs":     &MockBuilder{},
			"virtualbox-iso": &MockBuilder{},
		}).Get,

		ProvisionersSchemas: mapOfProvisioner(map[string]packer.Provisioner{
			"shell": &shell.Provisioner{},
			"file":  &file.Provisioner{},
		}),
		// PostProvisionersSchemas: mapOfProvisioner(map[string]packer.PostProcessor{
		// 	"amazon-import": &amazon_import.PostProcessor{},
		// }),
		// CommunicatorSchemas: mapOfDecodable(map[string]Decodable{
		// 	"ssh":   &communicator.SSH{},
		// 	"winrm": &communicator.WinRM{},
		// }).Get,
	}
}

type mapOfBuilder map[string]packer.Builder

func (mob mapOfBuilder) Get(builder string) (packer.Builder, error) {
	d, found := mob[builder]
	var err error
	if !found {
		err = fmt.Errorf("Unknown entry %s", builder)
	}
	return d, err
}

type mapOfProvisioner map[string]packer.Provisioner

func (mop mapOfProvisioner) Get(provisioner string) (packer.Provisioner, error) {
	p, found := mop[provisioner]
	var err error
	if !found {
		err = fmt.Errorf("Unknown provisioner %s", provisioner)
	}
	return p, err
}

func (mod mapOfProvisioner) List() []string {
	res := []string{}
	for k := range mod {
		res = append(res, k)
	}
	return res
}

type parseTestArgs struct {
	filename string
}

type parseTest struct {
	name      string
	parser    *Parser
	args      parseTestArgs
	wantCfg   *PackerConfig
	wantDiags bool
}

func testParse(t *testing.T, tests []parseTest) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, gotDiags := tt.parser.Parse(tt.args.filename)
			if tt.wantDiags == (gotDiags == nil) {
				t.Errorf("Parser.Parse() unexpected diagnostics. %s", gotDiags)
			}
			if diff := cmp.Diff(tt.wantCfg, gotCfg,
				cmpopts.IgnoreUnexported(cty.Value{}),
				cmpopts.IgnoreTypes(HCL2Ref{}),
				cmpopts.IgnoreTypes([]hcl.Range{}),
				cmpopts.IgnoreTypes(hcl.Range{}),
				cmpopts.IgnoreInterfaces(struct{ hcl.Expression }{}),
				cmpopts.IgnoreInterfaces(struct{ hcl.Body }{}),
			); diff != "" {
				t.Errorf("Parser.Parse() wrong packer config. %s", diff)
			}

		})
	}
}

var (
	basicNestedMockConfig = NestedMockConfig{
		String:   "string",
		Int:      42,
		Int64:    43,
		Bool:     true,
		Trilean:  config.TriTrue,
		Duration: 10 * time.Second,
		MapStringString: map[string]string{
			"a": "b",
			"c": "d",
		},
		SliceString: []string{
			"a",
			"b",
			"c",
		},
	}

	basicMockBuilder = &MockBuilder{
		Config: MockConfig{
			NestedMockConfig: basicNestedMockConfig,
			// Nested:           basicNestedMockConfig,
			NestedSlice: []NestedMockConfig{
				// 	basicNestedMockConfig,
			},
		},
	}
)

func TestParse_source(t *testing.T) {
	defaultParser := getBasicParser()

	tests := []parseTest{
		{"two basic sources",
			defaultParser,
			parseTestArgs{"testdata/sources/basic.pkr.hcl"},
			&PackerConfig{
				Sources: map[SourceRef]*Source{
					SourceRef{
						Type: "virtualbox-iso",
						Name: "ubuntu-1204",
					}: {
						Type:    "virtualbox-iso",
						Name:    "ubuntu-1204",
						Builder: basicMockBuilder,
					},
				},
			},
			false,
		},
	}
	testParse(t, tests)
}
