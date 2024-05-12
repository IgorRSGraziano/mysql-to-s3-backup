package compress

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

func CompressFile(filePath string, buf io.Writer) error {

	gw := gzip.NewWriter(buf)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	file, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(stat, stat.Name())

	if err != nil {
		return err
	}

	err = tw.WriteHeader(header)

	if err != nil {
		return err
	}

	_, err = io.Copy(tw, file)

	if err != nil {
		return err
	}

	return nil
}
