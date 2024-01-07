using System;

public static class LogAnalysis 
{
    public static string SubstringAfter(this string str, string delim) => str.Substring(delim,str.Length-delim);

    public static string SubstringBetween(this string str,string str1, string str2) => str.Substring(str.StartsWith(str1), str.EndsWith(str2));
    
    // TODO: define the 'Message()' extension method on the `string` type

    // TODO: define the 'LogLevel()' extension method on the `string` type
}