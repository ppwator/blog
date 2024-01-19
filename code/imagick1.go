package main

import (
	"fmt"
	"gopkg.in/gographics/imagick.v3/imagick"
	"log"
	"os"
	"path/filepath"
)

/**
https://github.com/gographics/imagick. 本地环境占用空间比较大
官网: https://imagemagick.org/
*/

func resizeImg(inputPath, outputPath string, targetWidth int) error {
	// 初始化 imagick 库
	imagick.Initialize()
	defer imagick.Terminate()

	// 创建 MagickWand
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// 读取输入图像文件
	if err := mw.ReadImage(inputPath); err != nil {
		fmt.Println("ReadImage", err)
		return err
	}

	// 获取原始图像的宽度和高度
	originalWidth := mw.GetImageWidth()

	// 计算等比例缩放后的高度
	originalHeight := mw.GetImageHeight()
	targetHeight := int(float64(originalHeight) * float64(targetWidth) / float64(originalWidth))

	// 等比例缩放
	if err := mw.ResizeImage(uint(targetWidth), uint(targetHeight), imagick.FILTER_LANCZOS); err != nil {
		fmt.Println("ResizeImage", err)
		return err
	}

	// 写入输出图像文件
	if err := mw.WriteImage(outputPath); err != nil {
		fmt.Println("WriteImage", err)
		return err
	}

	fmt.Println("缩略图生成成功:", outputPath)
	return nil
}

func main() {
	c, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	sourceImagePath := filepath.Join(c, "img", "1.png")
	thumbnailPath := filepath.Join(c, "img", "5.jpg")

	if err := resizeImg(sourceImagePath, thumbnailPath, 150); err != nil {
		fmt.Println("resizeImg", err)
		log.Fatal(err)
	}
}
