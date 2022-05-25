package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	params, err := validate(fromPath, offset, limit)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fromPath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if offset > 0 {
		_, err = file.Seek(params.offset, 0)
		if err != nil {
			return err
		}
	}

	tmpFile, err := ioutil.TempFile("", "temp.*")
	if err != nil {
		return err
	}

	err = copy(file, tmpFile)
	if err != nil {
		return err
	}

	err = os.Rename(tmpFile.Name(), toPath)
	if err != nil {
		return err
	}

	//newFile, err := os.Create(toPath)
	//if err != nil {
	//	return err
	//}
	//defer newFile.Close()

	//// 0 - вычисляем отступ от начала файла
	//_, err = file.Seek(offset, 0)
	//if err != nil {
	//	return err
	//}

	//buf := make([]byte, params.limit)
	//_, err = file.ReadAt(buf, params.offset)
	//if err != nil {
	//	if err != io.EOF {
	//		return err
	//	}
	//}
	//
	//_, err = newFile.Write(buf)
	//if err != nil {
	//	return err
	//}

	return nil
}

type copyParams struct {
	offset int64
	limit  int64
}

func validate(filePath string, offset int64, limit int64) (*copyParams, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	if !fileInfo.Mode().IsRegular() {
		return nil, ErrUnsupportedFile
	}

	if fileInfo.Size() < offset {
		return nil, ErrOffsetExceedsFileSize
	}

	if limit > (fileInfo.Size() - offset) {
		limit = fileInfo.Size() - offset
	}

	return &copyParams{
		offset: offset,
		limit:  limit,
	}, nil
}

func copy(src io.Reader, dst io.Writer) error {
	pBar := pb.Start64(limit)
	defer pBar.Finish()

	var total int64
	for {
		n, err := io.CopyN(dst, src, 1)
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		// sleep исключительно для демонстрации прогресс бара при работе с мелкими файлами
		time.Sleep(100 * time.Millisecond)
		pBar.Increment()

		total += n
		if total == limit {
			break
		}

	}

	return nil
}
