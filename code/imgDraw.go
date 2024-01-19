package main

import (
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func resizeImage(inputPath, outputPath string, targetWidth, targetHeight int) error {
	// 打开图像文件
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 使用文件扩展名判断图像格式
	var img image.Image
	switch filepath.Ext(inputPath) {
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

	// 计算缩放比例
	bounds := img.Bounds()
	scale := float64(targetWidth) / float64(bounds.Dx())

	// 创建目标图像
	thumbnail := image.NewRGBA(image.Rect(0, 0, targetWidth, int(float64(bounds.Dy())*scale)))

	// 使用 Bilinear 滤波器等比例缩放
	draw.BiLinear.Scale(thumbnail, thumbnail.Bounds(), img, img.Bounds(), draw.Over, nil)

	// 创建目标文件
	thumbnailFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer thumbnailFile.Close()

	// 将缩略图写入目标文件
	err = jpeg.Encode(thumbnailFile, thumbnail, nil)
	if err != nil {
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
	thumbnailPath := filepath.Join(c, "img", "4.jpg")

	if err := resizeImage(sourceImagePath, thumbnailPath, 150, 0); err != nil {
		log.Fatal(err)
	}
}
