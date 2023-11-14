package main

import "fmt"

type Parcel interface {
	PrintParcelInfo()
}

type ParcelAddresses struct {
	SenderAddress    string
	RecipientAddress string
}

func (pa ParcelAddresses) PrintAddresses() {
	fmt.Printf("Sender address: %s\n", pa.SenderAddress)
	fmt.Printf("Recipient address: %s\n", pa.RecipientAddress)
}

type Box struct {
	ParcelAddresses
}

func (box Box) PrintParcelInfo() {
	fmt.Println("Parcel type: Box")
	box.PrintAddresses()
}

type Envelope struct {
	ParcelAddresses
}

func (envelope Envelope) PrintParcelInfo() {
	fmt.Println("Parcel type: Envelope")
	envelope.PrintAddresses()
}

type SortingCenter struct {
}

func (sc SortingCenter) SendParcelByCargoCarrier(parcel Parcel) {
	fmt.Println("\nParcel sent by cargo carrier service.")
	parcel.PrintParcelInfo()
}

func (sc SortingCenter) SendParcelBySmallPackageCarrier(parcel Parcel) {
	fmt.Println("\nParcel sent by small package carrier service.")
	parcel.PrintParcelInfo()
}

func (sc SortingCenter) SendParcel(parcel Parcel) {
	switch p := parcel.(type) {
	case Box:
		sc.SendParcelByCargoCarrier(p)
	case Envelope:
		sc.SendParcelBySmallPackageCarrier(p)
	default:
		panic("unknown parcel type")
	}
}

func main() {
	fmt.Println("Lets GO!")

	box := Box{ParcelAddresses: ParcelAddresses{SenderAddress: "Kyiv", RecipientAddress: "London"}}
	envelop := Envelope{ParcelAddresses: ParcelAddresses{SenderAddress: "New York", RecipientAddress: "Berlin"}}

	sortCenter := SortingCenter{}
	sortCenter.SendParcel(box)
	sortCenter.SendParcel(envelop)

}
