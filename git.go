package gitfastimport

import (
	"io"
	"os"
	"strconv"
	"time"
)

const (
	dataStart = "data <<c9de958c\n"
	dataEnd   = "c9de958c\n"
	lf        = "\n"
	sp        = " "
)

type Signature struct {
	Name  string
	Email string
	When  time.Time
}

func WriteFileModify(w io.Writer, mode int, name string) (int, error) {
	return io.WriteString(w, "M 100"+strconv.Itoa(mode)+" inline "+name+"\n")
}

func WriteFileDelete(w io.Writer, name string) (int, error) {
	return io.WriteString(w, "D "+name+"\n")
}

func WriteFile(w io.Writer, f *os.File) (int, error) {
	total := 0

	n, err := WriteFileModify(w, 644, f.Name())
	if err != nil {
		return total, err
	}
	total += n

	stat, err := f.Stat()
	if err != nil {
		return total, err
	}

	n, err = io.WriteString(w, "data "+strconv.FormatInt(stat.Size(), 10)+"\n")
	if err != nil {
		return total, err
	}
	total += n

	n64, err := io.Copy(w, f)
	if err != nil {
		return total, err
	}
	total += int(n64)

	return total, nil
}

func WriteDataBegin(w io.Writer) (int, error) {
	return io.WriteString(w, dataStart)
}

func WriteDataEnd(w io.Writer) (int, error) {
	return io.WriteString(w, dataEnd)
}

func WriteSignature(w io.Writer, sig *Signature, sigType string) (int, error) {
	return io.WriteString(w,
		sigType+" "+
			sig.Name+" <"+sig.Email+"> "+
			strconv.FormatInt(sig.When.Unix(), 10)+" "+
			sig.When.Format("-0700")+"\n")
}

func WriteMessage(w io.Writer, message string) (int, error) {
	return io.WriteString(w,
		"data "+strconv.Itoa(len(message))+"\n"+
			message+"\n")
}

func WriteCommit(w io.Writer, branch string, message string, author *Signature, committer *Signature) (int, error) {
	total := 0

	n, err := io.WriteString(w, "commit "+branch+"\n")
	if err != nil {
		return total, err
	}
	total += n

	n, err = WriteSignature(w, author, "author")
	if err != nil {
		return total, err
	}
	total += n

	n, err = WriteSignature(w, committer, "committer")
	if err != nil {
		return total, err
	}
	total += n

	n, err = WriteMessage(w, message)
	if err != nil {
		return total, err
	}
	total += n

	return total, nil
}
