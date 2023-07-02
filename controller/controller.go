package controller

import (
	"github.com/FarhanRizkiM/kibackend"
	"github.com/FarhanRizkiM/presensi/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/whatsauth"
)

var DataMahasiswa = "Mahasiswa"
var DataPresensi = "Presensi"
var DataMataKuliah = "MataKuliah"
var DataJadwalKuliah = "JadwalKuliah"
var DataDosen = "Dosen"

// type HTTPRequest struct {
// 	Header string `json:"header"`
// 	Body   string `json:"body"`
// }

// func Sink(c *fiber.Ctx) error {
// 	var req HTTPRequest
// 	req.Header = string(c.Request().Header.Header())
// 	req.Body = string(c.Request().Body())
// 	return c.JSON(req)
// }

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetMahasiswa(c *fiber.Ctx) error {
	getstatus := kibackend.GetDataMahasiswa("Farhan Rizki Maulana")
	return c.JSON(getstatus)
}

func GetPresensi(c *fiber.Ctx) error {
	getstatus := kibackend.GetDataPresensi("HADIR")
	return c.JSON(getstatus)
}

func GetMataKuliah(c *fiber.Ctx) error {
	getstatus := kibackend.GetDataMataKuliah("TI41264")
	return c.JSON(getstatus)
}

func GetJadwalKuliah(c *fiber.Ctx) error {
	getstatus := kibackend.GetDataJadwalKuliah("Senin")
	return c.JSON(getstatus)
}

func GetDosen(c *fiber.Ctx) error {
	getstatus := kibackend.GetDataDosen("Rolly Maulana Awangga,S.T.,MT.,CAIP, SFPC.")
	return c.JSON(getstatus)
}
