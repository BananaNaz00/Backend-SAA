package routers

import (
	"myapp/controllers"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS("*"))

	//login dan register
	e.POST("auth/login", controllers.Login)
	e.POST("auth/register", controllers.Register)

	//protected := e.Group("/api")
	//e.Use(middlewares.RequiredAuth())

	//user
	e.GET("/user", controllers.DataUser)
	e.PUT("/user", controllers.UpdateProfile)
	e.DELETE("/user/:di", controllers.DeleteUser)
	e.Static("/foto/user", "Foto Ktp User")

	e.GET("/activitylog", controllers.DataActivityLog)
	e.POST("/activitylog", controllers.SimpanActivityLog)

	e.GET("/aksesjalan", controllers.DataAksesJalan)
	e.GET("/akseslistrik", controllers.DataAksesListrik)
	e.GET("/alasanpenyesuaian", controllers.DataAlasanPenyesuaian)

	//blok lahan
	e.GET("/bloklahan", controllers.DataBlokLahan)
	e.POST("/bloklahan", controllers.SimpanBlokLahan)
	e.PUT("/bloklahan/:id", controllers.UpdateBlokLahan)
	e.DELETE("/bloklahan/:id", controllers.DeleteBlokLahan)

	//dokumen kontrak
	e.GET("/dokumenkontrak", controllers.DataDokumenKontrak)
	e.GET("/dokumenkontrak/:id", controllers.GetDokumenKontrakByID)
	e.POST("/dokumenkontrak", controllers.SimpanDokumenKontrak, middleware.BodyLimit(10000))
	e.PUT("/dokumenkontrak/:id", controllers.UpdateDokumenKontrak)
	e.DELETE("/dokumenkontrak/:id", controllers.DeleteDokumenKontrak)
	e.Static("/files/dokumenkontrak", "Dokumen Kontrak")

	//fasilitas lahan
	e.GET("/fasilitaslahan", controllers.DataFasilitasLahan)
	e.POST("/fasilitaslahan", controllers.SimpanFasilitasLahan)
	e.PUT("/fasilitaslahan/:id", controllers.UpdateFasilitasLahan)
	e.DELETE("/fasilitaslahan/:id", controllers.DeleteFasilitasLahan)

	//insiden
	e.GET("/insiden", controllers.DataInsiden)
	e.POST("/insiden", controllers.SimpanInsiden)
	e.PUT("/insiden/:id", controllers.UpdateInsiden)
	e.DELETE("/insiden/:id", controllers.DeleteInsiden)

	//inspeksi alat
	e.GET("/inspeksialat", controllers.DataInspeksiAlat)
	e.POST("/inspeksialat", controllers.SimpanInspeksiAlat)
	e.PUT("/inspeksialat/:id", controllers.UpdateInspeksiAlat)
	e.DELETE("/inspeksialat/:id", controllers.DeleteInspeksiAlat)

	//inventory
	e.GET("/inventory", controllers.DataInventory)
	e.POST("/inventory", controllers.SimpanInventory)
	e.PUT("/inventory/:id", controllers.UpdateInventory)
	e.DELETE("/inventory/:id", controllers.DeleteInventory)

	//invetasi
	e.GET("/investasi", controllers.DataInvestasi)
	e.POST("/investasi", controllers.SimpanInvestasi)
	e.PUT("/investasi/:id", controllers.UpdateInvestasi)
	e.DELETE("/investasi/:id", controllers.DeleteInvestasi)

	//investor detail
	e.GET("/investordetail", controllers.DataInvestorDetail)
	e.POST("/investordetail", controllers.SimpanInvestorDetail)
	e.PUT("/investordetail/:id", controllers.UpdateInvestorDetail)
	e.DELETE("/investordetail/:id", controllers.DeleteInvestorDetail)

	e.GET("/kategoriaktivitas", controllers.DataKategoriAktivitas)
	e.GET("/kategoriinventory", controllers.DataKategoriInventory)
	e.GET("/kategorikendala", controllers.DataKategoriKendala)

	//komoditas
	e.GET("/komoditas", controllers.DataKomoditas)
	e.POST("/komoditas", controllers.SimpanKomoditas)
	e.PUT("/komoditas/:id", controllers.UpdateKomoditas)
	e.DELETE("/komoditas/:id", controllers.DeleteKomoditas)

	//komoditas lahan
	e.GET("/komoditaslahan", controllers.DataKomoditasLahan)
	e.POST("/komoditaslahan", controllers.SimpanKomoditasLahan)
	e.PUT("/komoditaslahan/:id", controllers.UpdateKomoditasLahan)
	e.DELETE("/komoditaslahan/:id", controllers.DeleteKomoditasLahan)

	//kondisi lingkungan
	e.GET("/kondisilingkungan", controllers.DataKondisiLingkungan)
	e.POST("/kondisilingkungan", controllers.SimpanKondisiLingkungan)
	e.PUT("/kondisilingkungan/:id", controllers.UpdateKondisiLingkungan)
	e.DELETE("/kondisilingkungan/:id", controllers.DeleteKondisiLingkungan)

	e.GET("/kualitassinyal", controllers.DataKualitasSinyal)

	//lahan
	e.GET("/lahan", controllers.DataLahan)
	e.POST("/lahan", controllers.SimpanLahan)
	e.GET("/lahan/:id", controllers.GetLahanById)
	e.DELETE("/lahan/:id", controllers.DeleteLahan)

	//lahan status log
	e.GET("/lahanstatuslog", controllers.DataLahanStatusLog)
	e.POST("/lahanstatuslog", controllers.SimpanLahanStatusLog)

	//opname material
	e.GET("/opnamematerial", controllers.DataOpnameMaterial)
	e.POST("/opnamematerial", controllers.SimpanOpnameMaterial)
	e.PUT("/opnamematerial/:id", controllers.UpdateOpnameMaterial)
	e.DELETE("/opnamematerial/:id", controllers.DeleteOpnameMaterial)

	//Role
	e.GET("/role", controllers.DataRole)
	e.POST("/role", controllers.SimpanRole)
	e.PUT("/role", controllers.UpdateRole)
	e.DELETE("/role/:id", controllers.DeleteRole)

	//Satuan
	e.GET("/satuan", controllers.Asatuan)

	//pemakaian material
	e.GET("/pemakaianmaterial", controllers.DataPemakaianMaterial)
	e.POST("/pemakaianmaterial", controllers.SimpanPemakaianMaterial)
	e.PUT("/pemakaianmaterial/:id", controllers.UpdatePemakaianMaterial)
	e.DELETE("/pemakaianmaterial/:id", controllers.DeletePemakaianMaterial)

	//tingkat Urgensi
	e.GET("/tingkaturgensi", controllers.DataTingkatUrgensi)

	//sumberAir
	e.GET("/sumberair", controllers.DataSumberAir)

	e.GET("/statususer", controllers.DataStatusUser)

	e.GET("/statuspenyelesaian", controllers.DataStatusPenyelesaian)

	e.GET("/statuslahan", controllers.DataStatusLahan)

	e.GET("/statusinvestasi", controllers.DataStatusInvestasi)

	e.GET("/statusalat", controllers.DataStatusAlat)

	// coba
	e.GET("/penerimaanbarang", controllers.DataPenerimaanBarang)
	e.POST("/penerimaanbarang", controllers.SimpanPenerimaanBarang)
	e.PUT("/penerimaanbarang/:id", controllers.UpdatePenerimaanBarang)
	e.DELETE("/penerimaanbarang/:id", controllers.DeletePenerimaanBarang)

	return e
}
