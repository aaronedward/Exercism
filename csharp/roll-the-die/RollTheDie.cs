using System;

public class Player
{
    public int RollDie()
    {
        //throw new NotImplementedException("Please implement the Player.RollDie() method");
        Random random = new Random();
        return random.Next(1,18);
    }

    public double GenerateSpellStrength()
    {
        //throw new NotImplementedException("Please implement the Player.GenerateSpellStrength() method");
        Random random = new Random();
        return random.NextDouble() * (99.99-0.0);
    }
}
