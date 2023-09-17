package errs

var tstSlice = "a;dksfhhggakdfjhgg"

func blah() {
	itr, err := Iterate(&tstSlice)

	if err == nil {
		itr.Filter(nil).ResSlice()
	}
}
