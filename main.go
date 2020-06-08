package main

/*
#cgo CFLAGS: -Ishared-lib
#cgo LDFLAGS: -Lshared-lib/output -ljudger
#include <runner.h>
*/
import "C"

import "fmt"

func main() {
	var (
		config C.struct_config
		result C.struct_result
	)

	config.max_cpu_time = 1000
	config.max_real_time = 2000
	config.max_memory = 512 * 1024 * 1024
	config.max_process_number = 200
	config.max_output_size = 10000
	config.max_stack = 32 * 1024 * 1024
	config.exe_path = C.CString("./a.out")
	// config.input_path = C.CString("1.in")
	// config.output_path = C.CString("1.out")
	config.error_path = C.CString("1.error")
	// config.args = C.CString("")
	// config.env = C.CString("")
	config.log_path = C.CString("judger.log")
	config.seccomp_rule_name = C.CString("c_pp")
	config.uid = 0
	config.gid = 0

	C.run(&config, &result)
	fmt.Println()
	fmt.Printf("debug:\n%+v\n%+v\n", config, result)
}
