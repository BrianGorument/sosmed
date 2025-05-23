package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/mail"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func getJwtSecret() []byte {
	jwtSecret := viper.GetString("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set in environment variables")
	}
	return []byte(jwtSecret)
}
func CreateJWTToken(userID uint ,userName string, userEmail string) (string, error) {

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userID,
		"userName":  userName,
		"userEmail": userEmail,
		"exp"  :  time.Now().Add(time.Hour * 1).Unix(),
	})
	
	jwtSecret := getJwtSecret()
	
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	jwtSecret := getJwtSecret()
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}


func ConvertToUint(value interface{}) (uint, error) {
	floatVal, ok := value.(float64)
	if !ok {
		return 0, fmt.Errorf("failed to convert %v to float64", value)
	}
	return uint(floatVal), nil
}

func ConvertToInt(value interface{}) (int, error) {
	floatVal, ok := value.(float64)
	if !ok {
		return 0, fmt.Errorf("failed to convert %v to float64", value)
	}
	return int(floatVal), nil
}

func UploadAndCompressMedia(fileHeader *multipart.FileHeader) (string, error) {
    if fileHeader == nil {
        return "", nil
    }
    // Buka file yang diunggah
    file, err := fileHeader.Open()
    if err != nil {
        return "", fmt.Errorf("failed to open file: %v", err)
    }
    defer file.Close()

    // Validasi tipe file (hanya JPEG dan PNG)
    allowedTypes := map[string]bool{
        "image/jpeg": true,
        "image/png":  true,
    }
    contentType := fileHeader.Header.Get("Content-Type")
    if !allowedTypes[contentType] {
        return "", fmt.Errorf("only JPEG and PNG are allowed")
    }

    // Validasi ukuran file (maksimal 2MB)
    const maxSize = 2 << 20 // 2MB
    if fileHeader.Size > maxSize {
        return "", fmt.Errorf("file size exceeds 2MB")
    }

    // Buat nama file unik berdasarkan timestamp
    ext := filepath.Ext(fileHeader.Filename)
    newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
    uploadPath := filepath.Join("config/uploads", newFileName)

    // Pastikan folder uploads ada
    if err := os.MkdirAll("config/uploads", os.ModePerm); err != nil {
        return "", fmt.Errorf("failed to create uploads directory: %v", err)
    }

    // Simpan file sementara
    tempPath := filepath.Join("config/uploads", "temp_"+newFileName)
    if err := os.WriteFile(tempPath, nil, 0644); err != nil {
        return "", fmt.Errorf("failed to create temp file: %v", err)
    }

    tempFile, err := os.OpenFile(tempPath, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return "", fmt.Errorf("failed to open temp file: %v", err)
    }

    // Salin isi file ke file sementara
    data, err := io.ReadAll(file)
    if err != nil {
        return "", fmt.Errorf("failed to read file: %v", err)
    }
    if _, err := tempFile.Write(data); err != nil {
        return "", fmt.Errorf("failed to write to temp file: %v", err)
    }

    // Kompresi gambar
    srcImg, err := imaging.Open(tempPath)
    if err != nil {
        return "", fmt.Errorf("failed to open image for compression: %v", err)
    }
    compressedImg := imaging.Resize(srcImg, 800, 0, imaging.Lanczos)

    // Simpan gambar yang sudah dikompresi
    if err := imaging.Save(compressedImg, uploadPath); err != nil {
        return "", fmt.Errorf("failed to save compressed image: %v", err)
    }

    // Hapus file sementara
    if err := os.Remove(tempPath); err != nil {
        // Hanya log, tidak perlu gagal
        fmt.Printf("Failed to delete temp file: %v\n", err)
    }
    
    return fmt.Sprintf("config/uploads/%s", newFileName), nil
}