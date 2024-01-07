class Lasagna
{
    public int ExpectedMinutesInOven() => 40;

    public int RemainingMinutesInOven(int minutesCooked) => ExpectedMinutesInOven() - minutesCooked;

    public int PreparationTimeInMinutes(int layers) {
        var layerTime = 2;
        return layers * layerTime;
    }

    public int ElapsedTimeInMinutes(int layers, int minutesCooked) => PreparationTimeInMinutes(layers) + minutesCooked;
}
