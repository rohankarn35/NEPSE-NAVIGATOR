package services

import (
	"strings"

	"github.com/rohankarn35/nepsemarketbot/models"
	"github.com/rohankarn35/nepsemarketbot/utils"
)

func FormatIPOMessage(ipo models.IPOAlertModel) string {
	openingDate := ConvertDate(ipo.OpeningDateAD, ipo.OpeningDateBS)
	closingDate := ConvertDate(ipo.ClosingDateAD, ipo.ClosingDateBS)
	status := "Upcoming"
	if ipo.Status != "Nearing" {
		status = ipo.Status
	}
	ipoType := ipo.ShareType
	if ipo.ShareType == "ordinary" {
		ipoType = "General Public"
	} else if ipo.ShareType == "Migrant Workers" {
		ipoType = "Foreign Employment"
	}

	returnText := "📢 *" + strings.ToUpper(status) + " " + ipo.Type + " ALERT* 📢\n" +
		"================================\n\n" +
		"🏢 *Company Name:* " + ipo.CompanyName + "\n" +
		"💼 *Symbol:* " + ipo.StockSymbol + "\n" +
		"📊 *Issue Type:* " + utils.CapitalizeFirstLetter(ipoType) + "\n" +
		"🏦 *Sector:* " + ipo.SectorName + "\n" +
		"📈 *Current Status:* " + strings.ToUpper(status) + "\n\n" +
		"💵 *Price Per Unit:* Rs. " + ipo.PricePerUnit + "\n" +
		"📅 *Min/Max Units:* " + ipo.MinUnits + " / " + ipo.MaxUnits + "\n\n" +
		"🔓 *Opening Date:* " + openingDate + "\n" +
		"🔒 *Closing Date:* " + closingDate + "\n" +
		"⏰ *Closing Time:* " + ipo.ClosingDateClosingTime + "\n\n" +
		"📜 *Share Registrar:* " + ipo.ShareRegistrar + "\n"

	if ipo.Rating != "" {
		returnText += "⭐ *Rating:* " + ipo.Rating + "\n"
	}

	returnText += "\n================================\n"

	return returnText

}

func FormatIPOAlertMessage(ipo models.IPOAlertModel) string {
	closingDate := ConvertDate(ipo.ClosingDateAD, ipo.ClosingDateBS)
	status := "Upcoming"
	if ipo.Status != "Nearing" {
		status = ipo.Status
	}
	ipoType := ipo.ShareType
	if ipo.ShareType == "ordinary" {
		ipoType = "General Public"
	} else if ipo.ShareType == "Migrant Workers" {
		ipoType = "Foreign Employment"
	}
	oversubs := GetIPOOverscribeData(ipo.StockSymbol)

	returnText := "⚠️ *HURRY UP! Only 1 Hour Left to Apply for " + strings.ToUpper(ipo.StockSymbol) + "!* ⚠️\n" +
		"================================\n\n" +
		"🏢 *Company Name:* " + ipo.CompanyName + "\n" +
		"💼 *Symbol:* " + strings.ToUpper(ipo.StockSymbol) + "\n" +
		"📊 *Issue Type:* " + utils.CapitalizeFirstLetter(ipoType) + "\n" +
		"🏦 *Sector:* " + ipo.SectorName + "\n" +
		"📈 *Current Status:* " + strings.ToUpper(status) + "\n\n" +
		"🔒 *Closing Date:* " + closingDate + "\n" +
		"⏰ *Closing Time:* " + ipo.ClosingDateClosingTime + "\n"

	if oversubs != "" {
		returnText += "🔢 *Oversubscription:* " + oversubs + "x\n"
	}

	returnText += "\n================================"

	return returnText

}
