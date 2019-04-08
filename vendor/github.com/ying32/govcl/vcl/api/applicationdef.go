package api

func Application_Instance() uintptr {
	ret, _, _ := application_Instance.Call()
	return ret
}

func Application_CreateForm(app uintptr, initScale bool) uintptr {
	ret, _, _ := application_CreateForm.Call(app, GoBoolToDBool(initScale))
	return ret
}

func Application_Run(app uintptr) {
	application_Run.Call(app)
}

func Application_Initialize(obj uintptr) {
	application_Initialize.Call(obj)
}
