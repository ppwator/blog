package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

func main() {
	fmt.Println(111)
	c, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println("c:", c)
	// 源图像文件路径
	sourceImagePath := filepath.Join(c, "/img/1.png")

	// 目标缩略图路径
	thumbnailPath := filepath.Join(c, "/img/3.jpg")
	fmt.Println(sourceImagePath, thumbnailPath)
	// 打开图像文件
	file, err := os.Open(sourceImagePath)
	if err != nil {
		fmt.Println("open", err)
		log.Fatal(err)
	}
	defer file.Close()

	// 使用文件扩展名判断图像格式
	var img image.Image
	switch filepath.Ext(sourceImagePath) {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	default:
		log.Fatal("Unsupported image format")
	}
	if err != nil {
		log.Fatal(err)
	}
	// 指定缩略图的宽度和高度
	width := 1080
	height := 1920

	//// 使用第三方库生成缩略图
	//thumbnail := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	// 使用第三方库等比例缩放图像
	thumbnail := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	// 创建目标文件
	thumbnailFile, err := os.Create(thumbnailPath)
	if err != nil {
		log.Fatal(err)
	}
	defer thumbnailFile.Close()

	// 将缩略图写入目标文件
	err = jpeg.Encode(thumbnailFile, thumbnail, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("缩略图生成成功:", thumbnailPath)
}
