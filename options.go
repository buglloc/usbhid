package usbhid

type EnumerateOption interface {
	isEnumerateOption()
}

type GetOption interface {
	isGetOption()
}

type optionDeviceFilterFunc struct {
	EnumerateOption
	GetOption
	filter DeviceFilterFunc
}

func WithDeviceFilterFunc(fn DeviceFilterFunc) optionDeviceFilterFunc {
	return optionDeviceFilterFunc{
		filter: fn,
	}
}

type optionVidFilter struct {
	EnumerateOption
	GetOption
	vid uint16
}

func WithVidFilter(vid uint16) optionVidFilter {
	return optionVidFilter{
		vid: vid,
	}
}

type optionPidFilter struct {
	EnumerateOption
	GetOption
	pid uint16
}

func WithPidFilter(pid uint16) optionPidFilter {
	return optionPidFilter{
		pid: pid,
	}
}

type optionOpen struct {
	GetOption
	open bool
	lock bool
}

func WithOpen(open, lock bool) optionOpen {
	return optionOpen{
		open: open,
		lock: lock,
	}
}
