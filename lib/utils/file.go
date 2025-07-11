package utils

import (
	"archive/tar"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

// TarDirectory 将目录 srcDir 打包为 tarFile
func TarDirectory(srcPath, tarPath string) error {
	file, err := os.Create(tarPath)
	if err != nil {
		return err
	}
	defer file.Close()

	tw := tar.NewWriter(file)
	defer tw.Close()

	return filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 获取相对路径，确保 tar 包内部路径正确
		relPath, err := filepath.Rel(srcPath, path)
		if err != nil {
			return err
		}

		// 跳过根目录
		if relPath == "." {
			return nil
		}

		// 创建 tar header
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name = relPath

		// 写入 header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// 如果是普通文件，写入内容
		if !info.Mode().IsRegular() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(tw, f)
		return err
	})
}

func ReadYaml(path string) (map[string]interface{}, error) {
	byt, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	temp := map[string]interface{}{}
	err = yaml.Unmarshal(byt, &temp)
	if err != nil {
		return nil, err
	}
	return toSimpleMap(temp), nil
}

// toSimpleMap 将多级的Map转换为key为 a.b.c的map
func toSimpleMap(origin map[string]interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	var loop func(prefix string, root map[string]interface{})
	loop = func(prefix string, root map[string]interface{}) {
		for k, v := range root {
			nKey := prefix + "." + k
			if node, ok := v.(map[string]interface{}); ok {
				loop(nKey, node)
			} else {
				res[nKey] = v
			}
		}
	}
	for k, v := range origin {
		if node, ok := v.(map[string]interface{}); ok {
			loop(k, node)
		} else {
			res[k] = v
		}
	}
	return res
}
