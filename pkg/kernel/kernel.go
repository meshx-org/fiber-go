package kernel

import (
	"log"
	"os"
)

type Kernel struct {
	klog *log.Logger
}

func NewKernel() *Kernel {
	return &Kernel{
		klog: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (k *Kernel) HandleDuplicate(handle Handle, out *Handle) {
	k.klog.Printf("HandleDuplicate %#v %#v", handle, out)
}

func (k *Kernel) HandleReplace(handle Handle, out *Handle) {
	k.klog.Printf("HandleReplace %#v %#v", handle, out)
}

func (k *Kernel) HandleClose(handle Handle) {
	k.klog.Printf("HandleClose %#v", handle)
}

func (k *Kernel) Wait() {}
