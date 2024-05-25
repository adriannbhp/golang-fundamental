package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name) // Baris ini memanggil fungsi filter dengan name sebagai argumennya.
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Asu" {
		return "..."
	} else {
		return name
	}
}

func goodWord(name string) string {
	if name == "Friend" {
		return "Ramah"
	} else {
		return name
	}
}

/*
*
Studi Kasus: Sistem Pemberitahuan otomatis
*/

/**
Implementasi
*/

// 1. Mendefinisikan Tipe Fungsi untuk Method Pengiriman
// Notification Func adalah tipe untuk function yang mengirimkan notifikasi
type NotificationFunc func(message string) error

// 2. Mengimplementasikan Function Pengiriman untuk Setiap Method
// KirimEmail untuk mengirim notifikasi melalui email
func KirimEmail(message string) error {
	fmt.Println("Mengirim email dengan pesan:", message)
	return nil
}

// KirimSMS untuk mengirim notifikasi melalui SMS
func KirimSMS(message string) error {
	fmt.Println("Mengirim SMS dengan pesan:", message)
	return nil
}

// KirimPushNotification mengirim notifikasi push
func KirimPushNotification(message string) error {
	fmt.Println("Mengirim push notification dengan pesan:", message)
	return nil
}

// 3. Function utama untuk mengirim Notifikasi
// KirimNotifikasi adalah function utama yang menerima method pengiriman sebagai parameter
func KirimNotifikasi(pengirim NotificationFunc, message string) error {
	return pengirim(message)
}

func main() {
	sayHelloWithFilter("Adrian", spamFilter)

	sayHelloWithFilter("Friend", goodWord)

	filter := spamFilter
	sayHelloWithFilter("Asu", filter)

	// 4. Pengunaan Function Utama dengan berbagai Method Pengiriman
	pesan := "Halo, ini adalah pesan notifikasi!"

	// MEngirim Email
	err := KirimNotifikasi(KirimEmail, pesan)
	if err != nil {
		fmt.Println("Gagal mengirim email:", err)
	}

	// Mengirim SMS
	err = KirimNotifikasi(KirimSMS, pesan)
	if err != nil {
		fmt.Println("Gagal mengirim SMS:", err)
	}

	// Mengirim push notification
	err = KirimNotifikasi(KirimPushNotification, pesan)
	if err != nil {
		fmt.Println("Gagal mengirim push norification:", err)
	}
}
