using System;

class RemoteControlCar
{
    private int battery=100, speed, distance=0, batteryDrain;
    public RemoteControlCar(int speed, int batteryDrain){
        this.speed = speed;
        this.batteryDrain = batteryDrain;
    }

    public bool BatteryDrained() => this.battery - batteryDrain < 0;

    public int DistanceDriven() => this.distance;

    public void Drive()
    {
        if (!BatteryDrained())
        {
            this.distance += speed;
            this.battery -= batteryDrain;            
        }
    }

    public static RemoteControlCar Nitro() => new RemoteControlCar(50,4);
}

class RaceTrack
{
    
    private int distance;
    public RaceTrack(int distance){
        this.distance = distance;
    }

    public bool TryFinishTrack(RemoteControlCar car)
    {
        while (!car.BatteryDrained())
        {
            car.Drive();
        }
        return car.DistanceDriven() >= distance;
    }
}
