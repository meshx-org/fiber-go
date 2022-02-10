package kernel

type ISyscalls interface {
	HandleDuplicate(handle Handle, out *Handle)
	HandleReplace(handle Handle, out *Handle)
	HandleClose(handle Handle)
}