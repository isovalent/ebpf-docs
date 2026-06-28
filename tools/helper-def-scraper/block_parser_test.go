package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlockParser(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		blocks []markdownBlock
	}{
		{
			// this is the standard pseudo list with a unique word `**BPF_MTU_CHK_SEGS**`
			// but that spans multiple lines.
			name: "pseudo_list_on_2_lines",
			input: []string{
				"\t**BPF_MTU_CHK_SEGS**",
				"\t\tfirst line",
				"",
				"\t\tsecond line",
				"",
				"\t\tthird line",
				"",
				"\tNew block",
			},
			blocks: []markdownBlock{
				{text: "* **BPF_MTU_CHK_SEGS**: first line second line third line"},
				{text: "New block"},
			},
		},
		{
			// this is a pseudo-list with multiple words on the first line.
			name: "pseudo_list_starts_with_multiple_words",
			input: []string{
				"\t**sizeof**(tuple->ipv4) example",
				"\t\tLook for an IPv4 socket.",
			},
			blocks: []markdownBlock{
				{text: "* **sizeof**(tuple->ipv4) example: Look for an IPv4 socket."},
			},
		},
		{
			// this is a regular list with multiple items and white lines in the middle.
			name: "regular_list",
			input: []string{
				"\tTest list format:",
				"\t\t* element 1. This line should be always in the same",
				"\t\t  list element",
				"",
				"\t\t*\telement 2.",
				"\t\t*\telement 3.",
				"",
				"",
				// if after the `*` we have a space and not a tab we don't touch the formatting.
				// it would be possible but it would increase the complexity of the parser.
				"\t\t* element 4.",
			},
			blocks: []markdownBlock{
				{text: "Test list format:"},
				{text: "* element 1. This line should be always in the same list element"},
				{text: "*\telement 2."},
				{text: "*\telement 3."},
				{text: "* element 4."},
			},
		},
		{
			name: "no_real_list",
			input: []string{
				"Returns",
				"\t**-EINVAL** if invalid *flags* are passed, zero otherwise.",
				"\t0 on success.",
				"\t**-ENOENT** if *task->mm* is NULL, or no vma contains *addr*.",
			},
			blocks: []markdownBlock{
				{text: "### Returns"},
				{text: "**-EINVAL** if invalid _flags_ are passed, zero otherwise. 0 on success. **-ENOENT** if _task->mm_ is NULL, or no vma contains _addr_."},
			},
		},
		{
			name: "code_block",
			input: []string{
				"",
				"\t::",
				"",
				"\t\tint ret;",
				"\t\tstruct bpf_tunnel_key key = {};",
				"",
				"\t\tret = bpf_skb_get_tunnel_key(skb, &key, sizeof(key), 0);",
				"\t\tif (ret < 0)",
				"\t\t\treturn TC_ACT_SHOT;\t// drop packet",
				"",
				"\t\tif (key.remote_ipv4 != 0x0a000001)",
				"\t\t\treturn TC_ACT_SHOT;\t// drop packet",
				"",
				"\t\treturn TC_ACT_OK;\t\t// accept packet",
				// a code block always terminate with a blank line so we put it also here in the test.
				"",
			},
			blocks: []markdownBlock{
				{text: "```\nint ret;\nstruct bpf_tunnel_key key = {};\n\nret = bpf_skb_get_tunnel_key(skb, &key, sizeof(key), 0);\nif (ret < 0)\n\treturn TC_ACT_SHOT;\t// drop packet\n\nif (key.remote_ipv4 != 0x0a000001)\n\treturn TC_ACT_SHOT;\t// drop packet\n\nreturn TC_ACT_OK;\t\t// accept packet\n```"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := newBlockParser(blockParserArgs{
				lines:  tt.input,
				debug:  true,
				output: os.Stdout,
			})
			blocks := parser.generateBlocks()
			require.Len(t, blocks, len(tt.blocks))
			for i, block := range blocks {
				require.Equal(t, tt.blocks[i].text, block.text)
			}
		})
	}
}
