// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//+build amd64

package linux

import (
	"gvisor.dev/gvisor/pkg/sentry/arch"
	"gvisor.dev/gvisor/pkg/sentry/kernel"
	"syscall"
	"os/exec"
)

func real_div() {

	pid,_,_ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if(pid == 0) {
		j := 0
		i := 1
		i = i/j 
		j = j/i 
		return
	}
	return
}

// Div0 Div 0 to trigger core dump
func Div0(t *kernel.Task, args arch.SyscallArguments) (uintptr, *kernel.SyscallControl, error) {
	exec.Command("export", "GOTRACEBACK=crash")
	go real_div();
	return 0, nil, nil
}