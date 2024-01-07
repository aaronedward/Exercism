using System;

static class LogLine
{
    public static string[] logLevels = {"[ERROR]:","[WARNING]:","[INFO]:"};
    public static string Message(string logLine)
    {
        string cleanedLogLine ="";
        foreach (string x in logLevels)
        {
            if (logLine.Contains(x))
            {
                cleanedLogLine = logLine.Replace(x,"").Trim();
            }
        }
        return cleanedLogLine;
    }

    public static string LogLevel(string logLine)
    {
        if (logLine.Contains("ERROR"))
        {
            return "error";
        } else if (logLine.Contains("WARNING")){
            return "warning";
        } else if (logLine.Contains("INFO")){
            return "info";
        } 
        return "";
    }

    public static string Reformat(string logLine) => $"{Message(logLine)} ({LogLevel(logLine)})";
}
