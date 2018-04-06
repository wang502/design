package errorHandling

// Errors are values, program the error to avoid repetitive error handling
/*
	_, err = fd.Write(p0[a:b])
	if err != nil {
    	return err
	}
	_, err = fd.Write(p1[c:d])
	if err != nil {
    	return err
	}
	_, err = fd.Write(p2[e:f])
	if err != nil {
    	return err
	}
	// and so on
*/
import "io"

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) write(buf []byte) {
	if ew.err != nil {
		return
	}
	_, ew.err = ew.w.Write(buf)
}

func write(fd io.Writer, p []byte) error {
	//...
	ew := &errWriter{w: fd}
	ew.write(p[0:10])
	ew.write(p[10:20])
	ew.write(p[20:30])
	if ew.err != nil {
		return ew.err
	}
	return nil
}
