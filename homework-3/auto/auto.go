package auto

import "fmt"

type Car struct {
	Brand           string
	Year            string
	CarVolume       float64
	IsEngineActive  bool
	IsWindowsOpen   bool
	FilingCarVolume int
}

type Truck struct {
	Brand             string
	Year              string
	TruckVolume       float64
	IsEngineActive    bool
	IsWindowsOpen     bool
	FilingTruckVolume int
}

func Cars() {
	FirstCar := Car{
		"Ford",
		"2020",
		20.00,
		true,
		false,
		10,
	}

	FirstTruck := Truck{
		"Volvo",
		"2010",
		2000.00,
		false,
		true,
		0,
	}

	fmt.Printf(
		"Автомобиль %s\n"+
			"Год выпуска %s\n"+
			"Объем багажника %.2f\n"+
			"Двигатель работает %t\n"+
			"Окна открыты %t\n"+
			"Багажник заполнен на %v проц.\n\n",
		FirstCar.Brand,
		FirstCar.Year,
		FirstCar.CarVolume,
		FirstCar.IsEngineActive,
		FirstCar.IsWindowsOpen,
		FirstCar.FilingCarVolume,
	)

	fmt.Printf("Грузовой автомобиль: %v\n", FirstTruck)
	fmt.Printf("Загружаем, закрываем окна и запускаем двигатель... \n\n")

	FirstTruck.FilingTruckVolume = 100
	FirstTruck.IsEngineActive = true
	FirstTruck.IsWindowsOpen = false

	fmt.Printf(
		"Грузовой автомобиль %s\n"+
			"Год выпуска %s\n"+
			"Объем кузова %.2f\n"+
			"Двигатель работает %t\n"+
			"Окна открыты %t\n"+
			"Кузов заполнен на %v проц.\n",
		FirstTruck.Brand,
		FirstTruck.Year,
		FirstTruck.TruckVolume,
		FirstTruck.IsEngineActive,
		FirstTruck.IsWindowsOpen,
		FirstTruck.FilingTruckVolume,
	)
}
