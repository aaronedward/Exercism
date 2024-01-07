using System;

public static class LogAnalysis 
{
    public static string SubstringAfter(this string str, string delim) => str.Substring(str.IndexOf(delim) + delim.Length);

    public static string SubstringBetween(this string str, string str1, string str2) => str.Substring(str.IndexOf(str1) + str1.Length, str.IndexOf(str2) - str.IndexOf(str1) - str1.Length);

    public static string Message(this string str) => SubstringAfter(str,": ");

    public static string LogLevel(this string str) => SubstringBetween(str, "[", "]");
}