package main

type exceptionHandler func(*blockParser)

const (
	bpf_sock_ops_cb_exception_first_line  = "**bpf_sock_ops_cb_flags_set(bpf_sock,**"
	bpf_sock_ops_cb_exception_second_line = "**bpf_sock->bpf_sock_ops_cb_flags & ~BPF_SOCK_OPS_RTO_CB_FLAG)**"
	bpf_sock_ops_cb_exception_replacement = "```c\nbpf_sock_ops_cb_flags_set(bpf_sock, bpf_sock->bpf_sock_ops_cb_flags & ~BPF_SOCK_OPS_RTO_CB_FLAG);\n```"
)

func getExceptionsMap() map[string]exceptionHandler {
	return map[string]exceptionHandler{
		bpf_sock_ops_cb_exception_first_line: func(p *blockParser) {
			if p.cursor+1 >= len(p.lines) ||
				trimIndent(p.lines[p.cursor+1]) != bpf_sock_ops_cb_exception_second_line {
				panic("exception handler for bpf_sock_ops_cb_flags_set did not find expected second line")
			}
			p.appendBlock(markdownBlock{
				text: bpf_sock_ops_cb_exception_replacement,
			})
			p.cursor += 2
		},
	}
}
