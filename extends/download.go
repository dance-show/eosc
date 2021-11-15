package extends

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/eolinker/eosc"
)

var (
	client             = &http.Client{}
	ErrorFileCorrupted = errors.New("the file corrupted")
)

type ExtenderInfo struct {
	Go       string `json:"go"`
	Arch     string `json:"arch"`
	Eosc     string `json:"eosc"`
	Sha      string `json:"Sha"`
	Status   int    `json:"status"`
	IsLatest bool   `json:"is_latest"`
	Create   string `json:"create"`
	Update   string `json:"update"`
	URL      string `json:"url"`
}

//DownLoadToRepository 下载指定版本的插件项目，并解压到仓库
func DownLoadToRepository(group, project, version string) error {
	info, err := ExtenderInfoRequest(group, project, version)
	if err != nil {
		return err
	}
	dest := LocalExtenderPath(group, project, version)
	tarPath := LocalExtendTarPath(group, project, version)
	resp, err := http.Get(info.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create(tarPath)
	if err != nil {
		return err
	}
	defer f.Close()
	size, err := io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	tarSha, err := eosc.FileSha1(f, size)
	if err != nil {
		return err
	}
	if tarSha != info.Sha {
		return ErrorFileCorrupted
	}
	return eosc.Decompress(tarPath, dest)
}

//DownLoadToRepositoryById 下载插件， id格式为 {group}:{project}[:{version}]
func DownLoadToRepositoryById(id string) error {
	group, project, version, err := DecodeExtenderId(id)
	if err != nil {
		return err
	}
	if version == "" {
		version = "latest"
	}
	return DownLoadToRepository(group, project, version)
}
