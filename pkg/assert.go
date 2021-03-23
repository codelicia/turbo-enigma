package pkg

// TODO We should get rid of this
func Assert(e error) {
	if e != nil {
		panic(e)
	}
}
