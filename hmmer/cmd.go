package hmmer

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func run(command, hmmfile, seqfile string) (hits []Hit) {
	tempFile, err := ioutil.TempFile(os.TempDir(), command)
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name())

	cmd := exec.Command(command, "--tblout", tempFile.Name(), "--noali", hmmfile, seqfile)
	stderr := new(bytes.Buffer)
	stdout := new(bytes.Buffer)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err := cmd.Run(); err != nil {
		log.Printf("%s STDOUT: %s\n", command, stdout.String())
		log.Printf("%s STDERR: %s\n", command, stderr.String())
		log.Panic(err)
	}

	rd, err := os.Open(tempFile.Name())
	if err != nil {
		panic(err)
	}
	defer rd.Close()

	hits = ParseTblReport(rd)
	return
}

func Search(hmmfile, seqfile string) []Hit {
	command := "hmmsearch"
	return run(command, hmmfile, seqfile)
}

func Scan(hmmfile, seqfile string) []Hit {
	command := "hmmscan"
	return run(command, hmmfile, seqfile)
}
