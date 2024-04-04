using System;

class RemoteControlCar{

    //should implement string interpolation for DisplayDistance and BatteryDisplay
    public int meters = 0;
    public double battery = 100.0;

    public static RemoteControlCar Buy() => new RemoteControlCar();

    public string DistanceDisplay() =>  "Driven "+meters+" meters";

    public string BatteryDisplay() => (battery>0) ? "Battery at "+battery+"%" : "Battery empty";

    public void Drive()
    {
        if (battery == 0)
        {
            return;
        }

        meters += 20;
        battery -= 1;
    }
}
