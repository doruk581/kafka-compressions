package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/golang/snappy"
	"github.com/klauspost/compress/zstd"
	"github.com/pierrec/lz4/v4"
)

func main() {
	// JSON dosyasının yolu
	filePath := "data.json"

	// JSON dosyasını oku
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("❌ Dosya okunurken hata oluştu:", err)
		return
	}

	// Orijinal JSON boyutu
	originalSize := len(jsonData)

	// 🔹 Zstd ile sıkıştırma
	zstdSize := compressZstd(jsonData)

	// 🔹 Gzip ile sıkıştırma
	gzipSize := compressGzip(jsonData)

	// 🔹 Snappy ile sıkıştırma
	snappySize := compressSnappy(jsonData)

	// 🔹 LZ4 ile sıkıştırma
	lz4Size := compressLz4(jsonData)

	// 📊 Sonuçları yazdır
	fmt.Printf("\n📌 Orijinal JSON boyutu: %.2f MB\n", bytesToMB(originalSize))
	fmt.Printf("✅ Zstd Sıkıştırılmış Boyut: %.2f MB | Oran: %.2f%% Küçülme: %.2f%%\n", bytesToMB(zstdSize), float64(zstdSize)/float64(originalSize)*100, 100-float64(zstdSize)/float64(originalSize)*100)
	fmt.Printf("✅ Gzip Sıkıştırılmış Boyut: %.2f MB | Oran: %.2f%% Küçülme: %.2f%%\n", bytesToMB(gzipSize), float64(gzipSize)/float64(originalSize)*100, 100-float64(gzipSize)/float64(originalSize)*100)
	fmt.Printf("✅ Snappy Sıkıştırılmış Boyut: %.2f MB | Oran: %.2f%% Küçülme: %.2f%%\n", bytesToMB(snappySize), float64(snappySize)/float64(originalSize)*100, 100-float64(snappySize)/float64(originalSize)*100)
	fmt.Printf("✅ LZ4 Sıkıştırılmış Boyut: %.2f MB | Oran: %.2f%% Küçülme: %.2f%%\n", bytesToMB(lz4Size), float64(lz4Size)/float64(originalSize)*100, 100-float64(lz4Size)/float64(originalSize)*100)
}

// 🔹 Zstd Sıkıştırma Fonksiyonu
func compressZstd(data []byte) int {
	encoder, err := zstd.NewWriter(nil)
	if err != nil {
		fmt.Println("❌ Zstd encoder oluşturulamadı:", err)
		return len(data)
	}
	defer encoder.Close()
	compressed := encoder.EncodeAll(data, nil)
	return len(compressed)
}

// 🔹 Gzip Sıkıştırma Fonksiyonu
func compressGzip(data []byte) int {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	_, err := gzipWriter.Write(data)
	if err != nil {
		fmt.Println("❌ Gzip sıkıştırma hatası:", err)
		return len(data)
	}
	gzipWriter.Close()
	return buf.Len()
}

// 🔹 Snappy Sıkıştırma Fonksiyonu
func compressSnappy(data []byte) int {
	compressed := snappy.Encode(nil, data)
	return len(compressed)
}

// 🔹 LZ4 Sıkıştırma Fonksiyonu
func compressLz4(data []byte) int {
	var buf bytes.Buffer
	lz4Writer := lz4.NewWriter(&buf)
	_, err := lz4Writer.Write(data)
	if err != nil {
		fmt.Println("❌ LZ4 sıkıştırma hatası:", err)
		return len(data)
	}
	lz4Writer.Close()
	return buf.Len()
}

// 🔹 Byte'ı MB'ye çeviren yardımcı fonksiyon
func bytesToMB(size int) float64 {
	return float64(size) / (1024 * 1024)
}
