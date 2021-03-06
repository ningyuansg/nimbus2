package main

import (
    "nimbus2/lib"
)

func main() {
    download := lib.NewLatestDownload()
    lastEntry := lib.GetLastEntry()
    if download.FileName != lastEntry[0] {
        overlayImg := download.GetImage()
        overlayImg = lib.ResizeToMap(overlayImg)
        lib.AppendCSV(download.FileName, lib.AlphaPercentMask(overlayImg))
        overlayImg = lib.AddMap(overlayImg)
        lib.SaveToRecords(overlayImg, download.FileName)
    }
}
